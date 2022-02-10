package resource

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

func ResourceAwsDeployment() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceAwsDeploymentCreate,
		ReadContext:   resourceAwsDeploymentRead,
		UpdateContext: resourceAwsDeploymentUpdate,
		DeleteContext: resourceAwsDeploymentDelete,
		Schema:        schemas.AwsDeploymentSchema(),
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func resourceAwsDeploymentCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	accessToken := m.(map[string]string)["access_token"]
	url := m.(map[string]string)["baseUrl"] + "deployments"
	deployment := schemas.ToAwsDeploymentObj(d)
	res, err := utils.PostRequest(url, *deployment, accessToken)
	if err != nil {
		return diag.FromErr(err)
	}
	resBody, _ := ioutil.ReadAll(res.Body)
	resDeploy := models.Deployment{}
	err = json.Unmarshal(resBody, &resDeploy)
	if err != nil {
		return diag.FromErr(err)
	}
	err = schemas.UpdateCommonDeploymentResource(&resDeploy, d)
	if err != nil {
		return diag.FromErr(err)
	}
	err = schemas.UpdateAwsDeploymentResource(resDeploy.AwsDeployment, d)
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(resDeploy.ID)
	return diags
}

func resourceAwsDeploymentRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	accessToken := m.(map[string]string)["access_token"]
	url := m.(map[string]string)["baseUrl"] + "deployments/" + d.Id()
	res, err := utils.GetRequest(url, accessToken)

	if err != nil {
		return diag.FromErr(err)
	}

	resBody, err := ioutil.ReadAll(res.Body)
	if res.StatusCode != 200 {
		return diag.FromErr(errors.New(string(resBody)))
	}

	if err != nil {
		return diag.FromErr(err)
	}

	resDeploy := models.Deployment{}
	err = json.Unmarshal(resBody, &resDeploy)
	if err != nil {
		return diag.FromErr(err)
	}
	err = schemas.UpdateCommonDeploymentResource(&resDeploy, d)
	if err != nil {
		return diag.FromErr(err)
	}
	err = schemas.UpdateAwsDeploymentResource(resDeploy.AwsDeployment, d)
	if err != nil {
		return diag.FromErr(err)
	}
	return diags
}

func resourceAwsDeploymentUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	accessToken := m.(map[string]string)["access_token"]
	url := m.(map[string]string)["baseUrl"] + "deployments/" + d.Id()
	deployment := schemas.ToAwsDeploymentObj(d)
	res, err := utils.PutRequest(url, *deployment, accessToken)
	if err != nil {
		return diag.FromErr(err)
	}
	resBody, _ := ioutil.ReadAll(res.Body)
	resDeploy := models.Deployment{}
	err = json.Unmarshal(resBody, &resDeploy)
	if err != nil {
		return diag.FromErr(err)
	}
	err = schemas.UpdateCommonDeploymentResource(&resDeploy, d)
	if err != nil {
		return diag.FromErr(err)
	}
	err = schemas.UpdateAwsDeploymentResource(resDeploy.AwsDeployment, d)
	if err != nil {
		return diag.FromErr(err)
	}
	return diags
}

func resourceAwsDeploymentDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	accessToken := m.(map[string]string)["access_token"]
	url := m.(map[string]string)["baseUrl"] + "deployments"
	err := utils.DeleteRequest(url, d.Id(), accessToken, schemas.ToAwsDeploymentObj(d))
	if err != nil {
		return diag.FromErr(err)
	}
	return diags
}
