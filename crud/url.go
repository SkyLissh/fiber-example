package crud

import (
	"skylissh/fiber-example/db"
	"skylissh/fiber-example/models"
	"skylissh/fiber-example/schemas"
)

type urlCrud struct {
	Base[models.Url, schemas.Url]
}

func (u *urlCrud) FindByKey(key string, url *models.Url) error {
	if err := db.DB.Where("url = ?", key).First(url).Error; err != nil {
		return err
	}

	return nil
}

func (u *urlCrud) FindByTargetUrl(targetUrl string, url *models.Url) error {
	if err := db.DB.Where("target_url = ?", targetUrl).First(url).Error; err != nil {
		return err
	}

	return nil
}

var Url = &urlCrud{Base[models.Url, schemas.Url]{DB: db.DB}}
