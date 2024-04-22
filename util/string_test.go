package util

import "testing"

func TestCamelCaseToSplitChar(t *testing.T) {
	t.Log(CamelCaseToSplitChar("createUserNameAbC", '_'))

}

func TestSplitCharToCamelCase(t *testing.T) {
	t.Log(SplitCharToCamelCase("create_user_name_ab_c", '_'))
}
