# <resource name> arrcusmcn_aws_deployment

The arrcusmcn_aws_deployment data source provides details of a specific ArcEdge created using the Arrcus ArcOrchestrator.

## Example Usage

```hcl
data "arrcusmcn_aws_deployment" "arrcusmcn_aws" {
  name = "aws_hub"
}
```

## Argument Reference

* `name ` - (Required) A unique name for specific ArcEdge deployment running on AWS.

## Attribute Reference

* `id` - a unique identifier for the resource
* `credentials_id ` - The id of an AWS credential which owns the ArcEdge deployment.
* `region ` - Region where ArcEdge is deployed.
* `instance_type ` -  Instance size of the ArcEdge deployed.
* `instance_key ` -  Instance key needed to ssh into the deployed ArcEdge.
* `vpc_id ` -  VPC ID where the ArcEdge is deployed.
* `public_subnet ` - Public subnet ID where ArcEdge is deployed.
* `private_subnet ` -  Private subnet ID if ArcEdge is deployed as a spoke.
* `enable_high_availability ` -  If set to true, indicates pair of ArcEdges have been deployed to implement high availability .
* `enable_private_subnet ` -  Will be set to true, if ArcEdge is deployed as a spoke.
* `arc_orch_ip ` -  IP of the ArcOrchestrator.
* `active_ip` - Public IP of the ArcEdge that has been deployed.
* `action ` -  Action initiated by the ArcOrch (Will be either CREATE/UPDATE/DELETE)..
* `status ` -   Current status of the ArcEdge deployment (verbose).
* `status_id ` -   Current status of the ArcEdge deployment (numerical).
* `active_ip_gateway ` -  Default gateway for the deployed ArcEdge.
* `active_private_ip ` -  ArcEdge private IP.
* `backup_ip ` -  Backup ArcEdge public IP (when enable_high_availability is set to true).
* `backup_private_ip ` -  Backup ArcEdge private IP (when enable_high_availability is set to true).
* `private_cidr ` - CIDR block of the private subnet.
* `ingress_sg ` -  Security group created for the ArcEdge deployment.
* `hub_number ` -  Hub number if the ArcEdge is created as a hub.
* `coordinates_lat ` -  Latitude where ArcEdge is deployed.
* `coordinates_long ` -  Longitude where ArcEdge is deployed.
