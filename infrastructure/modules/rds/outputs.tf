output "rds_dbname" {
  value       = var.rds_name
  description = "value of the database name"
}

output "rds_username" {
  value       = "gobank"
  description = "value of the database username"
}

output "rds_password_secret_id" {
  value       = aws_secretsmanager_secret.db_password.id
  sensitive   = true
  description = "value of the database password secret id"
}

output "rds_hostname" {
  value       = module.db.db_instance_address
  description = "value of the database host"
}
