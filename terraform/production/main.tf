terraform {
  backend "s3" {
    bucket               = "com.codeclimate.terraform"
    dynamodb_table       = "terraform-state-locks"
    encrypt              = true
    key                  = "hestia/terraform.tfstate"
    region               = "us-east-1"
    workspace_key_prefix = "env"
  }
}

variable "aws_assume_role_arn" {}
variable "release_s3_bucket" {}
variable "release_s3_key" {}
variable "syslog_udp_destination" {}

locals {
  tags = "${map("Terraform", "true", "ProductLine", "Hestia", "Environment", "production")}"
}

provider "aws" {
  region = "us-east-1"

  assume_role {
    role_arn = "${var.aws_assume_role_arn}"
  }
}


data "aws_caller_identity" "current" {}

data "aws_region" "current" {
  current = true
}

resource "aws_iam_role" "hestia" {
  name = "hestia-${terraform.workspace}"

  assume_role_policy = <<EOF
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Action": "sts:AssumeRole",
      "Principal": {
        "Service": "lambda.amazonaws.com"
      },
      "Effect": "Allow",
      "Sid": ""
    }
  ]
}
EOF
}

resource "aws_iam_policy" "hestia" {
  name        = "hestia-${terraform.workspace}"
  path        = "/"
  description = "IAM policy for hestia in ${terraform.workspace} environment."

  policy = <<EOF
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Effect": "Allow",
      "Action": [
        "lambda:InvokeFunction"
      ],
      "Resource": "*"
    },
    {
        "Effect": "Allow",
        "Action": [
            "ssm:DescribeParameters",
            "ssm:GetParameter",
            "ssm:GetParameters"
        ],
        "Resource": "*"
    }
  ]
}
EOF
}

resource "aws_iam_role_policy_attachment" "hestia" {
  role       = "${aws_iam_role.hestia.name}"
  policy_arn = "${aws_iam_policy.hestia.arn}"
}

resource "aws_iam_role_policy_attachment" "attachment" {
    role       = "${aws_iam_role.hestia.name}"
    policy_arn = "arn:aws:iam::aws:policy/service-role/AWSLambdaBasicExecutionRole"
}

resource "aws_lambda_function" "handler" {
  description   = "Hestia handler function - ${terraform.workspace}"
  function_name = "hestia-${terraform.workspace}-handler"
  handler       = "handler"
  publish       = true
  role          = "${aws_iam_role.hestia.arn}"
  runtime       = "go1.x"
  s3_bucket     = "${var.release_s3_bucket}"
  s3_key        = "${var.release_s3_key}"

  tags = "${local.tags}"
}

resource "aws_lambda_function" "api" {
  description   = "Hestia api function - ${terraform.workspace}"
  function_name = "hestia-${terraform.workspace}-api"
  handler       = "api"
  publish       = true
  role          = "${aws_iam_role.hestia.arn}"
  runtime       = "go1.x"
  s3_bucket     = "${var.release_s3_bucket}"
  s3_key        = "${var.release_s3_key}"

  tags = "${local.tags}"

  environment = {
    variables = "${map("COMMAND_LAMBDA_FUNCTION_NAME", "${aws_lambda_function.handler.function_name}", "COMMAND_LAMBDA_VERSION", "${aws_lambda_function.handler.version}")}"
  }
}

resource "aws_api_gateway_rest_api" "hestia" {
  name        = "hestia - ${terraform.workspace}"
  description = "hestia api - ${terraform.workspace}"
}

resource "aws_api_gateway_resource" "events" {
  rest_api_id = "${aws_api_gateway_rest_api.hestia.id}"
  parent_id   = "${aws_api_gateway_rest_api.hestia.root_resource_id}"
  path_part   = "events"
}

resource "aws_api_gateway_method" "events" {
  rest_api_id   = "${aws_api_gateway_rest_api.hestia.id}"
  resource_id   = "${aws_api_gateway_resource.events.id}"
  http_method   = "POST"
  authorization = "NONE"
}

