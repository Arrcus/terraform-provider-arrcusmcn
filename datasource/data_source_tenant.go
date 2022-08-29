package datasource

import (
	"context"
	"encoding/json"
	"io/ioutil"

	"github.com/Arrcus/terraform-provider-arrcusmcn/models"
	"github.com/Arrcus/terraform-provider-arrcusmcn/schemas"
	"github.com/Arrcus/terraform-provider-arrcusmcn/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DataSourceTenant() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceTenantRead,
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"organization": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"domain": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"defaultuser_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"defaultuser_username": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"defaultuser_password": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"defaultuser_email": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"defaultuser_is_default": &schema.Schema{
				Type:     schema.TypeBool,
				Computed: true,
			},
			"defaultuser_is_default_password": &schema.Schema{
				Type:     schema.TypeBool,
				Computed: true,
			},
			"defaultuser_roles": &schema.Schema{
				Type:     schema.TypeList,
				Required: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"numdeployments": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"numconnections": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"is_default": &schema.Schema{
				Type:     schema.TypeBool,
				Computed: true,
			},
		},
	}
}

func dataSourceTenantRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	accessToken := m.(map[string]string)["access_token"]
	url := m.(map[string]string)["baseUrl"] + "get_user_info"
	res, err := utils.GetRequest(url, accessToken)
	if err != nil {
		return diag.FromErr(err)
	}
	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return diag.FromErr(err)
	}

	user := models.User{}
	err = json.Unmarshal(resBody, &user)
	if err != nil {
		return diag.FromErr(err)
	}

	err = schemas.ToUserSchema(&user, d)
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(user.ID.String())
	return diags
}
