package schemas

import (
	"errors"

	"github.com/Arrcus/terraform-provider-arrcusmcn/models"
	"github.com/Arrcus/terraform-provider-arrcusmcn/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func TenantSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
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
	}
}

func ToTenantSchema(tenant *models.Tenant, d *schema.ResourceData) error {
	err := d.Set("name", *tenant.Name)
	if err != nil {
		return err
	}
	err = d.Set("organization", tenant.Organization)
	if err != nil {
		return err
	}
	err = d.Set("domain", *tenant.Domain)
	if err != nil {
		return err
	}
	err = d.Set("numdeployments", tenant.Numdeployments)
	if err != nil {
		return err
	}
	err = d.Set("numconnections", tenant.Numconnections)
	if err != nil {
		return err
	}
	err = d.Set("is_default", *tenant.IsDefault)
	if err != nil {
		return err
	}
	err = d.Set("defaultuser_name", *tenant.Defaultuser.Name)
	if err != nil {
		return err
	}
	err = d.Set("defaultuser_username", *tenant.Defaultuser.Username)
	if err != nil {
		return err
	}
	err = d.Set("defaultuser_email", *tenant.Defaultuser.Email)
	if err != nil {
		return err
	}
	err = d.Set("defaultuser_password", "")
	if err != nil {
		return err
	}
	err = d.Set("defaultuser_is_default", *tenant.Defaultuser.IsDefault)
	if err != nil {
		return err
	}
	err = d.Set("defaultuser_is_default_password", *tenant.Defaultuser.IsDefaultPassword)
	if err != nil {
		return err
	}
	err = d.Set("defaultuser_roles", tenant.Defaultuser.Roles)
	if err != nil {
		return err
	}
	return nil
}

func ToTenantObj(d *schema.ResourceData) (*models.Tenant, error) {
	tenant := models.Tenant{}
	tenant.Defaultuser = &models.User{}
	if v, exists := d.GetOk("name"); exists {
		tenant.Name = utils.StrPtr(v.(string))
	} else {
		return nil, errors.New("name is missing.")
	}
	if v, exists := d.GetOk("organization"); exists {
		tenant.Organization = v.(string)
	}

	if v, exists := d.GetOk("domain"); exists {
		tenant.Domain = utils.StrPtr(v.(string))
	} else {
		return nil, errors.New("domain is missing.")
	}

	if v, exists := d.GetOk("numdeployments"); exists {
		tenant.Numdeployments = v.(int64)
	}

	if v, exists := d.GetOk("numconnections"); exists {
		tenant.Numconnections = v.(int64)
	}

	if v, exists := d.GetOk("is_default"); exists {
		vptr := v.(bool)
		tenant.IsDefault = &vptr
	} else {
		tenant.IsDefault = nil
	}

	if v, exists := d.GetOk("defaultuser_name"); exists {
		tenant.Defaultuser.Name = utils.StrPtr(v.(string))
	} else {
		return nil, errors.New("defaultuser_name is missing.")
	}

	if v, exists := d.GetOk("defaultuser_username"); exists {
		tenant.Defaultuser.Username = utils.StrPtr(v.(string))
	} else {
		return nil, errors.New("defaultuser_username is missing.")
	}

	if v, exists := d.GetOk("defaultuser_password"); exists {
		tenant.Defaultuser.Password = utils.StrPtr(v.(string))
	} else {
		return nil, errors.New("defaultuser_password is missing.")
	}

	if v, exists := d.GetOk("defaultuser_email"); exists {
		tenant.Defaultuser.Email = utils.StrPtr(v.(string))
	} else {
		return nil, errors.New("defaultuser_email is missing.")
	}

	if v, exists := d.GetOk("defaultuser_name"); exists {
		tenant.Defaultuser.Username = utils.StrPtr(v.(string))
	} else {
		return nil, errors.New("defaultuser_name is missing.")
	}

	if v, exists := d.GetOk("defaultuser_is_default"); exists {
		vptr := v.(bool)
		tenant.Defaultuser.IsDefault = &vptr
	} else {
		tenant.Defaultuser.IsDefault = nil
	}

	if v, exists := d.GetOk("defaultuser_is_default_password"); exists {
		vptr := v.(bool)
		tenant.Defaultuser.IsDefaultPassword = &vptr
	} else {
		tenant.Defaultuser.IsDefaultPassword = nil
	}

	tenant.Defaultuser.Roles = []models.Rolename{models.RolenameTenantAdmin}
	// if v, exists := d.GetOk("defaultuser_roles"); exists {
	// 	roleStrs := v.([]interface{})
	// 	if len(roleStrs) == 0 {
	// 		return nil, errors.New("roles can't be empty")
	// 	}
	// 	roles := make([]models.Rolename, 0)
	// 	for _, role := range roleStrs {
	// 		if r, err := checkRoleString(role.(string)); err != nil {
	// 			return nil, err
	// 		} else {
	// 			roles = append(roles, *r)
	// 		}
	// 	}
	// 	tenant.Defaultuser.Roles = roles
	// } else {
	// 	return nil, errors.New("roles can't be empty")
	// }

	return &tenant, nil
}
