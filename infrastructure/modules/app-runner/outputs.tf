output "app_runner_sg_id" {
  value       = aws_security_group.app_runner_sg.id
  description = "The security group ID for the App Runner service"
}
