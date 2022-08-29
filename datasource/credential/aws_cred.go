package datasource

import (
	"context"
	"encoding/json"
	"errors"
	"io/ioutil"

	"github.com/Arrcus/terraform-provider-arrcusmcn/models"
	schemas "github.com/Arrcus/terraform-provider-arrcusmcn/schemas/credential"
	"github.com/Arrcus/terraform-provider-arrcusmcn/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DataSourceAwsCred() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceAwsCredRead,
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"access_key": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"secret_key": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceAwsCredRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	accessToken := m.(map[string]string)["access_token"]
	tenant := m.(map[string]string)["tenant"]
	url := m.(map[string]string)["baseUrl"] + "cloud_credentials?provider=aws" + "&tenant=" + tenant
	res, err := utils.GetRequest(url, accessToken)
	if err != nil {
		return diag.FromErr(err)
	}

	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return diag.FromErr(err)
	}
	creds := make([]models.Credentials, 0)
	err = json.Unmarshal(resBody, &creds)
	if err != nil {
		return diag.FromErr(err)
	}

	for _, c := range creds {
		if *c.Name == d.Get("name").(string) {
			schemas.ToAwsCredSchema(&c, d)
			d.SetId(c.ID)
			return diags
		}
	}
	return diag.FromErr(errors.New("Can't find aws credentail with given name."))
}
