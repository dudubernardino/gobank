###############################
####### Project Config ########
###############################

variable "region" {
  type        = string
  description = "AWS region"
}

###############################
####### S3 Bucket #############
###############################

variable "tfstate_bucket" {
  type        = string
  default     = "emb-tfstate-bucket"
  description = "S3 bucket for Terraform state"
}

###############################
############ ECR ##############
###############################

variable "ecr_name" {
  type        = string
  default     = "gobank-ecr"
  description = "ECR name"
}

variable "ecr_repository_type" {
  type        = string
  default     = "private"
  description = "ECR repository type"
}

variable "ecr_image" {
  type        = string
  description = "ECR image"
}

###############################
######### App Runner ##########
###############################

variable "app_runner_name" {
  type        = string
  default     = "gobank-app-runner"
  description = "App Runner name"
}


