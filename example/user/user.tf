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

/*
Run terraform apply directly will create a new user with given info.
If you want to update user without create it, please follow:
1. Remove `resource "arrcusmcn_user" "user"` and then run terraform apply to get id of the user currently using.
2. Add `resource "arrcusmcn_user" "user"` with corresponding info.
3. Run `terraform import arrcusmcn_user.user {id}` to import existing user.
4. Run `terraform apply`
*/

resource "arrcusmcn_user" "user" {
  name = "coke1 coke2"
  username = "coke@coke1.com"
  password = "coke123"
  email = "coke@coke1.com"
  roles = ["TenantOperator"]
}

output "user" {
  value = resource.arrcusmcn_user.user
}
