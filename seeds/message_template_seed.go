package seeds

import (
	"github.com/mukezhz/logging-template/models"
	"gorm.io/gorm"
)

func Seed(db *gorm.DB) {
	data := []models.MessageTemplate{
		{
			TenantID:  "tenant_a",
			Code:      "ORDER_LIMIT_EXCEEDED",
			Template:  "Order {orderId} exceeded limit {limit}",
			ToastType: "error",
		},
		{
			TenantID:  "tenant_b",
			Code:      "ORDER_LIMIT_EXCEEDED",
			Template:  "You cannot place more than {limit} orders",
			ToastType: "warning",
		},
	}

	for _, d := range data {
		db.FirstOrCreate(&d, models.MessageTemplate{
			TenantID: d.TenantID,
			Code:     d.Code,
		})
	}
}
