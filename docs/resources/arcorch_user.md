# <resource name> arrcusmcn_user

The arrcusmcn_user resource is for the use of creation and management of Arrcus MCN accounts.

## Example Usage

```hcl
resource "arrcusmcn_user" "user" {
  username = "username"
  password = "password"
  email = "username@email.com"
}

```

## Argument Reference

* `username` - (Required) Name for the account.
* `password` - (Required) Password for the account.
* `email` - (Required) Email address for the account.

## Attribute Reference

* `id` - a unique identifier for the resource