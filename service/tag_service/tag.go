package tag_service

import (
	"encoding/json"

	"github.com/simanchou/go-gin-example/models"
	"github.com/simanchou/go-gin-example/pkg/gredis"
	"github.com/simanchou/go-gin-example/pkg/logging"
	"github.com/simanchou/go-gin-example/service/cache_service"
)

// Tag tag struct
type Tag struct {
	ID         int
	Name       string
	CreatedBy  string
	ModifiedBy string
	State      int

	PageNum  int
	PageSize int
}

// ExistByName check exist by name
func (t *Tag) ExistByName() (bool, error) {
	return models.ExistTagByName(t.Name)
}

//ExistByID check exist by id
func (t *Tag) ExistByID() (bool, error) {
	return models.ExistTagByID(t.ID)
}

//Add add tag
func (t *Tag) Add() error {
	return models.AddTag(t.Name, t.State, t.CreatedBy)
}

// Edit edit tag
func (t *Tag) Edit() error {
	data := make(map[string]interface{})
	data["modified_by"] = t.ModifiedBy
	data["name"] = t.Name
	if t.State >= 0 {
		data["state"] = t.State
	}

	return models.EditTag(t.ID, data)
}

//Delete delete tag
func (t *Tag) Delete() error {
	return models.DeleteTag(t.ID)
}

//Count count all tags
func (t *Tag) Count() (int, error) {
	return models.GetTagTotal(t.getMaps())
}

//GetAll get all tags
func (t *Tag) GetAll() ([]models.Tag, error) {
	var (
		tags, cacheTags []models.Tag
	)

	cache := cache_service.Tag{
		State:    t.State,
		PageNum:  t.PageNum,
		PageSize: t.PageSize,
	}
	key := cache.GetTagsKey()
	if gredis.Exists(key) {
		data, err := gredis.Get(key)
		if err != nil {
			logging.Info(err)
		} else {
			json.Unmarshal(data, &cacheTags)
			return cacheTags, nil
		}
	}
	tags, err := models.GetTags(t.PageNum, t.PageSize, t.getMaps())
	if err != nil {
		return nil, err
	}

	gredis.Set(key, tags, 3600)
	return tags, nil
}

func (t *Tag) getMaps() map[string]interface{} {
	maps := make(map[string]interface{})
	maps["deleted_on"] = 0

	if t.Name != "" {
		maps["name"] = t.Name
	}
	if t.State >= 0 {
		maps["state"] = t.State
	}
	return maps
}
