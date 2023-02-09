package models

import (
	"encoding/json"
	"time"

	"github.com/moody/helpers"
)

type Checkout struct {
	ID         string    `json:"id,omitempty" gorm:"primaryKey;type:char(36)"`
	Username   string    `json:"username,omitempty"`
	Location   string    `json:"location,omitempty"`
	CheckoutAt time.Time `json:"checkout_at,omitempty"`
	CreatedAt  time.Time `json:"created_at,omitempty"`
	UpdatedAt  time.Time `json:"updated_at,omitempty"`
}

func (Checkout) TableName() string {
	return "Checkout"
}

func (o *Checkout) Schema() map[string]interface{} {
	return map[string]interface{}{
		"table": map[string]string{"name": "checkouts", "as": "co"},
		"fields": map[string]map[string]string{
			"id":          {"name": "co.id", "as": "id", "type": "string"},
			"username":    {"name": "co.username", "as": "username", "type": "string"},
			"location":    {"name": "co.location", "as": "location", "type": "string"},
			"checkout_at": {"name": "co.checkout_at", "as": "checkout_at", "type": "time"},
			"created_at":  {"name": "co.created_at", "as": "created_at", "type": "time"},
			"updated_at":  {"name": "co.updated_at", "as": "updated_at", "type": "time"},
		},
	}
}

func (o *Checkout) GetById(ctx helpers.Context, id string, params map[string][]string) map[string]interface{} {
	return helpers.GetById(ctx, "checkouts", "id", id, params, o.Schema(), map[string]interface{}{})
}

func (o *Checkout) GetPaginated(ctx helpers.Context, params map[string][]string) map[string]interface{} {
	return helpers.GetPaginated(ctx, params, o.Schema(), map[string]interface{}{})
}

func (o *Checkout) UpdateById(ctx helpers.Context) map[string]interface{} {
	helpers.GetDB(ctx).Model(Checkout{}).Where("id = ?", o.ID).Updates(o)
	return o.GetById(ctx, helpers.Convert(o.ID).String(), map[string][]string{})
}

func (o *Checkout) DeleteById(ctx helpers.Context) map[string]interface{} {
	id := helpers.Convert(o.ID).String()
	helpers.GetDB(ctx).Model(Checkout{}).Where("id = ?", o.ID).Delete(&Checkout{})
	return helpers.DeletedMessage("checkout", "id", id)
}

func (c *Checkout) SetDefaultValue(ctx helpers.Context) (map[string]interface{}, error) {
	params := map[string]interface{}{}
	c.ID = helpers.NewUUID()
	decoder := json.NewDecoder(ctx.Request().Body)
	decoder.Decode(c)
	c.CheckoutAt = time.Now()
	return params, nil
}

func (o *Checkout) Create(ctx helpers.Context) map[string]interface{} {
	params, err := o.SetDefaultValue(ctx)
	if err != nil {
		return params
	}

	if err := helpers.GetDB(ctx).Create(o).Error; err != nil {
		return helpers.Map{
			"code":    500,
			"message": "Error while creating checkout: " + err.Error(),
		}
	}

	return helpers.Map{
		"code":    201,
		"message": "Checkout created successfully",
		"checkin": helpers.Map{
			"id":          o.ID,
			"username":    o.Username,
			"location":    o.Location,
			"checkout_at": o.CheckoutAt,
			"created_at":  o.CreatedAt,
			"updated_at":  o.UpdatedAt,
		},
	}
}
