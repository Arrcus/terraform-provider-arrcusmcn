terraform {
  required_providers {
    arcorch = {
      version = "1.0.0"
      source = "arrcus.com/arrcus/arrcusmcn"
    }
  }
}

provider "arcorch" {
  username = "admin"
  password = "password"
  serverip = "1.2.3.4"
  port = "8000"
}

data "arcorch_user" "user" {
  
}

/*
Run terraform apply directly will create a new user with given info.
If you want to update user without create it, please follow:
1. Remove `resource "arcorch_user" "user"` and then run terraform apply to get id of the user currently using.
2. Add `resource "arcorch_user" "user"` with corresponding info.
3. Run `terraform import arcorch_user.user {id}` to import existing user.
4. Run `terraform apply`
*/
resource "arcorch_user" "user" {
  username = "username"
  password = "user_password"
  email = "example@email.com"
}

output "user" {
  value = data.arcorch_user.user
}