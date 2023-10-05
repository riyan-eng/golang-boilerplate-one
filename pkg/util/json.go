package util

import "encoding/json"

func UnmarshalConverter[T any](s string) (data T) {
	json.Unmarshal([]byte(s), &data)
	return
}
