variable "app_runner_name" {
  type        = string
  description = "Name of the App Runner service"
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
