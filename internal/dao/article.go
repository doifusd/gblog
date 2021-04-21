package dao

/*
func (d *Dao) CountArticle(name string, state uint8) (int64, error) {
	article := model.Article{Title: name, State: state}
	return article.Count(d.engine)
}

func (d *Dao) GetArticleList(name string, state uint8, page, pageSize int) ([]*model.Article, error) {
	article := model.Article{Title: name, State: state}
	pageOffset := app.GetPageOffset(page, pageSize)
	return article.List(d.engine, pageOffset, pageSize)
}

func (d *Dao) CreateArticle(name string, state uint8, createdBy string) error {
	article := model.Article{
		Title: name,
		State: state,
		Model: &model.Model{CreatedBy: createdBy},
	}
	return article.Create(d.engine)
}

func (d *Dao) UpdateArticle(id uint32, name string, state uint8, modifiedBy string) error {
	article := model.Article{
		Model: &model.Model{ID: id},
	}
	values := map[string]interface{}{
		"state":       state,
		"modified_by": modifiedBy,
	}
	if name != "" {
		values["name"] = name
	}
	return article.Update(d.engine, values)
}

func (d *Dao) DeleteArticle(id uint32) error {
	article := model.Article{Model: &model.Model{ID: id}}
	return article.Delete(d.engine)
}

func (d *Dao) GetArticle(name, createdBy string) (int64, error) {
	article := model.Article{Title: name, Model: &model.Model{CreatedBy: createdBy}}
	return article.GetOne(d.engine)
}*/
