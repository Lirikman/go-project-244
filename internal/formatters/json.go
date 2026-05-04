package formatters

import (
	"encoding/json"
)

func FormmaterJson(tree map[string]map[string]any) string {
	jsonData, err := json.MarshalIndent(tree, "", "  ")
	if err != nil {
		return "serialization error"
	}
	return string(jsonData)
}
