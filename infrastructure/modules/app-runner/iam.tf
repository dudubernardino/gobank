resource "aws_iam_role" "app_runner_role" {
  name        = "${var.app_runner_name}-role"
  description = "App-Runner role for GitHub Actions"

  assume_role_policy = jsonencode({
    Version : "2012-10-17",
    Statement : [
      {
        Effect : "Allow",
        Principal : {
          Service : "build.apprunner.amazonaws.com"
        },
        Action : "sts:AssumeRole"
      }
    ]
  })

  tags = var.app_runner_tags
}

resource "aws_iam_role_policy_attachment" "container_registry_read_only_access" {
  role       = aws_iam_role.app_runner_role.name
  policy_arn = "arn:aws:iam::aws:policy/AmazonEC2ContainerRegistryReadOnly"
}
