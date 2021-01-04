package utils

import (
	"bytes"
	"encoding/json"
)

// ToBufferString 将对象转换为字节流
func ToBufferString(data interface{}) string {
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	if err := encoder.Encode(data); err != nil {
		return ""
	}
	return buffer.String()
}
