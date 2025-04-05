module "ecr" {
  source = "terraform-aws-modules/ecr/aws"

  repository_name = var.ecr_name

  repository_image_tag_mutability = "MUTABLE"

  repository_image_scan_on_push = true

  repository_read_write_access_arns = []
  repository_lifecycle_policy = jsonencode({
    rules = [
      {
        rulePriority = 1,
        description  = "Keep last 30 images",
        selection = {
          tagStatus     = "tagged",
          tagPrefixList = ["v"],
          countType     = "imageCountMoreThan",
          countNumber   = 30
        },
        action = {
          type = "expire"
        }
      }
    ]
  })

  repository_type = var.ecr_repository_type

  tags = var.ecr_tags
}
