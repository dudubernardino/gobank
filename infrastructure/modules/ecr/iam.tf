resource "aws_iam_role" "ecr_role" {
  name        = "${var.ecr_name}-role"
  description = "ECR role for GitHub Actions"

  assume_role_policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Effect = "Allow"
        Action = "sts:AssumeRoleWithWebIdentity"
        Principal = {
          Federated = "arn:aws:iam::994851302567:oidc-provider/token.actions.githubusercontent.com"
        }
        Condition = {
          StringEquals = {
            "token.actions.githubusercontent.com:aud" = [
              "sts.amazonaws.com"
            ]
            "token.actions.githubusercontent.com:sub" = [
              "repo:dudubernardino/gobank:ref:refs/heads/main"
            ]
          }
        }
      }
    ]
  })

  tags = var.ecr_tags
}

resource "aws_iam_role_policy" "ecr_role_policy" {
  name = "${var.ecr_name}-role-policy"
  role = aws_iam_role.ecr_role.id


  policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Sid      = "Statement1"
        Action   = "apprunner:*"
        Effect   = "Allow"
        Resource = "*"
      },
      {
        Sid = "Statement2"
        Action = [
          "iam:PassRole",
          "iam:CreateServiceLinkedRole",
        ]
        Effect   = "Allow"
        Resource = "*"
      },
      {
        Sid = "Statement3"
        Action = [
          "ecr:GetDownloadUrlForLayer",
          "ecr:BatchGetImage",
          "ecr:BatchCheckLayerAvailability",
          "ecr:PutImage",
          "ecr:InitiateLayerUpload",
          "ecr:UploadLayerPart",
          "ecr:CompleteLayerUpload",
          "ecr:GetAuthorizationToken",
        ]
        Effect   = "Allow"
        Resource = "*"
      }
    ]
  })


  depends_on = [
    aws_iam_role.ecr_role,
  ]
}
