package schemas

import (
	"errors"

	"github.com/Arrcus/terraform-provider-arcorch/models"
	"github.com/Arrcus/terraform-provider-arcorch/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func UserSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"organization": &schema.Schema{
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
	}
}

func ToUserSchema(user *models.User, d *schema.ResourceData) error {
	err := d.Set("username", *user.Username)
	if err != nil {
		return err
	}
	err = d.Set("email", *user.Email)
	if err != nil {
		return err
	}
	err = d.Set("password", "")
	if err != nil {
		return err
	}
	err = d.Set("organization", user.Organization)
	if err != nil {
		return err
	}
	err = d.Set("is_default", user.IsDefault)
	if err != nil {
		return err
	}

	return nil
}

func ToUserObj(d *schema.ResourceData) (*models.User, error) {
	user := models.User{}
	if v, exists := d.GetOk("username"); exists {
		user.Name = utils.StrPtr(v.(string))
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

	user.Organization = *utils.StrPtr(d.Get("organization").(string))
	return &user, nil
}
