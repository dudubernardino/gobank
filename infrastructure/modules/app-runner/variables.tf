variable "app_runner_name" {
  type        = string
  description = "Name of the App Runner service"
}

variable "app_runner_env_secrets" {
  type        = map(string)
  default     = {}
  description = "Environment secrets for the App Runner service"
}
variable "app_runner_env_variables" {
  type        = map(string)
  default     = {}
  description = "Environment variables for the App Runner service"
}

variable "app_runner_tags" {
  type        = map(string)
  default     = {}
  description = "Tags to apply to the App Runner service"
}

variable "ecr_image" {
  type        = string
  description = "ECR image URI for the App Runner service"
}

variable "ecr_repository_type" {
  type        = string
  description = "Type of the ECR repository"
}
variable "vpc_id" {
  type        = string
  description = "VPC ID for the App Runner service"
}

variable "vpc_private_subnets" {
  type        = list(string)
  description = "List of private subnets for the VPC connector"
}


