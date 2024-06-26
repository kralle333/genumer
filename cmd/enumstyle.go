// Code generated by enumgen; DO NOT EDIT
package cmd

import (
	"encoding/json"
	"fmt"
)

type enumStyle string

const (
	enumStyleCamelCase          enumStyle = "camelCase"
	enumStylePascalCase         enumStyle = "pascalCase"
	enumStyleUpperCase          enumStyle = "upperCase"
	enumStyleKebabCase          enumStyle = "kebabCase"
	enumStyleScreamingKebabCase enumStyle = "screamingKebabCase"
	enumStyleSnakeCase          enumStyle = "snakeCase"
	enumStyleScreamingSnakeCase enumStyle = "screamingSnakeCase"
)

var allEnumStyles = []enumStyle{
	enumStyleCamelCase,
	enumStylePascalCase,
	enumStyleUpperCase,
	enumStyleKebabCase,
	enumStyleScreamingKebabCase,
	enumStyleSnakeCase,
	enumStyleScreamingSnakeCase,
}

type enumStyleObj struct{}

var enumStyleMap = map[enumStyle]enumStyleObj{
	enumStyleCamelCase:          {},
	enumStylePascalCase:         {},
	enumStyleUpperCase:          {},
	enumStyleKebabCase:          {},
	enumStyleScreamingKebabCase: {},
	enumStyleSnakeCase:          {},
	enumStyleScreamingSnakeCase: {},
}

func (e enumStyle) MarshalJSON() ([]byte, error) {
	return json.Marshal(e.String())
}

func (e *enumStyle) UnmarshalJSON(bytes []byte) error {
	var name string
	err := json.Unmarshal(bytes, &name)
	if err != nil {
		return err
	}
	val, err := enumStyleFromString(name)
	if err != nil {
		return err
	}
	*e = val

	return nil
}
func enumStyleFromString(v string) (enumStyle, error) {
	for _, known := range allEnumStyles {
		if v == known.String() {
			return known, nil
		}
	}
	return "", fmt.Errorf("unable to find enumStyle with value %s", v)
}

func (e enumStyle) String() string {
	return string(e)
}

func (e enumStyle) IsValid() bool {
	_, ok := enumStyleMap[e]
	return ok
}

func (e enumStyle) IsGeneratedGoEnum() bool {
	return true
}
