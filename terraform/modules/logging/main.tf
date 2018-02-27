resource "aws_cloudwatch_log_group" "api" {
  name              = "/aws/lambda/hestia-${terraform.workspace}-api"
  retention_in_days = 30

  tags = "${var.tags}"
}

resource "aws_cloudwatch_log_group" "handler" {
  name              = "/aws/lambda/hestia-${terraform.workspace}-handler"
  retention_in_days = 30

  tags = "${var.tags}"
}

resource "aws_cloudwatch_log_subscription_filter" "api_cloudwatch_log_subscription" {
  name            = "hestia-${terraform.workspace}-api-log-subscription"
  log_group_name  = "${aws_cloudwatch_log_group.api.name}"
  filter_pattern  = ""
  destination_arn = "${aws_lambda_function.log_forwarder.arn}"

  depends_on = ["aws_lambda_permission.allow_cloudwatch"]
}

resource "aws_cloudwatch_log_subscription_filter" "handler_cloudwatch_log_subscription" {
  name            = "hestia-${terraform.workspace}-handler-log-subscription"
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
  principal     = "logs.${var.region}.amazonaws.com"
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

  tags = "${var.tags}"
}
