variable "ecr_name" {
  type        = string
  description = "The name of the ECR repository"
}

variable "ecr_repository_type" {
  type        = string
  description = "The type of the ECR repository"
}

variable "ecr_tags" {
  type        = map(string)
  default     = {}
  description = "Tags for the ECR repository"
}
