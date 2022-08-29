# <resource name> arrcusmcn_user

The arrcusmcn_user resource is for the use of creation and management of Arrcus MCN user accounts under a tenant.

## Example Usage

```hcl
resource "arrcusmcn_user" "user" {
  name = "John Example"
  username = "john@example.com"
  password = "example123"
  email = "john@example.com"
  roles = ["TenantOperator"]
}

```

## Argument Reference

* `name` - (Required) Full name for the account.
* `username` - (Required) Name for the account. Name must be ended with a tenant domain.
* `password` - (Required) Password used for login.
* `email` - (Required) Email for the account.
* `roles` - (Required) Role of the account.

## Attribute Reference

* `id` - a unique identifier for the resource