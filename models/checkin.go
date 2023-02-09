package models

import (
	"encoding/json"
	"time"

	"github.com/moody/helpers"
)

type Checkin struct {
	ID        string    `json:"id,omitempty" gorm:"primaryKey;type:char(36)"`
	Username  string    `json:"username,omitempty"`
	Location  string    `json:"location,omitempty"`
	CheckinAt time.Time `json:"checkin_at,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

func (Checkin) TableName() string {
	return "checkins"
}

func (o *Checkin) Schema() map[string]interface{} {
	return map[string]interface{}{
		"table": map[string]string{"name": "checkins", "as": "ck"},
		"fields": map[string]map[string]string{
			"id":         {"name": "ck.id", "as": "id", "type": "string"},
			"username":   {"name": "username", "as": "username", "type": "string"},
			"location":   {"name": "ck.location", "as": "location", "type": "string"},
			"checkin_at": {"name": "ck.checkin_at", "as": "checkin_at"},
			"created_at": {"name": "ck.created_at", "as": "created_at"},
			"updated_at": {"name": "ck.updated_at", "as": "updated_at"},
		},
	}
}

func (o *Checkin) GetById(ctx helpers.Context, id string, params map[string][]string) map[string]interface{} {
	return helpers.GetById(ctx, "checkins", "id", id, params, o.Schema(), map[string]interface{}{})
}

func (o *Checkin) GetPaginated(ctx helpers.Context, params map[string][]string) map[string]interface{} {
	return helpers.GetPaginated(ctx, params, o.Schema(), map[string]interface{}{})
}

func (o *Checkin) UpdateById(ctx helpers.Context) map[string]interface{} {
	helpers.GetDB(ctx).Model(Checkin{}).Where("id = ?", o.ID).Updates(o)
	return o.GetById(ctx, helpers.Convert(o.ID).String(), map[string][]string{})
}

func (o *Checkin) DeleteById(ctx helpers.Context) map[string]interface{} {
	id := helpers.Convert(o.ID).String()
	helpers.GetDB(ctx).Model(Checkin{}).Where("id = ?", o.ID).Delete(&Checkin{})
	return helpers.DeletedMessage("checkin", "id", id)
}

func (c *Checkin) SetDefaultValue(ctx helpers.Context) (map[string]interface{}, error) {
	params := map[string]interface{}{}
	c.ID = helpers.NewUUID()
	c.CheckinAt = time.Now()
	decoder := json.NewDecoder(ctx.Request().Body)
	err := decoder.Decode(c)
	if err != nil {
		return nil, err
	}
	return params, nil
}

func (o *Checkin) Create(ctx helpers.Context) map[string]interface{} {
	params, err := o.SetDefaultValue(ctx)
	if err != nil {
		return params
	}

	if err := helpers.GetDB(ctx).Create(o).Error; err != nil {
		return helpers.Map{
			"code":    500,
			"message": "Error while creating checkin: " + err.Error(),
		}
	}

	return helpers.Map{
		"code":    201,
		"message": "Checkin created successfully",
		"checkin": helpers.Map{
			"id":         o.ID,
			"username":   o.Username,
			"location":   o.Location,
			"checkin_at": o.CheckinAt,
			"created_at": o.CreatedAt,
			"updated_at": o.UpdatedAt,
		},
	}
}
