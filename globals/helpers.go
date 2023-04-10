package globals

func KBadRequest() map[string]interface{} {
	return map[string]interface{}{
		"success": false,
		"object":  nil,
		"message": "bad request",
	}
}

func KHelperErr() map[string]interface{} {
	return map[string]interface{}{
		"success": false,
		"object":  nil,
		"message": "helper error",
	}
}
func KEncoderDecoderErr() map[string]interface{} {
	return map[string]interface{}{
		"success": false,
		"object":  nil,
		"message": "Internal Server Error",
	}
}
