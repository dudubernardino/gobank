variable "rds_name" {
  type        = string
  description = "The name of the database to create when the DB instance is created"
}

variable "rds_tags" {
  type        = map(string)
  default     = {}
  description = "Tags to apply to the RDS service"
}

variable "vpc_id" {
  type        = string
  description = "VPC ID for the App Runner service"
}

variable "vpc_private_subnets" {
  type        = list(string)
  description = "List of private subnets for the VPC connector"
}

variable "app_runner_sg" {
  type        = string
  description = "The security group ID of the App Runner service"
}
