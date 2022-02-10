package resource

import (
	"context"
	"encoding/json"
	"io/ioutil"

	"github.com/Arrcus/terraform-provider-arrcusmcn/models"
	schemas "github.com/Arrcus/terraform-provider-arrcusmcn/schemas"
	"github.com/Arrcus/terraform-provider-arrcusmcn/utils"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ResourceUser() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceUserCreate,
		ReadContext:   resourceUserRead,
		UpdateContext: resourceUserUpdate,
		DeleteContext: resourceUserDelete,
		Schema:        schemas.UserSchema(),
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func resourceUserCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	tflog.Info(ctx, "Create Called")
	accessToken := m.(map[string]string)["access_token"]
	url := m.(map[string]string)["baseUrl"] + "register"
	user, err := schemas.ToUserObj(d)
	res, err := utils.PostRequest(url, *user, accessToken)
	if err != nil {
		return diag.FromErr(err)
	}
	resBody, _ := ioutil.ReadAll(res.Body)
	resUser := models.Credentials{}
	err = json.Unmarshal(resBody, &resUser)
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(resUser.ID)
	return diags
}

func resourceUserRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics
	accessToken := m.(map[string]string)["access_token"]
	url := m.(map[string]string)["baseUrl"] + "users/" + d.Id()
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

	resUser := models.User{}
	err = json.Unmarshal(resBody, &resUser)
	if err != nil {
		return diag.FromErr(err)
	}
	err = schemas.ToUserSchema(&resUser, d)
	if err != nil {
		return diag.FromErr(err)
	}
	return diags
}

func resourceUserUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	accessToken := m.(map[string]string)["access_token"]
	url := m.(map[string]string)["baseUrl"] + "users/" + d.Id()
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

func resourceUserDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	accessToken := m.(map[string]string)["access_token"]
	url := m.(map[string]string)["baseUrl"] + "users"
	// return diag.FromErr(errors.New(fmt.Sprint((d))))
	err := utils.DeleteRequest(url, d.Id(), accessToken, nil)
	if err != nil {
		return diag.FromErr(err)
	}
	return diags
}
