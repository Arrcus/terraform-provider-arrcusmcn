package schemas

import (
	"errors"

	"github.com/Arrcus/terraform-provider-arrcusmcn/models"
	"github.com/Arrcus/terraform-provider-arrcusmcn/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func UserSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"name": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		"username": &schema.Schema{
			Type:     schema.TypeString,
			Required: true,
		},
		"password": &schema.Schema{
			Type:     schema.TypeString,
			Required: true,
		},
		"email": &schema.Schema{
			Type:     schema.TypeString,
			Required: true,
		},
		"is_default": &schema.Schema{
			Type:     schema.TypeBool,
			Computed: true,
		},
		"is_default_password": &schema.Schema{
			Type:     schema.TypeBool,
			Computed: true,
		},
		"roles": &schema.Schema{
			Type:     schema.TypeList,
			Required: true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		},
	}
}

func ToUserSchema(user *models.User, d *schema.ResourceData) error {
	err := d.Set("name", *user.Name)
	if err != nil {
		return err
	}
	err = d.Set("username", *user.Username)
	if err != nil {
		return err
	}
	if user.Email != nil {
		err = d.Set("email", *user.Email)
		if err != nil {
			return err
		}
	}
	err = d.Set("password", "")
	if err != nil {
		return err
	}
	err = d.Set("is_default", user.IsDefault)
	if err != nil {
		return err
	}
	err = d.Set("is_default_password", user.IsDefaultPassword)
	if err != nil {
		return err
	}
	err = d.Set("roles", user.Roles)
	if err != nil {
		return err
	}
	return nil
}

func ToUserObj(d *schema.ResourceData) (*models.User, error) {
	user := models.User{}
	if v, exists := d.GetOk("name"); exists {
		user.Name = utils.StrPtr(v.(string))
	} else {
		return nil, errors.New("name is missing.")
	}
	if v, exists := d.GetOk("username"); exists {
		user.Username = utils.StrPtr(v.(string))
	} else {
		return nil, errors.New("username is missing.")
	}

	if v, exists := d.GetOk("email"); exists {
		user.Email = utils.StrPtr(v.(string))
	} else {
		return nil, errors.New("email is missing.")
	}

	if v, exists := d.GetOk("password"); exists {
		user.Password = utils.StrPtr(v.(string))
	} else {
		return nil, errors.New("password is missing.")
	}

	if v, exists := d.GetOk("is_default"); exists {
		vptr := v.(bool)
		user.IsDefault = &vptr
	} else {
		user.IsDefault = nil
	}

	if v, exists := d.GetOk("is_default_password"); exists {
		vptr := v.(bool)
		user.IsDefaultPassword = &vptr
	} else {
		user.IsDefaultPassword = nil
	}

	if v, exists := d.GetOk("roles"); exists {
		roleStrs := v.([]interface{})
		if len(roleStrs) == 0 {
			return nil, errors.New("roles can't be empty")
		}
		roles := make([]models.Rolename, 0)
		for _, role := range roleStrs {
			if r, err := checkRoleString(role.(string)); err != nil {
				return nil, err
			} else {
				roles = append(roles, *r)
			}
		}
		user.Roles = roles
	} else {
		return nil, errors.New("roles can't be empty")
	}
	return &user, nil
}

func checkRoleString(str string) (*models.Rolename, error) {
	switch str {
	case string(models.RolenameArcOrchAdmin):
		return models.RolenameArcOrchAdmin.Pointer(), nil
	case string(models.RolenameTenantAdmin):
		return models.RolenameTenantAdmin.Pointer(), nil
	case string(models.RolenameTenantOperator):
		return models.RolenameTenantOperator.Pointer(), nil
	}
	return nil, errors.New("invalid role name")
}
