variable "vpc_name" {
  type        = string
  description = "The name of the VPC"
}

variable "vpc_cidr" {
  type        = string
  description = "The CIDR block for the VPC"
}

variable "vpc_azs" {
  type        = list(string)
  description = "The availability zones for the VPC"
}

variable "vpc_private_subnets" {
  type        = list(string)
  description = "The private subnets for the VPC"
}

variable "vpc_public_subnets" {
  type        = list(string)
  description = "The public subnets for the VPC"
}

variable "vpc_enable_nat_gateway" {
  type        = bool
  description = "Enable NAT Gateway for the VPC"
}

variable "vpc_enable_vpn_gateway" {
  type        = bool
  description = "Enable VPN Gateway for the VPC"
}

variable "vpc_tags" {
  type        = map(string)
  default     = {}
  description = "Tags for the VPC distribution"
}
