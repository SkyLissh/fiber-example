package crud

import (
	"github.com/goccy/go-json"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Base[T any, E any] struct {
	DB *gorm.DB
}

func (b *Base[T, E]) decode(data *T, updated *E) error {
	u, e := json.Marshal(updated)

	if e != nil {
		return e
	}

	if err := json.Unmarshal(u, data); err != nil {
		return err
	}

	return nil
}

func (b *Base[T, E]) Find(data *[]T) chan error {
	c := make(chan error)

	go func() {
		if err := b.DB.Find(data).Error; err != nil {
			c <- err
		}

		close(c)
	}()

	return c
}

func (b *Base[T, E]) FindByID(id string, data *T) error {
	if err := b.DB.Where("id = ?", id).First(data).Error; err != nil {
		return err
	}

	return nil
}

func (b *Base[T, E]) Insert(data *T, updated *E) error {
	if err := b.decode(data, updated); err != nil {
		return err
	}
	
	if err := b.DB.Create(data).Error; err != nil {
		return err
	}
	
	return nil
}

func (b *Base[T, E]) Update(id string, data *T, updated *E) error {
	if err := b.decode(data, updated); err != nil {
		return err
	}

	if err := b.DB.Clauses(clause.Returning{}).Where("id = ?", id).Updates(*data).Error; err != nil {
		return err
	}

	return nil
}

func (b *Base[T, E]) Remove(id string, data *T) error {
	if err := b.DB.Where("id = ?", id).Delete(data).Error; err != nil {
		return err
	}

	return nil
}
