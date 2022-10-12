package belajargolangjson

import (
	"encoding/json"
	"fmt"
	"testing"
)

// data nya bsa dinamis, tidak bisa pake json tag
// tidak harus mendeklarasikan di struct
func TestMapDecode(t *testing.T) {
	jsonString := `{"id":"01","name":"Dell","image_url":"http://www.example.com/image.jpg"}`
	jsonBytes := []byte(jsonString)

	var result map[string]interface{}

	json.Unmarshal(jsonBytes, &result)

	fmt.Println(result)
	fmt.Println(result["id"])
}

func TestMapEncode(t *testing.T) {
	product := map[string]interface{}{
		"id":        "002",
		"name":      "dell",
		"image_url": "example.png",
	}

	bytes, _ := json.Marshal(product)

	fmt.Println(string(bytes))
}
