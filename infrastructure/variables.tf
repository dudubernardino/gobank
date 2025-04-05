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

####################
####### VPC ########
####################

variable "vpc_name" {
  type        = string
  default     = "gobank-vpc"
  description = "The name of the VPC"
}

variable "vpc_cidr" {
  type        = string
  default     = "10.0.0.0/16"
  description = "The CIDR block for the VPC"
}

variable "vpc_azs" {
  type        = list(string)
  default     = ["us-east-2a", "us-east-2b", "us-east-2c"]
  description = "The availability zones for the VPC"
}

variable "vpc_private_subnets" {
  type        = list(string)
  default     = ["10.0.1.0/24", "10.0.2.0/24", "10.0.3.0/24"]
  description = "The private subnets for the VPC"
}

variable "vpc_public_subnets" {
  type        = list(string)
  default     = ["10.0.101.0/24", "10.0.102.0/24", "10.0.103.0/24"]
  description = "The public subnets for the VPC"
}

variable "vpc_enable_nat_gateway" {
  type        = bool
  default     = false
  description = "Enable NAT Gateway for the VPC"
}

variable "vpc_enable_vpn_gateway" {
  type        = bool
  default     = false
  description = "Enable VPN Gateway for the VPC"
}

###############################
############ RDS ##############
###############################

variable "rds_name" {
  type        = string
  default     = "gobankdb"
  description = "RDS name"
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


