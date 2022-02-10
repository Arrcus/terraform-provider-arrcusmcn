package schemas

import (
	"github.com/Arrcus/terraform-provider-arcorch/models"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func UpdateCommonDeploymentResource(deployment *models.Deployment, d *schema.ResourceData) error {
	err := d.Set("arc_orch_ip", deployment.ArcOrchIP)
	if err != nil {
		return err
	}
	err = d.Set("action", deployment.Action)
	if err != nil {
		return err
	}
	err = d.Set("status", deployment.Status)
	if err != nil {
		return err
	}
	err = d.Set("status_id", int(deployment.StatusID))
	if err != nil {
		return err
	}
	err = d.Set("active_ip", deployment.ActiveIP)
	if err != nil {
		return err
	}
	err = d.Set("active_ip_gateway", deployment.ActiveIPGateway)
	if err != nil {
		return err
	}
	err = d.Set("active_private_ip", deployment.ActivePrivateIP)
	if err != nil {
		return err
	}
	err = d.Set("backup_ip", deployment.BackupIP)
	if err != nil {
		return err
	}
	err = d.Set("backup_private_ip", deployment.BackupPrivateIP)
	if err != nil {
		return err
	}
	err = d.Set("private_cidr", deployment.PrivateCidr)
	if err != nil {
		return err
	}
	err = d.Set("ingress_sg", deployment.IngressSg)
	if err != nil {
		return err
	}
	err = d.Set("hub_number", int(deployment.HubNumber))
	if err != nil {
		return err
	}
	err = d.Set("enable_high_availability", deployment.EnableHighAvailability)
	if err != nil {
		return err
	}
	err = d.Set("enable_private_subnet", deployment.EnablePrivateSubnet)
	if err != nil {
		return err
	}
	err = d.Set("credentials_id", deployment.CredentialsID)
	return nil
}
