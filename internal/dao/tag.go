package dao

import (
	"blog/internal/model"
	"blog/pkg/app"
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

// func (d *Dao) CreateTag(name string, state uint8, createdBy string) error {
func (d *Dao) CreateTag(name string, createdBy string) error {
	tag := model.Tag{
		Name:      name,
		Model:     &model.Model{State: 1},
		CreatedBy: createdBy,
	}
	return tag.Create(d.engine)
}

func (d *Dao) UpdateTag(id uint32, name string, state uint8, modifiedBy string) error {
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

func (d *Dao) DeleteTag(id uint32) error {
	tag := model.Tag{Model: &model.Model{ID: id}}
	return tag.Delete(d.engine)
}

func (d *Dao) GetTag(name, createdBy string) (int64, error) {
	// tag := model.Tag{Name: name, Model: &model.Model{CreatedBy: createdBy}}
	tag := model.Tag{Name: name, CreatedBy: createdBy}
	return tag.GetOne(d.engine)
}
