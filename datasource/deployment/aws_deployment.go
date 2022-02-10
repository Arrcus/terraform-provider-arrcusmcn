package datasource

import (
	"context"
	"encoding/json"
	"errors"
	"io/ioutil"

	"github.com/Arrcus/terraform-provider-arcorch/models"
	schemas "github.com/Arrcus/terraform-provider-arcorch/schemas/deployment"
	"github.com/Arrcus/terraform-provider-arcorch/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DataSourceAwsDeployment() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataAwsDeploymentRead,
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"credentials_id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"enable_high_availability": &schema.Schema{
				Type:     schema.TypeBool,
				Computed: true,
			},
			"enable_private_subnet": &schema.Schema{
				Type:     schema.TypeBool,
				Computed: true,
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
				Computed: true,
			},
			"region": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"vpc_id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"instance_type": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"public_subnet": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"private_subnet": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"coordinates_lat": &schema.Schema{
				Type:     schema.TypeFloat,
				Computed: true,
			},
			"coordinates_long": &schema.Schema{
				Type:     schema.TypeFloat,
				Computed: true,
			},
		},
	}
}

func dataAwsDeploymentRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	accessToken := m.(map[string]string)["access_token"]
	url := m.(map[string]string)["baseUrl"] + "deployments"
	res, err := utils.GetRequest(url, accessToken)
	if err != nil {
		return diag.FromErr(err)
	}

	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return diag.FromErr(err)
	}
	creds := make([]models.Deployment, 0)
	err = json.Unmarshal(resBody, &creds)
	if err != nil {
		return diag.FromErr(err)
	}

	for _, c := range creds {
		if *c.Name == d.Get("name").(string) {
			err = schemas.UpdateCommonDeploymentResource(&c, d)
			if err != nil {
				return diag.FromErr(err)
			}
			err = schemas.UpdateAwsDeploymentResource(c.AwsDeployment, d)
			if err != nil {
				return diag.FromErr(err)
			}
			d.SetId(c.ID)
			return diags
		}
	}
	return diag.FromErr(errors.New("Can't find AWS deployment with given name."))
}
