package resource

import (
	"context"
	"encoding/json"
	"io/ioutil"

	"github.com/Arrcus/terraform-provider-arrcusmcn/models"
	schemas "github.com/Arrcus/terraform-provider-arrcusmcn/schemas"
	"github.com/Arrcus/terraform-provider-arrcusmcn/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ResourceTenant() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceTenantCreate,
		ReadContext:   resourceTenantRead,
		UpdateContext: resourceTenantUpdate,
		DeleteContext: resourceTenantDelete,
		Schema:        schemas.TenantSchema(),
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func resourceTenantCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	accessToken := m.(map[string]string)["access_token"]

	url := m.(map[string]string)["baseUrl"] + "tenants"
	tenant, err := schemas.ToTenantObj(d)
	if err != nil {
		return diag.FromErr(err)
	}
	// return diag.FromErr(errors.New(fmt.Sprintf("%v", *tenant)))
	res, err := utils.PostRequest(url, *tenant, accessToken)
	if err != nil {
		return diag.FromErr(err)
	}
	resBody, _ := ioutil.ReadAll(res.Body)
	resTenant := models.Tenant{}
	err = json.Unmarshal(resBody, &resTenant)
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(resTenant.ID.String())
	return diags
}

func resourceTenantRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics
	accessToken := m.(map[string]string)["access_token"]
	url := m.(map[string]string)["baseUrl"] + "tenants/" + d.Id()
	// return diag.FromErr(errors.New("url"))
	res, err := utils.GetRequest(url, accessToken)
	// return diag.FromErr(err)
	if err != nil {
		return diag.FromErr(err)
	}

	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return diag.FromErr(err)
	}

	resTenant := models.Tenant{}
	err = json.Unmarshal(resBody, &resTenant)
	if err != nil {
		return diag.FromErr(err)
	}
	err = schemas.ToTenantSchema(&resTenant, d)
	if err != nil {
		return diag.FromErr(err)
	}
	return diags
}

func resourceTenantUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	accessToken := m.(map[string]string)["access_token"]
	url := m.(map[string]string)["baseUrl"] + "tenants/" + d.Id()
	user, err := schemas.ToUserObj(d)
	if err != nil {
		return diag.FromErr(err)
	}
	res, err := utils.PutRequest(url, *user, accessToken)
	if err != nil {
		return diag.FromErr(err)
	}
	resBody, _ := ioutil.ReadAll(res.Body)
	resUser := models.User{}
	err = json.Unmarshal(resBody, &resUser)
	if err != nil {
		return diag.FromErr(err)
	}
	// d.SetId(resUser.ID)
	return diags
}

func resourceTenantDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	accessToken := m.(map[string]string)["access_token"]
	url := m.(map[string]string)["baseUrl"] + "tenants/" + d.Id()
	err := utils.DeleteRequest(url, accessToken)
	if err != nil {
		return diag.FromErr(err)
	}
	return diags
}
