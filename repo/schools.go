package repo

import (
	"a21hc3NpZ25tZW50/model"

	"gorm.io/gorm"
)

// SchoolRepo adalah struct untuk repository sekolah
type SchoolRepo struct {
	db *gorm.DB
}

// NewSchoolRepo constructor untuk SchoolRepo
func NewSchoolRepo(db *gorm.DB) SchoolRepo {
	return SchoolRepo{db}
}

// Init menyimpan data sekolah ke tabel sekolah
func (s SchoolRepo) Init(data []model.School) error {
	if err := s.db.Create(&data).Error; err != nil {
		return err
	}
	return nil
}
