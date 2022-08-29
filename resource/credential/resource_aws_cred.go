package resource

import (
	"context"
	"encoding/json"
	"io/ioutil"

	"github.com/Arrcus/terraform-provider-arrcusmcn/models"
	schemas "github.com/Arrcus/terraform-provider-arrcusmcn/schemas/credential"
	"github.com/Arrcus/terraform-provider-arrcusmcn/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ResourceAwsCredential() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceAwsCredentialCreate,
		ReadContext:   resourceAwsCredentialRead,
		UpdateContext: resourceAwsCredentialUpdate,
		DeleteContext: resourceAwsCredentialDelete,
		Schema:        schemas.AwsCredentialSchema(),
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func resourceAwsCredentialCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	accessToken := m.(map[string]string)["access_token"]
	tenant := m.(map[string]string)["tenant"]
	url := m.(map[string]string)["baseUrl"] + "cloud_credentials?tenant=" + tenant
	cred, err := schemas.ToAwsCredObj(d)
	res, err := utils.PostRequest(url, *cred, accessToken)
	if err != nil {
		return diag.FromErr(err)
	}
	resBody, _ := ioutil.ReadAll(res.Body)
	resCred := models.Credentials{}
	err = json.Unmarshal(resBody, &resCred)
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(resCred.ID)
	return diags
}

func resourceAwsCredentialRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	accessToken := m.(map[string]string)["access_token"]
	tenant := m.(map[string]string)["tenant"]
	url := m.(map[string]string)["baseUrl"] + "cloud_credentials/" + d.Id() + "?tenant=" + tenant
	res, err := utils.GetRequest(url, accessToken)
	if err != nil {
		return diag.FromErr(err)
	}

	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return diag.FromErr(err)
	}

	resCred := models.Credentials{}
	err = json.Unmarshal(resBody, &resCred)
	if err != nil {
		return diag.FromErr(err)
	}
	schemas.ToAwsCredSchema(&resCred, d)
	if err != nil {
		return diag.FromErr(err)
	}
	return diags
}

func resourceAwsCredentialUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	accessToken := m.(map[string]string)["access_token"]
	tenant := m.(map[string]string)["tenant"]
	url := m.(map[string]string)["baseUrl"] + "cloud_credentials/" + d.Id() + "?tenant=" + tenant
	cred, err := schemas.ToAwsCredObj(d)
	if err != nil {
		return diag.FromErr(err)
	}
	_, err = utils.PutRequest(url, *cred, accessToken)
	if err != nil {
		return diag.FromErr(err)
	}
	return diags
}

func resourceAwsCredentialDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	accessToken := m.(map[string]string)["access_token"]
	tenant := m.(map[string]string)["tenant"]
	url := m.(map[string]string)["baseUrl"] + "cloud_credentials/" + d.Id() + "?tenant=" + tenant
	err := utils.DeleteRequest(url, accessToken)
	if err != nil {
		return diag.FromErr(err)
	}
	return diags
}
