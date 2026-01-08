package models

import (
	"fmt"
	"time"
)

type MessageTemplate struct {
	ID        uint   `gorm:"primaryKey"`
	TenantID  string `gorm:"size:50;not null;index:idx_tenant_code,unique"`
	Code      string `gorm:"size:100;not null;index:idx_tenant_code,unique"`
	Template  string `gorm:"type:text;not null"`
	ToastType string `gorm:"size:20;not null"`
	Enabled   bool   `gorm:"default:true"`
	CreatedAt time.Time
	UpdatedAt time.Time

	Variables []MessageTemplateVariable `gorm:"foreignKey:TemplateID"`
}

type MessageTemplateVariable struct {
	ID           uint   `gorm:"primaryKey"`
	TemplateID   uint   `gorm:"index;not null"`
	VariableKey  string `gorm:"size:50;not null"`
	Required     bool
	DefaultValue string `gorm:"size:255"`
}

func ResolveVariables(
	defs []MessageTemplateVariable,
	input map[string]any,
) (map[string]any, error) {

	result := make(map[string]any)

	for _, d := range defs {
		val, ok := input[d.VariableKey]

		if !ok {
			if d.Required {
				return nil, fmt.Errorf(
					"missing required variable: %s",
					d.VariableKey,
				)
			}
			result[d.VariableKey] = d.DefaultValue
			continue
		}

		result[d.VariableKey] = val
	}

	return result, nil
}
