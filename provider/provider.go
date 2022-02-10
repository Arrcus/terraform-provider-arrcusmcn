package provider

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"

	"github.com/Arrcus/terraform-provider-arrcusmcn/datasource"
	dc "github.com/Arrcus/terraform-provider-arrcusmcn/datasource/credential"
	dd "github.com/Arrcus/terraform-provider-arrcusmcn/datasource/deployment"
	"github.com/Arrcus/terraform-provider-arrcusmcn/resource"
	rc "github.com/Arrcus/terraform-provider-arrcusmcn/resource/credential"
	rd "github.com/Arrcus/terraform-provider-arrcusmcn/resource/deployment"
	"github.com/Arrcus/terraform-provider-arrcusmcn/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"username": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"password": &schema.Schema{
				Type:      schema.TypeString,
				Required:  true,
				Sensitive: true,
			},
			"serverip": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"port": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"arrcusmcn_aws_deployment": rd.ResourceAwsDeployment(),
			"arrcusmcn_aws_cred":       rc.ResourceAwsCredential(),
			"arrcusmcn_user":           resource.ResourceUser(),
		},
		DataSourcesMap: map[string]*schema.Resource{
			"arrcusmcn_user":           datasource.DataSourceUser(),
			"arrcusmcn_aws_cred":       dc.DataSourceAwsCred(),
			"arrcusmcn_aws_deployment": dd.DataSourceAwsDeployment(),
		},
		ConfigureContextFunc: providerConfigure,
	}
}

func providerConfigure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	ipAddr := d.Get("serverip").(string)
	port := d.Get("port").(string)
	username := d.Get("username").(string)
	password := d.Get("password").(string)
	loginUrl := fmt.Sprintf(`https://%s:%s/api/login`, ipAddr, port)
	var diags diag.Diagnostics

	user := map[string]string{
		"username": username,
		"password": password,
	}
	resp, err := utils.PostRequest(loginUrl, user, "")
	if err != nil {
		return nil, diag.FromErr(err)
	}
	if resp.StatusCode != 200 {
		body, _ := ioutil.ReadAll(resp.Body)
		sb := string(body)
		return nil, diag.FromErr(errors.New(sb))
	}
	var result map[string]string
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, diag.FromErr(err)
	}
	result["baseUrl"] = fmt.Sprintf(`https://%s:%s/api/`, ipAddr, port)
	return result, diags
}
