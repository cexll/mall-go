package json

import "github.com/goccy/go-json"

func JsonDecode(data string) (map[string]any, error) {
	var dat map[string]any
	err := json.Unmarshal([]byte(data), &dat)
	return dat, err
}

func JsonEncode(data any) (string, error) {
	jsons, err := json.Marshal(data)
	return string(jsons), err
}
