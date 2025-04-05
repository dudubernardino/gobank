###############################
####### OIDC Permissions ######
###############################
resource "aws_iam_openid_connect_provider" "oidc-github" {
  url = "https://token.actions.githubusercontent.com"

  client_id_list = [
    "sts.amazonaws.com",
  ]

  tags = {
    Iac = "true"
  }
}

###############################
####### TF Permissions ########
###############################
resource "aws_iam_role" "tf-role" {
  name        = "tf-role"
  description = "Terraform role for GitHub Actions"

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

  tags = {
    Iac = "true"
  }
}

resource "aws_iam_role_policy" "tf_role_policy" {
  name = "tf-role-policy"
  role = aws_iam_role.tf-role.id


  policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Sid      = "Statement1"
        Action   = "ecr:*"
        Effect   = "Allow"
        Resource = "*"
      },
      {
        Sid      = "Statement2"
        Action   = "iam:*"
        Effect   = "Allow"
        Resource = "*"
      },
      {
        Sid      = "Statement3"
        Action   = "s3:*"
        Effect   = "Allow"
        Resource = "*"
      }
    ]
  })


  depends_on = [
    aws_iam_role.tf-role,
  ]
}

