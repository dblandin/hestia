resource "aws_lambda_function" "function" {
  function_name = "hestia-${terraform.workspace}-${var.handler}"
  handler       = "${var.handler}"
  publish       = true
  role          = "${var.aws_iam_role}"
  runtime       = "go1.x"
  s3_bucket     = "${var.release_s3_bucket}"
  s3_key        = "${var.release_s3_key}"

  tags = "${var.tags}"

  environment = {
    variables = "${var.variables}"
  }
}
