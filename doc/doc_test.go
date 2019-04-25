package doc

import (
	"reflect"
	"testing"

	_ "github.com/coveo/gotemplate/v3/hcl"
	_ "github.com/coveo/gotemplate/v3/json"
	_ "github.com/coveo/gotemplate/v3/yaml"
)

func Test_getDescStruct(t *testing.T) {
	tests := []struct {
		name  string
		desc  string
		want  interface{}
		want1 string
	}{
		{
			"YAML with default and description",
			`
				default: default_works
				description: description_works
			`, "default_works", "description_works",
		},
		{
			"YAML with description, no default",
			`
			description: description_works
			`, nil, "description_works",
		},
		{
			"Raw description no struct",
			`
			raw description works
			`, nil, "raw description works",
		},
		{
			"JSON with default and description",
			`
			{"default":"default_works", "description":"description_works"}
			`, "default_works", "description_works",
		},
		{
			"HCL with default and description",
			`
			default = "default_works"
			description =  "description_works"
		`, "default_works", "description_works",
		},
		{
			"YAML with default struct and description",
			`
			default:
				key1: value1
				key2: value2
			description: Struct default value
		`, map[string]interface{}{"key1": "value1", "key2": "value2"}, "Struct default value",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := getDescStruct(tt.desc)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getDescStruct() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("getDescStruct() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
