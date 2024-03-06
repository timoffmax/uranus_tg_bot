package utils

func StringToBool(s string) bool {
	result := false

	if ("true" == s) || ("1" == s) {
		result = true
	}

	return result
}

func BoolToString(b bool) string {
	result := "false"

	if true == b {
		result = "true"
	}

	return result
}
