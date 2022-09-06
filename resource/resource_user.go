package resource

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/Arrcus/terraform-provider-arrcusmcn/models"
	schemas "github.com/Arrcus/terraform-provider-arrcusmcn/schemas"
	"github.com/Arrcus/terraform-provider-arrcusmcn/utils"
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
	tenant := m.(map[string]string)["tenant"]
	accessToken := m.(map[string]string)["access_token"]

	url := m.(map[string]string)["baseUrl"] + "users?tenant=" + tenant
	user, err := schemas.ToUserObj(d)
	if err != nil {
		return diag.FromErr(err)
	}
	res, err := utils.PostRequest(url, *user, accessToken)
	if err != nil {
		return diag.FromErr(err)
	}
	resBody, _ := ioutil.ReadAll(res.Body)
	resUser := models.User{}
	err = json.Unmarshal(resBody, &resUser)
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(resUser.ID.String())
	return diags
}

func resourceUserRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics
	tenant := m.(map[string]string)["tenant"]
	accessToken := m.(map[string]string)["access_token"]
	url := m.(map[string]string)["baseUrl"] + "users/" + d.Id() + "?tenant=" + tenant
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
	if d.HasChange("name") {
		return diag.Errorf(fmt.Sprintf("name can't be changed"))
	}
	if d.HasChange("username") {
		return diag.Errorf("username can't be changed")
	}
	tenant := m.(map[string]string)["tenant"]
	accessToken := m.(map[string]string)["access_token"]
	url := m.(map[string]string)["baseUrl"] + "users/" + d.Id() + "?tenant=" + tenant
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
	return diags
}

func resourceUserDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	tenant := m.(map[string]string)["tenant"]
	accessToken := m.(map[string]string)["access_token"]
	url := m.(map[string]string)["baseUrl"] + "users/" + d.Id() + "?tenant=" + tenant
	err := utils.DeleteRequest(url, accessToken)
	if err != nil {
		return diag.FromErr(err)
	}
	return diags
}
