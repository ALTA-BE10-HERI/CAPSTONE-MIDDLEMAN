package helper

func ResponseFailed(msg string) map[string]interface{} {
	return map[string]interface{}{
		"code":    404,
		"message": msg,
	}
}

func ResponseOkNoData(msg string) map[string]interface{} {
	return map[string]interface{}{
		"code":    200,
		"message": msg,
	}
}

func ResponseOkWithData(msg string, data interface{}) map[string]interface{} {
	return map[string]interface{}{
		"message": msg,
		"data":    data,
		"code":    200,
	}
}
