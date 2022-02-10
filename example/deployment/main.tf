terraform {
  required_providers {
    arcorch = {
      version = "1.0.0"
      source = "arrcus.com/arrcus/arrcusmcn"
    }
  }
}

provider "arrcusmcn" {
  username = "admin"
  password = "password"
  serverip = "1.2.3.4"
  port = "8000"
}

data "arcorch_aws_cred" "aws_cred" {
  name = "aws_cred"
}

resource "arcorch_aws_deployment" "arcorch_aws" {
  name = "aws_hub"
  credentials_id = data.arcorch_aws_cred.aws_cred.id
  public_subnet = "aws-subnet"
  region = "us-east-1"
  vpc_id = "aws-vpc"
  instance_key = "arcedge"
  instance_type = "t2.medium"
  private_subnet = ""
  enable_high_availability = false
  enable_private_subnet = false
}

output "arcedge" {
  value = arcorch_aws_deployment.arcorch_aws
}