package provider

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/Arrcus/terraform-provider-arcorch/datasource"
	dc "github.com/Arrcus/terraform-provider-arcorch/datasource/credential"
	dd "github.com/Arrcus/terraform-provider-arcorch/datasource/deployment"
	"github.com/Arrcus/terraform-provider-arcorch/resource"
	rc "github.com/Arrcus/terraform-provider-arcorch/resource/credential"
	rd "github.com/Arrcus/terraform-provider-arcorch/resource/deployment"
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
			"arcorch_aws_deployment": rd.ResourceAwsDeployment(),
			"arcorch_aws_cred":       rc.ResourceAwsCredential(),
			"arcorch_user":           resource.ResourceUser(),
		},
		DataSourcesMap: map[string]*schema.Resource{
			"arcorch_user":           datasource.DataSourceUser(),
			"arcorch_aws_cred":       dc.DataSourceAwsCred(),
			"arcorch_aws_deployment": dd.DataSourceAwsDeployment(),
		},
		ConfigureContextFunc: providerConfigure,
	}
}

func providerConfigure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	ipAddr := d.Get("serverip").(string)
	port := d.Get("port").(string)
	username := d.Get("username").(string)
	password := d.Get("password").(string)
	loginUrl := fmt.Sprintf(`http://%s:%s/api/login`, ipAddr, port)
	var diags diag.Diagnostics

	postBody, _ := json.Marshal(map[string]string{
		"username": username,
		"password": password,
	})

	responseBody := bytes.NewBuffer(postBody)

	resp, err := http.Post(loginUrl, "application/json", responseBody)
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
	result["baseUrl"] = fmt.Sprintf(`http://%s:%s/api/`, ipAddr, port)
	return result, diags
}
