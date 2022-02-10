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
  password = "password"
  serverip = "1.2.3.4"
  port = "8000"
}

data "arrcusmcn_user" "user" {
  
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
  username = "username"
  password = "user_password"
  email = "example@email.com"
}

output "user" {
  value = data.arrcusmcn_user.user
}
