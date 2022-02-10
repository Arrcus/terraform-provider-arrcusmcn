# <resource name> arrcusmcn_aws_deployment

arrcusmcn_aws_deployment resource will be used to create ArcEdge deployments in the desired AWS account. 

## Example Usage

```hcl
data "arrcusmcn_aws_cred" "aws_cred" {
  name = "aws_cred"
}

resource "arrcusmcn_aws_deployment" "arrcusmcn_aws" {
  name = "aws_hub"
  credentials_id = data.arrcusmcn_aws_cred.aws_cred.id
  public_subnet = "aws-subnet"
  region = "us-east-1"
  vpc_id = "aws-vpc"
  instance_key = "arcedge"
  instance_type = "t2.medium"
  private_subnet = ""
  enable_high_availability = false
  enable_private_subnet = false
}
```

## Argument Reference

* `name ` - (Required) A unique name for an ArcEdge deployment running on.
* `credentials_id ` - (Required) The id of the AWS credential where the ArcEdge will be deployed.
* `region ` - (Required) Region where ArcEdge will be deployed.
* `instance_type ` - (Required) Instance size of the ArcEdge deployed.
* `instance_key ` - (Required) Instance key needed to ssh into the deployed ArcEdge.
* `vpc_id ` - (Required) VPC ID where the ArcEdge will be deployed.
* `public_subnet ` - (Required) Public subnet (should contain an attached Internet Gateway).
* `private_subnet ` - (Required) Private subnet when the ArcEdge is deployed as a spoke.
* `enable_high_availability ` - (Required) Set to true if high availability is desired. A pair of ArcEdge in active/standby mode will be deployed.
* `enable_private_subnet ` - (Required) Set to true, if ArcEdge is to be deployed as a spoke.

## Attribute Reference

* `id` - a unique identifier for the resource
* `arc_orch_ip ` -  IP of the ArcOrchestrator.
* `action ` -  Action initiated by the ArcOrch (Will be either CREATE/UPDATE/DELETE).
* `status ` -  Current status of the ArcEdge deployment (verbose).
* `status_id ` -  Current status of the ArcEdge deployment (numerical).
* `active_ip` - Public IP of the ArcEdge that has been deployed.
* `active_ip_gateway ` -  Default gateway for the deployed ArcEdge.
* `active_private_ip ` -  ArcEdge private IP.
* `backup_ip ` -  Backup ArcEdge public IP (when enable_high_availability is set to true).
* `backup_private_ip ` -  Backup ArcEdge private IP (when enable_high_availability is set to true).
* `private_cidr ` -  CIDR block of the private subnet.
* `ingress_sg ` -  Security group created for the ArcEdge deployment.
* `hub_number ` -  Hub number of the overlay network.
* `coordinates_lat ` -  Latitude where ArcEdge is deployed.
* `coordinates_long ` -  Longitude where ArcEdge is deployed.
