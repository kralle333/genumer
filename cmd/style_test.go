package cmd

import "testing"
import "github.com/stretchr/testify/require"

func TestEnumStyleApply(t *testing.T) {

	testValues := []string{"firstValue", "secondVal", "lastOne"}
	require.ElementsMatch(t, applyEnumStyle(testValues, enumStyleCamelCase), []string{"firstValue", "secondVal", "lastOne"})
	require.ElementsMatch(t, applyEnumStyle(testValues, enumStylePascalCase), []string{"FirstValue", "SecondVal", "LastOne"})
	require.ElementsMatch(t, applyEnumStyle(testValues, enumStyleUpperCase), []string{"FIRSTVALUE", "SECONDVAL", "LASTONE"})
	require.ElementsMatch(t, applyEnumStyle(testValues, enumStyleSnakeCase), []string{"first_value", "second_val", "last_one"})
	require.ElementsMatch(t, applyEnumStyle(testValues, enumStyleScreamingSnakeCase), []string{"FIRST_VALUE", "SECOND_VAL", "LAST_ONE"})
	require.ElementsMatch(t, applyEnumStyle(testValues, enumStyleKebabCase), []string{"first-value", "second-val", "last-one"})
	require.ElementsMatch(t, applyEnumStyle(testValues, enumStyleScreamingKebabCase), []string{"FIRST-VALUE", "SECOND-VAL", "LAST-ONE"})
}
