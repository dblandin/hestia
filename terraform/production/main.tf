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

module "policies" {
  source = "../modules/policies"
}

module "logging" {
  source = "../modules/logging"

  region                 = "${data.aws_region.current.name}"
  syslog_udp_destination = "${var.syslog_udp_destination}"

  tags = "${local.tags}"
}

module "api" {
  source = "../modules/api"

  api_lambda_arn = "${module.function_api.arn}"
  region         = "${data.aws_region.current.name}"
  account_id     = "${data.aws_caller_identity.current.account_id}"
}

module "function_api" {
  source = "../modules/function"

  aws_iam_role      = "${module.policies.hestia_iam_role_arn}"
  handler           = "api"
  release_s3_bucket = "${var.release_s3_bucket}"
  release_s3_key    = "${var.release_s3_key}"

  variables = "${map("COMMAND_LAMBDA_FUNCTION_NAME", "${module.function_handler.name}", "COMMAND_LAMBDA_VERSION", "${module.function_handler.version}")}"

  tags = "${local.tags}"
}

module "function_handler" {
  source = "../modules/function"

  aws_iam_role      = "${module.policies.hestia_iam_role_arn}"
  handler           = "handler"
  release_s3_bucket = "${var.release_s3_bucket}"
  release_s3_key    = "${var.release_s3_key}"

  variables = "${map("NOOP", "noop")}"

  tags = "${local.tags}"
}
