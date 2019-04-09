package print

import "testing"

func TestDefaultValue(t *testing.T) {
	descriptions := []string{
		`
			default: works
			description: yaml indented
		`,
		`
default: works
description: yaml unindented
		`,
		`
default =  works
description = hcl unindented
		`,
		`
{
	default: "works"
	description: "json unindented"
}
		`,
	}
	for _, d := range descriptions {
		v := defaultValue(d)
		t.Logf("Default value: %v\n", v)
	}
}
