resource "aws_cloudwatch_log_group" "api" {
  name              = "/aws/lambda/hestia/${terraform.workspace}/api"
  retention_in_days = 30

  tags = "${var.tags}"
}

resource "aws_cloudwatch_log_group" "handler" {
  name              = "/aws/lambda/hestia/${terraform.workspace}/handler"
  retention_in_days = 30

  tags = "${var.tags}"
}
