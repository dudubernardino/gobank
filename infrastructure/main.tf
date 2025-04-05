module "ecr" {
  source = "./modules/ecr"

  ecr_name            = var.ecr_name
  ecr_repository_type = var.ecr_repository_type

  ecr_tags = {
    Iac = "true"
  }
}

module "app-runner" {
  source = "./modules/app-runner"

  app_runner_name = var.app_runner_name

  ecr_image           = var.ecr_image
  ecr_repository_type = var.ecr_repository_type

  app_runner_tags = {
    Iac = "true"
  }
}
