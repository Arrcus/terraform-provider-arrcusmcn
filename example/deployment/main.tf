terraform {
  required_providers {
    arrcusmcn = {
      version = "1.0.0"
      source = "arrcus.com/arrcus/arrcusmcn"
    }
  }
}

provider "arrcusmcn" {
  username = "admin@coke1.com"
  password = "coke123"
  serverip = "172.16.102.113"
  port = "443"
}

data "arrcusmcn_aws_cred" "aws_cred" {
  name = "aws_cred"
}

resource "arrcusmcn_aws_deployment" "arrcusmcn_aws" {
  name = "tingcih-aws"
  credentials_id = data.arrcusmcn_aws_cred.aws_cred.id
  public_subnet = "subnet-0eafefaf92e79895d"
  region = "us-east-1"
  vpc_id = "vpc-07f11c57735551934"
  instance_key = "arcedge"
  instance_type = "t2.medium"
  private_subnet = "subnet-0e7cf30920754f767"
  enable_high_availability = false
  enable_private_subnet = false
}

output "arcedge" {
  value = arrcusmcn_aws_deployment.arrcusmcn_aws
}