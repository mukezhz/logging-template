package utils

import (
	"fmt"
	"strings"
)

func RenderTemplate(template string, vars map[string]any) string {
	msg := template
	for k, v := range vars {
		msg = strings.ReplaceAll(
			msg,
			"{"+k+"}",
			fmt.Sprintf("%v", v),
		)
	}
	return msg
}
