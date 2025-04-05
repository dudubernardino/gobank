terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "5.94.1"
    }
  }

  backend "s3" {
    bucket  = "emb-tfstate-bucket"
    region  = "us-east-2"
    key     = "state/terraform.tfstate"
    encrypt = true
  }
}

provider "aws" {
  region = var.region
}

resource "aws_s3_bucket" "terraform_state" {
  bucket        = var.tfstate_bucket
  force_destroy = true

  lifecycle {
    prevent_destroy = true
  }

  tags = {
    Iac = "true"
  }
}

resource "aws_s3_bucket_versioning" "terraform_state_version" {
  bucket = aws_s3_bucket.terraform_state.bucket

  versioning_configuration {
    status = "Enabled"
  }

  depends_on = [aws_s3_bucket.terraform_state]
}
