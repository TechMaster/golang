package demo

import (
	jsondefault "encoding/json"

	goccy "github.com/goccy/go-json"
)

func Goccy_JSON2Map(json string) (data map[string]interface{}) {
	_ = goccy.Unmarshal([]byte(json), &data)
	return data
}

func Goccy_JSON2Struct(json string) (person Person) {
	_ = goccy.Unmarshal([]byte(json), &person)
	return
}

func Goccy_Struct2JSON(person Person) string {
	if buff, err := goccy.Marshal(APerson); err == nil {
		return string(buff)
	} else {
		return ""
	}
}

//----------------------------------------------------------------

func Default_JSON2Map(json string) (data map[string]interface{}) {
	_ = jsondefault.Unmarshal([]byte(json), &data)
	return data
}

func Default_JSON2Struct(json string) (person Person) {
	_ = jsondefault.Unmarshal([]byte(json), &person)
	return
}

//----------------------------------------------------------------
func Jsoniter_JSON2Map(json string) (data map[string]interface{}) {
	_ = Jsoniter.Unmarshal([]byte(json), &data)
	return data
}

func Jsoniter_JSON2Struct(json string) (person Person) {
	_ = Jsoniter.Unmarshal([]byte(json), &person)
	return
}
