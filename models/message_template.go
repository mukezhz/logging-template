package models

import "time"

type MessageTemplate struct {
	ID        uint   `gorm:"primaryKey"`
	TenantID  string `gorm:"size:50;not null;index:idx_tenant_code,unique"`
	Code      string `gorm:"size:100;not null;index:idx_tenant_code,unique"`
	Template  string `gorm:"type:text;not null"`
	ToastType string `gorm:"size:20;not null"` // success, error, warning
	Enabled   bool   `gorm:"default:true"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
