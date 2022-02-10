# <resource name> arcorch_user

The arcorch_user data source provides details of Arrcus MCN account used in Provder which currently logging in.

## Example Usage

```hcl
data "arcorch_user" "user" {}

```

## Attribute Reference

* `id` - a unique identifier for the resource
* `username` - Name for the account.
* `password` - Password for the account.
* `email` - Email address for the account.