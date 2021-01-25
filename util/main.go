package util

func DerefOptionalStringPtr(ptr *string) string {
	if ptr != nil {
		return *ptr
	}

	return ""
}
