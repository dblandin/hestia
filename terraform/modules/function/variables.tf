variable "aws_iam_role" {}
variable "handler" {}
variable "release_s3_bucket" {}
variable "release_s3_key" {}

variable "tags" {
  type = "map"
}

variable "variables" {
  type = "map"
  default = {}
}
