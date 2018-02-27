output "rest_api_id" {
  value = "${aws_api_gateway_deployment.hestia.rest_api_id}"
}

output "stage_name" {
  value = "${aws_api_gateway_deployment.hestia.stage_name}"
}
