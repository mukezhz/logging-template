package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/mukezhz/logging-template/dtos"
	"github.com/mukezhz/logging-template/repos"
	"github.com/mukezhz/logging-template/utils"
)

func ToastHandler(repo *repos.TemplateRepo) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tenantID := r.Header.Get("X-Tenant-ID")

		tpl, err := repo.Get(tenantID, "ORDER_LIMIT_EXCEEDED")
		if err != nil {
			http.Error(w, "template not found", 404)
			return
		}

		message := utils.RenderTemplate(
			tpl.Template,
			map[string]any{
				"orderId": "ORD-123",
				"limit":   5,
			},
		)

		resp := dtos.ToastResponse{
			Code:      tpl.Code,
			Message:   message,
			ToastType: tpl.ToastType,
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
	}
}
