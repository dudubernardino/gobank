module "vpc" {
  source = "./modules/vpc"

  vpc_name = var.vpc_name
  vpc_cidr = var.vpc_cidr

  vpc_azs             = var.vpc_azs
  vpc_private_subnets = var.vpc_private_subnets
  vpc_public_subnets  = var.vpc_public_subnets

  vpc_enable_nat_gateway = var.vpc_enable_nat_gateway
  vpc_enable_vpn_gateway = var.vpc_enable_vpn_gateway

  vpc_tags = {
    Iac  = "true"
    Name = "gobank-vpc"
  }
}

module "ecr" {
  source = "./modules/ecr"

  ecr_name            = var.ecr_name
  ecr_repository_type = var.ecr_repository_type

  ecr_tags = {
    Iac = "true"
  }
}

data "aws_secretsmanager_secret_version" "db_password" {
  secret_id = module.rds.rds_password_secret_id
}
module "app-runner" {
  source     = "./modules/app-runner"
  depends_on = [module.vpc, module.ecr]

  app_runner_name = var.app_runner_name

  ecr_image           = var.ecr_image
  ecr_repository_type = "ECR"

  vpc_id              = module.vpc.vpc_id
  vpc_private_subnets = module.vpc.vpc_private_subnets

  app_runner_env_variables = {
    GOBANK_DATABASE_PORT     = 5432
    GOBANK_DATABASE_NAME     = module.rds.rds_dbname
    GOBANK_DATABASE_USER     = module.rds.rds_username
    GOBANK_DATABASE_PASSWORD = data.aws_secretsmanager_secret_version.db_password.secret_string
    GOBANK_DATABASE_HOST     = module.rds.rds_hostname
  }

  app_runner_tags = {
    Iac = "true"
  }
}

module "rds" {
  source     = "./modules/rds"
  depends_on = [module.vpc]

  rds_name = var.rds_name

  vpc_id              = module.vpc.vpc_id
  vpc_private_subnets = module.vpc.vpc_private_subnets

  app_runner_sg = module.app-runner.app_runner_sg_id
  rds_tags = {
    Iac = "true"
  }
}
