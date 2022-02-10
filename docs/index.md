# <provider> Arrcus-MCN

Arrcus-MCN provider is used to manage ArcEdge deployments running on different cloud platforms using Arrcus ArcOrchestrator. 

## Example Usage

```hcl
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
  port = "80"
}
```

## Argument Reference

* `username` - (Required) account username which will be used to login to  ArcOrchestrator.
* `password` - (Required) account password corresponding to given username
* `serverip` - (Required) ArcOrchestrator ip address
* `port` - (Required) ArcOrchestrator port
