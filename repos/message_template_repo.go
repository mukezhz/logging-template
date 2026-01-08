package repos

import (
	"github.com/mukezhz/logging-template/models"
	"gorm.io/gorm"
)

type TemplateRepo struct {
	db *gorm.DB
}

func NewTemplateRepo(db *gorm.DB) *TemplateRepo {
	return &TemplateRepo{db: db}
}

func (r *TemplateRepo) Get(
	tenantID, code string,
) (*models.MessageTemplate, error) {

	var tpl models.MessageTemplate

	err := r.db.
		Preload("Variables").
		Where("tenant_id = ? AND code = ? AND enabled = true",
			tenantID, code).
		First(&tpl).Error

	return &tpl, err
}