resource "aws_api_gateway_integration" "api" {
  rest_api_id             = "${aws_api_gateway_rest_api.hestia.id}"
  resource_id             = "${aws_api_gateway_resource.events.id}"
  http_method             = "${aws_api_gateway_method.events.http_method}"
  integration_http_method = "POST"
  type                    = "AWS_PROXY"
  uri                     = "arn:aws:apigateway:${data.aws_region.current.name}:lambda:path/2015-03-31/functions/${aws_lambda_function.api.arn}/invocations"
}

resource "aws_lambda_permission" "apigw_lambda" {
  statement_id  = "AllowExecutionFromAPIGateway"
  action        = "lambda:InvokeFunction"
  function_name = "${aws_lambda_function.api.arn}"
  principal     = "apigateway.amazonaws.com"

  # More: http://docs.aws.amazon.com/apigateway/latest/developerguide/api-gateway-control-access-using-iam-policies-to-invoke-api.html
  source_arn = "arn:aws:execute-api:${data.aws_region.current.name}:${data.aws_caller_identity.current.account_id}:${aws_api_gateway_rest_api.hestia.id}/*/${aws_api_gateway_method.events.http_method}${aws_api_gateway_resource.events.path}"
}

resource "aws_api_gateway_deployment" "hestia_prod" {
  depends_on = [
    "aws_api_gateway_method.events",
    "aws_api_gateway_integration.api"
  ]
  rest_api_id = "${aws_api_gateway_rest_api.hestia.id}"
  stage_name = "production"
}

output "prod_url" {
  value = "https://${aws_api_gateway_deployment.hestia_prod.rest_api_id}.execute-api.${data.aws_region.current.name}.amazonaws.com/${aws_api_gateway_deployment.hestia_prod.stage_name}"
}

resource "aws_cloudwatch_log_group" "api" {
  name              = "/aws/lambda/${aws_lambda_function.api.function_name}"
  retention_in_days = 30

  tags = "${local.tags}"
}

resource "aws_cloudwatch_log_group" "handler" {
  name              = "/aws/lambda/${aws_lambda_function.handler.function_name}"
  retention_in_days = 30

  tags = "${local.tags}"
}

resource "aws_cloudwatch_log_subscription_filter" "api_cloudwatch_log_subscription" {
  name            = "${aws_lambda_function.api.function_name}-log-subscription"
  log_group_name  = "${aws_cloudwatch_log_group.api.name}"
  filter_pattern  = ""
  destination_arn = "${aws_lambda_function.log_forwarder.arn}"

  depends_on = ["aws_lambda_permission.allow_cloudwatch"]
}

resource "aws_cloudwatch_log_subscription_filter" "handler_cloudwatch_log_subscription" {
  name            = "${aws_lambda_function.handler.function_name}-log-subscription"
  log_group_name  = "${aws_cloudwatch_log_group.handler.name}"
  filter_pattern  = ""
  destination_arn = "${aws_lambda_function.log_forwarder.arn}"

  depends_on = ["aws_lambda_permission.allow_cloudwatch"]
}

resource "aws_iam_role" "log_forwarder" {
  name = "hestia-${terraform.workspace}-log-forwarder"

  assume_role_policy = <<EOF
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Action": "sts:AssumeRole",
      "Principal": {
        "Service": "lambda.amazonaws.com"
      },
      "Effect": "Allow",
      "Sid": ""
    }
  ]
}
EOF
}

resource "aws_lambda_permission" "allow_cloudwatch" {
  statement_id  = "hestia-${terraform.workspace}-allow-cloudwatch"
  action        = "lambda:InvokeFunction"
  function_name = "${aws_lambda_function.log_forwarder.function_name}"
  principal     = "logs.${data.aws_region.current.name}.amazonaws.com"
}

resource "aws_lambda_function" "log_forwarder" {
  description   = "Forwards hestia ${terraform.workspace} logs to Papertrail"
  function_name = "hestia-${terraform.workspace}-log-forwarder"
  handler       = "main"
  publish       = true
  role          = "${aws_iam_role.log_forwarder.arn}"
  runtime       = "go1.x"
  s3_bucket     = "com.codeclimate.hestia"
  s3_key        = "log_forwarder.zip"

  environment {
    variables = {
      SYSLOG_UDP_DESTINATION = "${var.syslog_udp_destination}"
    }
  }

  tags = "${local.tags}"
}
