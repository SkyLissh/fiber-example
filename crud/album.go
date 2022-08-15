package crud

import (
	"skylissh/fiber-example/db"
	"skylissh/fiber-example/models"
	"skylissh/fiber-example/schemas"
)

type album struct {
	Base[models.Album, schemas.Album]
}

var Album = &album{Base[models.Album, schemas.Album]{DB: db.DB}}
