resource "aws_security_group" "app_runner_sg" {
  name        = "${var.app_runner_name}-sg"
  description = "Security group for App Runner to access RDS"
  vpc_id      = var.vpc_id

  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }

  tags = var.app_runner_tags
}

resource "aws_apprunner_vpc_connector" "app_runner_connector" {
  vpc_connector_name = "${var.app_runner_name}-vpc-connector"
  subnets            = var.vpc_private_subnets
  security_groups    = [aws_security_group.app_runner_sg.id]
}

resource "aws_apprunner_service" "app_runner_service" {
  service_name = var.app_runner_name

  source_configuration {
    image_repository {
      image_configuration {
        port                          = "3000"
        runtime_environment_secrets   = var.app_runner_env_secrets
        runtime_environment_variables = var.app_runner_env_variables
      }
      image_identifier      = var.ecr_image
      image_repository_type = var.ecr_repository_type
    }

    authentication_configuration {
      access_role_arn = aws_iam_role.app_runner_role.arn
    } # XXX: no example found in the provider docs

    auto_deployments_enabled = false
  }

  network_configuration {
    egress_configuration {
      egress_type       = "VPC"
      vpc_connector_arn = aws_apprunner_vpc_connector.app_runner_connector.arn
    }
  }


  tags = var.app_runner_tags
}
