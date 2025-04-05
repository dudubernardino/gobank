resource "aws_security_group" "dudu_sg" {
  name        = "${var.rds_name}-dudu-sg"
  description = "Allow My IP to connect to RDS"
  vpc_id      = var.vpc_id

  ingress {
    from_port   = 22
    to_port     = 22
    protocol    = "tcp"
    cidr_blocks = ["177.91.141.190/32"]
    description = "SSH from my machine"
  }

  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }

  tags = var.rds_tags
}


resource "aws_security_group" "rds_sg" {
  name        = "${var.rds_name}-sg"
  description = "Allow App Runner to connect to RDS"
  vpc_id      = var.vpc_id

  ingress {
    from_port       = 5432
    to_port         = 5432
    protocol        = "tcp"
    security_groups = [aws_security_group.dudu_sg.id]
  }

  ingress {
    from_port       = 5432
    to_port         = 5432
    protocol        = "tcp"
    security_groups = [var.app_runner_sg]
  }

  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }

  tags = var.rds_tags
}

resource "random_password" "rds_password" {
  length  = 16
  special = true
}


module "db" {
  source = "terraform-aws-modules/rds/aws"

  identifier = "${var.rds_name}-rds"

  engine            = "postgres"
  engine_version    = "15"
  family            = "postgres15"
  instance_class    = "db.t3.micro"
  allocated_storage = 20

  db_name  = var.rds_name
  username = "gobank"
  password = random_password.rds_password.result
  port     = "5432"

  publicly_accessible = false
  skip_final_snapshot = true

  vpc_security_group_ids = [aws_security_group.rds_sg.id]
  create_db_subnet_group = true
  subnet_ids             = var.vpc_private_subnets
  multi_az               = false

  storage_encrypted   = true
  deletion_protection = false

  tags = var.rds_tags
}

resource "aws_secretsmanager_secret" "db_password" {
  name = "${var.rds_name}-db-password"
}

resource "aws_secretsmanager_secret_version" "db_password_version" {
  secret_id     = aws_secretsmanager_secret.db_password.id
  secret_string = random_password.rds_password.result
}

resource "aws_key_pair" "bastion_key" {
  key_name   = "bastion-key"
  public_key = file("~/.ssh/bastion-key.pub")
}

resource "aws_instance" "bastion" {
  ami                         = "ami-0c55b159cbfafe1f0"
  instance_type               = "t3.micro"
  subnet_id                   = var.vpc_public_subnets[0]
  associate_public_ip_address = true
  vpc_security_group_ids      = [aws_security_group.rds_sg.id, aws_security_group.dudu_sg.id]
  key_name                    = aws_key_pair.bastion_key.key_name


  depends_on = [aws_key_pair.bastion_key]

  tags = {
    Iac = true
  }
}
