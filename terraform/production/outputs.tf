output "endpoint_url" {
  value = "https://${module.api.rest_api_id}.execute-api.${data.aws_region.current.name}.amazonaws.com/${module.api.stage_name}"
}
