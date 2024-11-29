package repo

import (
	"a21hc3NpZ25tZW50/model"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type TeacherRepo struct {
	db *gorm.DB
}

// NewTeacherRepo: Konstruktor untuk TeacherRepo
func NewTeacherRepo(db *gorm.DB) TeacherRepo {
	return TeacherRepo{db}
}

// Save: Menyimpan data guru ke dalam tabel teachers
func (t TeacherRepo) Save(data model.Teacher) error {
	if err := t.db.Create(&data).Error; err != nil {
		return fmt.Errorf("failed to save teacher: %w", err)
	}
	return nil
}

// Query: Mengambil semua data guru yang belum dihapus (soft delete)
func (t TeacherRepo) Query() ([]model.Teacher, error) {
	var teachers []model.Teacher
	err := t.db.Select("*").Where("deleted_at IS NULL").Find(&teachers).Error
	if err != nil {
		return nil, fmt.Errorf("failed to query teachers: %w", err)
	}
	return teachers, nil
}

// Update: Memperbarui nama guru berdasarkan ID
func (t TeacherRepo) Update(id uint, name string) error {
	var teacher model.Teacher
	// Menyaring data guru berdasarkan ID
	if err := t.db.First(&teacher, id).Error; err != nil {
		return fmt.Errorf("teacher with id %d not found: %w", id, err)
	}

	// Update nama guru
	teacher.Name = name
	if err := t.db.Save(&teacher).Error; err != nil {
		return fmt.Errorf("failed to update teacher: %w", err)
	}
	return nil
}

// Delete: Melakukan soft delete pada data guru berdasarkan ID
func (t TeacherRepo) Delete(id uint) error {
	var teacher model.Teacher
	// Menyaring data guru berdasarkan ID
	if err := t.db.First(&teacher, id).Error; err != nil {
		return fmt.Errorf("teacher with id %d not found: %w", id, err)
	}

	// Melakukan soft delete dengan mengubah nilai DeletedAt
	teacher.Model.DeletedAt = gorm.DeletedAt{Time: time.Now(), Valid: true}
	if err := t.db.Save(&teacher).Error; err != nil {
		return fmt.Errorf("failed to delete teacher: %w", err)
	}

	return nil
}
