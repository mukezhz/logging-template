package seeds

import (
	"github.com/mukezhz/logging-template/models"
	"gorm.io/gorm"
)

func Seed(db *gorm.DB) {
	templates := []models.MessageTemplate{
		{
			TenantID:  "tenant_a",
			Code:      "ORDER_LIMIT_EXCEEDED",
			Template:  "Order {orderId} exceeded limit {limit}",
			ToastType: "error",
			Variables: []models.MessageTemplateVariable{
				{
					VariableKey: "orderId",
					Required:    true,
				},
				{
					VariableKey: "limit",
					Required:    true,
				},
			},
		},
	}

	for _, t := range templates {
		db.FirstOrCreate(&t, models.MessageTemplate{
			TenantID: t.TenantID,
			Code:     t.Code,
		})
	}
}
