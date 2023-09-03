output "sensor_url" {
  value       = aws_lambda_function_url.sensor.function_url
  description = "Public Lambda function URL"
}
