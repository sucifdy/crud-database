package repo

import (
	"a21hc3NpZ25tZW50/model"

	"gorm.io/gorm"
)

// ClassRepo adalah struct untuk repository kelas
type ClassRepo struct {
	db *gorm.DB
}

// NewClassRepo constructor untuk ClassRepo
func NewClassRepo(db *gorm.DB) ClassRepo {
	return ClassRepo{db}
}

// Init menyimpan data kelas ke tabel kelas
func (c ClassRepo) Init(data []model.Class) error {
	if err := c.db.Create(&data).Error; err != nil {
		return err
	}
	return nil
}
