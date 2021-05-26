package dao

import (
	"blog/internal/model"
	"blog/pkg/app"
	"time"
)

func (d *Dao) CountTag(name string, state uint8) (int64, error) {
	tag := model.Tag{Name: name, Model: &model.Model{State: state}}
	return tag.Count(d.engine)
}

func (d *Dao) GetTagList(name string, state uint8, page, pageSize int) ([]*model.Tag, error) {
	tag := model.Tag{Name: name, Model: &model.Model{State: state}}
	pageOffset := app.GetPageOffset(page, pageSize)
	return tag.List(d.engine, pageOffset, pageSize)
}

func (d *Dao) CreateTag(name string, createdBy uint32) error {
	tag := model.Tag{
		Name:      name,
		Model:     &model.Model{State: 1},
		CreatedBy: createdBy,
	}
	return tag.Create(d.engine)
}

func (d *Dao) UpdateTag(id uint32, name string, state uint8, modifiedBy uint32) error {
	tag := model.Tag{
		Model: &model.Model{ID: id},
	}
	values := map[string]interface{}{
		"state":       state,
		"modified_by": modifiedBy,
	}
	if name != "" {
		values["name"] = name
	}
	return tag.Update(d.engine, values)
}

func (d *Dao) DeleteTag(id uint32, modifiedBy uint32) error {
	now := time.Now().Format("2006-01-02 15:04:05")
	tag := model.Tag{Model: &model.Model{ID: id, State: 1}}
	values := map[string]interface{}{
		"state":       0,
		"modified_by": modifiedBy,
		"deleted_on":  now,
	}
	return tag.Delete(d.engine, values)
}

func (d *Dao) GetTag(name string, createdBy uint32) (int64, error) {
	// tag := model.Tag{Name: name, Model: &model.Model{CreatedBy: createdBy}}
	tag := model.Tag{Name: name, CreatedBy: createdBy}
	return tag.GetOne(d.engine)
}
