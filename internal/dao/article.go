package dao

import (
	"blog/internal/model"
	"blog/pkg/app"
	"time"
)

func (d *Dao) CountArticle(createdBy uint32, state uint8) (int64, error) {
	article := model.Article{CreatedBy: createdBy, Model: &model.Model{State: state}}
	return article.Count(d.engine)
}

func (d *Dao) GetArticleList(createdBy uint32, state uint8, page, pageSize int) ([]*model.Article, error) {
	article := model.Article{CreatedBy: createdBy, Model: &model.Model{State: state}}
	pageOffset := app.GetPageOffset(page, pageSize)
	return article.List(d.engine, pageOffset, pageSize)
}

func (d *Dao) GetArticleInfo(articleId uint32) (*model.Article, error) {
	article := model.Article{Model: &model.Model{ID: articleId}}
	return article.Info(d.engine, articleId)
}

func (d *Dao) CreateArticle(title string, desc string, content string, cover string, createdBy uint32, tagIds []uint32) error {
	article := model.Article{
		Model:     &model.Model{State: 1},
		Title:     title,
		Descs:     desc,
		Content:   content,
		Cover:     cover,
		CreatedBy: createdBy,
	}
	return article.Create(d.engine, tagIds, createdBy)
}

func (d *Dao) UpdateArticle(id uint32, title, descs, content, cover string, state uint8, modifiedBy uint32, tags []uint32) error {
	article := model.Article{
		Model: &model.Model{ID: id},
	}
	values := map[string]interface{}{
		"modified_by": modifiedBy,
	}
	if state == 1 {
		values["state"] = 1
	} else if state == 2 {
		values["state"] = 0
	}

	if title != "" {
		values["title"] = title
	}
	if content != "" {
		values["content"] = content
	}
	if descs != "" {
		values["descs"] = descs
	}
	if cover != "" {
		values["cover"] = cover
	}
	return article.Update(d.engine, values, tags)
}

func (d *Dao) DeleteArticle(id uint32, modifiedBy uint32) error {
	article := model.Article{Model: &model.Model{ID: id}}

	values := map[string]interface{}{
		"modified_by": modifiedBy,
		"state":       0,
		"deleted_on":  time.Now().Format("2006-01-02 15:04:05"),
	}
	return article.Delete(d.engine, values)
}
