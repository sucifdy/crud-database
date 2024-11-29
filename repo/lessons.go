package repo

import (
	"a21hc3NpZ25tZW50/model"

	"gorm.io/gorm"
)

// LessonRepo adalah struct untuk repository pelajaran
type LessonRepo struct {
	db *gorm.DB
}

// NewLessonRepo constructor untuk LessonRepo
func NewLessonRepo(db *gorm.DB) LessonRepo {
	return LessonRepo{db}
}

// Init menyimpan data pelajaran ke tabel pelajaran
func (l LessonRepo) Init(data []model.Lesson) error {
	if err := l.db.Create(&data).Error; err != nil {
		return err
	}
	return nil
}
