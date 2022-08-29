terraform {
  required_providers {
    arrcusmcn = {
      version = "1.0.0"
      source = "arrcus.com/arrcus/arrcusmcn"
    }
  }
}

provider "arrcusmcn" {
  username = "admin"
  password = "Arrcus2018"
  serverip = "172.16.102.113"
  port = "443"
}

resource "arrcusmcn_tenant" "tenant" {
  name = "coke1"
  organization = "coke1"
  domain = "coke1.com"
  defaultuser_name = "coke1 coke1"
  defaultuser_username = "admin@coke1.com"
  defaultuser_password = "coke123"
  defaultuser_email = "coke1@coke1.com"
  defaultuser_roles = ["TenantAdmin", "TenantOperator"]
}
