# <resource name> arcorch_aws_cred

The arcorch_aws_cred data source provides details of a specific AWS credential created on the Arrcus ArcOrchestrator.

## Example Usage

```hcl
data "arcorch_aws_cred" "aws_cred" {
  name = "aws_cred"
}
```

## Argument Reference

* `name` - (Required) A unique name for the specific AWS credential.

## Attribute Reference

* `id` - a unique identifier for the resource
* `access_key` - Access key of an AWS account.
* `secret_key` - Secret key of an AWS account.
