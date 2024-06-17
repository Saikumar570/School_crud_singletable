package repositories

import (
	"School_api/models"

	"github.com/jinzhu/gorm"
)

type SchoolRepository interface {
	FindAll() ([]models.School, error)
	FindByID(id uint) (models.School, error)
	Create(school models.School) (models.School, error)
	Update(school models.School) (models.School, error)
	Delete(school models.School) error
}

type schoolRepository struct {
	db *gorm.DB
}

func NewSchoolRepository(db *gorm.DB) SchoolRepository {
	return &schoolRepository{db}
}

func (r *schoolRepository) FindAll() ([]models.School, error) {
	var schools []models.School
	err := r.db.Find(&schools).Error
	return schools, err
}

func (r *schoolRepository) FindByID(id uint) (models.School, error) {
	var school models.School
	err := r.db.First(&school, id).Error
	return school, err
}

func (r *schoolRepository) Create(school models.School) (models.School, error) {
	//school.BeforeSave(r.db)
	err := r.db.Create(&school).Error
	return school, err
}

func (r *schoolRepository) Update(school models.School) (models.School, error) {
	err := r.db.Save(&school).Error
	return school, err
}

func (r *schoolRepository) Delete(school models.School) error {
	return r.db.Delete(&school).Error
}
