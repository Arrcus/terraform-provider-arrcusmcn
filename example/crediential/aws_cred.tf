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

resource "arrcusmcn_aws_cred" "aws_cred" {
  name = "aws_cred"
  access_key = "AKIAU3M22ZAFDSLBIPKC"
  secret_key = "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"
}