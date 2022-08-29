package schemas

import (
	"github.com/Arrcus/terraform-provider-arrcusmcn/models"
	"github.com/Arrcus/terraform-provider-arrcusmcn/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func AwsDeploymentSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"name": &schema.Schema{
			Type:     schema.TypeString,
			Required: true,
		},
		"credentials_id": &schema.Schema{
			Type:     schema.TypeString,
			Required: true,
		},
		"enable_high_availability": &schema.Schema{
			Type:     schema.TypeBool,
			Required: true,
		},
		"enable_private_subnet": &schema.Schema{
			Type:     schema.TypeBool,
			Required: true,
		},
		"arc_orch_ip": &schema.Schema{
			Type:     schema.TypeString,
			Computed: true,
		},
		"action": &schema.Schema{
			Type:     schema.TypeString,
			Computed: true,
		},
		"status": &schema.Schema{
			Type:     schema.TypeString,
			Computed: true,
		},
		"status_id": &schema.Schema{
			Type:     schema.TypeInt,
			Computed: true,
		},
		"active_ip": &schema.Schema{
			Type:     schema.TypeString,
			Computed: true,
		},
		"active_ip_gateway": &schema.Schema{
			Type:     schema.TypeString,
			Computed: true,
		},
		"active_private_ip": &schema.Schema{
			Type:     schema.TypeString,
			Computed: true,
		},
		"backup_ip": &schema.Schema{
			Type:     schema.TypeString,
			Computed: true,
		},
		"backup_private_ip": &schema.Schema{
			Type:     schema.TypeString,
			Computed: true,
		},
		"private_cidr": &schema.Schema{
			Type:     schema.TypeString,
			Computed: true,
		},
		"ingress_sg": &schema.Schema{
			Type:     schema.TypeString,
			Computed: true,
		},
		"hub_number": &schema.Schema{
			Type:     schema.TypeInt,
			Computed: true,
		},
		"instance_key": &schema.Schema{
			Type:     schema.TypeString,
			Required: true,
		},
		"region": &schema.Schema{
			Type:     schema.TypeString,
			Required: true,
		},
		"vpc_id": &schema.Schema{
			Type:     schema.TypeString,
			Required: true,
		},
		"instance_type": &schema.Schema{
			Type:     schema.TypeString,
			Required: true,
		},
		"public_subnet": &schema.Schema{
			Type:     schema.TypeString,
			Required: true,
		},
		"private_subnet": &schema.Schema{
			Type:     schema.TypeString,
			Required: true,
		},
		"coordinates_lat": &schema.Schema{
			Type:     schema.TypeFloat,
			Computed: true,
		},
		"coordinates_long": &schema.Schema{
			Type:     schema.TypeFloat,
			Computed: true,
		},
		// "aws_access_key": &schema.Schema{
		// 	Type:     schema.TypeString,
		// 	Computed: true,
		// },
		// "aws_secret_key": &schema.Schema{
		// 	Type:     schema.TypeString,
		// 	Computed: true,
		// },
	}
}

func ToAwsDeploymentObj(d *schema.ResourceData) *models.Deployment {
	res := models.Deployment{}
	res.Name = utils.StrPtr(d.Get("name").(string))
	res.CredentialsID = *utils.StrPtr(d.Get("credentials_id").(string))
	res.Provider = models.ProvidersAws
	res.AwsDeployment = &models.AwsDeployment{
		InstanceKey:   *utils.StrPtr(d.Get("instance_key").(string)),
		InstanceType:  *utils.StrPtr(d.Get("instance_type").(string)),
		PrivateSubnet: *utils.StrPtr(d.Get("private_subnet").(string)),
		PublicSubnet:  *utils.StrPtr(d.Get("public_subnet").(string)),
		Region:        *utils.StrPtr(d.Get("region").(string)),
		VpcID:         *utils.StrPtr(d.Get("vpc_id").(string)),
	}
	res.ArcOrchIP = *utils.StrPtr(d.Get("arc_orch_ip").(string))
	res.Action = *utils.StrPtr(d.Get("action").(string))
	res.Status = *utils.StrPtr(d.Get("status").(string))
	res.StatusID = *utils.Int64Ptr(d.Get("status_id").(int))
	res.ActiveIP = *utils.StrPtr(d.Get("active_ip").(string))
	res.ActiveIPGateway = *utils.StrPtr(d.Get("active_ip_gateway").(string))
	res.ActivePrivateIP = *utils.StrPtr(d.Get("active_private_ip").(string))
	res.BackupIP = *utils.StrPtr(d.Get("backup_ip").(string))
	res.BackupPrivateIP = *utils.StrPtr(d.Get("backup_private_ip").(string))
	res.PrivateCidr = *utils.StrPtr(d.Get("private_cidr").(string))
	res.IngressSg = *utils.StrPtr(d.Get("ingress_sg").(string))
	res.HubNumber = *utils.Int64Ptr(d.Get("hub_number").(int))
	res.Coordinates = &models.Coordinates{
		Lat:  utils.Float64Ptr(d.Get("coordinates_lat").(float64)),
		Long: utils.Float64Ptr(d.Get("coordinates_long").(float64)),
	}
	res.EnableHighAvailability = d.Get("enable_high_availability").(bool)
	res.EnablePrivateSubnet = d.Get("enable_private_subnet").(bool)

	return &res
}

func UpdateAwsDeploymentResource(deployment *models.AwsDeployment, d *schema.ResourceData) error {
	err := d.Set("instance_key", deployment.InstanceKey)
	if err != nil {
		return err
	}

	err = d.Set("region", deployment.Region)
	if err != nil {
		return err
	}

	err = d.Set("vpc_id", deployment.VpcID)
	if err != nil {
		return err
	}

	err = d.Set("instance_type", deployment.InstanceType)
	if err != nil {
		return err
	}

	err = d.Set("public_subnet", deployment.PublicSubnet)
	if err != nil {
		return err
	}

	err = d.Set("private_subnet", deployment.PrivateSubnet)
	if err != nil {
		return err
	}

	return nil
}
