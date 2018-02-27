resource "aws_api_gateway_rest_api" "hestia" {
  name        = "hestia-${terraform.workspace}-api"
  description = "hestia ${terraform.workspace} api"
}

resource "aws_api_gateway_resource" "github" {
  rest_api_id = "${aws_api_gateway_rest_api.hestia.id}"
  parent_id   = "${aws_api_gateway_rest_api.hestia.root_resource_id}"
  path_part   = "github"
}

resource "aws_api_gateway_resource" "github_webhooks" {
  rest_api_id = "${aws_api_gateway_rest_api.hestia.id}"
  parent_id   = "${aws_api_gateway_resource.github.id}"
  path_part   = "webhooks"
}

resource "aws_api_gateway_method" "github_webhooks" {
  rest_api_id   = "${aws_api_gateway_rest_api.hestia.id}"
  resource_id   = "${aws_api_gateway_resource.github_webhooks.id}"
  http_method   = "POST"
  authorization = "NONE"
}

resource "aws_api_gateway_integration" "github_webhooks" {
  rest_api_id             = "${aws_api_gateway_rest_api.hestia.id}"
  resource_id             = "${aws_api_gateway_resource.github_webhooks.id}"
  http_method             = "${aws_api_gateway_method.github_webhooks.http_method}"
  integration_http_method = "POST"
  type                    = "AWS_PROXY"
  uri                     = "arn:aws:apigateway:${var.region}:lambda:path/2015-03-31/functions/${var.api_lambda_arn}/invocations"
}

resource "aws_api_gateway_resource" "slack" {
  rest_api_id = "${aws_api_gateway_rest_api.hestia.id}"
  parent_id   = "${aws_api_gateway_rest_api.hestia.root_resource_id}"
  path_part   = "slack"
}

resource "aws_api_gateway_resource" "slack_webhooks" {
  rest_api_id = "${aws_api_gateway_rest_api.hestia.id}"
  parent_id   = "${aws_api_gateway_resource.slack.id}"
  path_part   = "webhooks"
}

resource "aws_api_gateway_method" "slack_webhooks" {
  rest_api_id   = "${aws_api_gateway_rest_api.hestia.id}"
  resource_id   = "${aws_api_gateway_resource.slack_webhooks.id}"
  http_method   = "POST"
  authorization = "NONE"
}

resource "aws_api_gateway_integration" "slack_webhooks" {
  rest_api_id             = "${aws_api_gateway_rest_api.hestia.id}"
  resource_id             = "${aws_api_gateway_resource.slack_webhooks.id}"
  http_method             = "${aws_api_gateway_method.slack_webhooks.http_method}"
  integration_http_method = "POST"
  type                    = "AWS_PROXY"
  uri                     = "arn:aws:apigateway:${var.region}:lambda:path/2015-03-31/functions/${var.api_lambda_arn}/invocations"
}

resource "aws_lambda_permission" "lambda_permission" {
  statement_id  = "hestia-${terraform.workspace}-lambda-permission"
  action        = "lambda:InvokeFunction"
  function_name = "${var.api_lambda_arn}"
  principal     = "apigateway.amazonaws.com"

  # More: http://docs.aws.amazon.com/apigateway/latest/developerguide/api-gateway-control-access-using-iam-policies-to-invoke-api.html
  source_arn = "arn:aws:execute-api:${var.region}:${var.account_id}:${aws_api_gateway_rest_api.hestia.id}/${terraform.workspace}/*"
}

resource "aws_api_gateway_deployment" "hestia" {
  depends_on = [
    "aws_api_gateway_integration.github_webhooks",
    "aws_api_gateway_integration.slack_webhooks",
    "aws_api_gateway_method.github_webhooks",
    "aws_api_gateway_method.slack_webhooks",
  ]

  rest_api_id = "${aws_api_gateway_rest_api.hestia.id}"
  stage_name  = "${terraform.workspace}"
}
