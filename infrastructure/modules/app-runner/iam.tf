resource "aws_iam_role" "app_runner_role" {
  name        = "${var.app_runner_name}-role"
  description = "App-Runner role for GitHub Actions"

  assume_role_policy = jsonencode({
    Version : "2012-10-17",
    Statement : [
      {
        Effect : "Allow",
        Principal : {
          Service : "tasks.apprunner.amazonaws.com"
        },
        Action : "sts:AssumeRole"
      }
    ]
  })

  tags = var.app_runner_tags
}

resource "aws_iam_role_policy" "app_runner_role_policy" {
  name = "${var.app_runner_name}-role-policy"
  role = aws_iam_role.app_runner_role.id


  policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Sid = "ECRReadOnlyAccess"
        Action = [
          "ecr:GetAuthorizationToken",
          "ecr:BatchCheckLayerAvailability",
          "ecr:GetDownloadUrlForLayer",
          "ecr:BatchGetImage",
          "ecr:GetRepositoryPolicy",
          "ecr:DescribeRepositories",
          "ecr:ListImages",
          "ecr:DescribeImages",
          "ecr:BatchGetRepositoryScanningConfiguration",
          "ecr:GetLifecyclePolicy"
        ]
        Effect   = "Allow"
        Resource = "*"
      }
    ]
  })


  depends_on = [
    aws_iam_role.app_runner_role,
  ]
}
