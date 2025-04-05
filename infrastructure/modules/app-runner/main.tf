# resource "aws_apprunner_service" "app_runner_service" {
#   service_name = var.app_runner_name

#   source_configuration {
#     image_repository {
#       image_configuration {
#         port = "3000"
#       }
#       image_identifier      = var.ecr_image
#       image_repository_type = var.ecr_repository_type
#     }

#     auto_deployments_enabled = false
#   }

#   tags = var.app_runner_tags
# }
