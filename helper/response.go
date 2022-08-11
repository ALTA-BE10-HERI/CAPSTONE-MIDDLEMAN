package helper

func ResponseFailed(msg string) map[string]interface{} {
	return map[string]interface{}{
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
		"code":    200,
		"message": msg,
		"data":    data,
	}
}

func ResponseNoAccess(msg string) map[string]interface{} {
	return map[string]interface{}{
		"code":    401,
		"message": msg,
	}
}

func ResponseInternalServerError(msg string) map[string]interface{} {
	return map[string]interface{}{
		"code":    500,
		"message": msg,
	}
}

func ResponseBadRequest(msg string) map[string]interface{} {
	return map[string]interface{}{
		"code":    400,
		"message": msg,
	}
}

func ResponseCreate(msg string) map[string]interface{} {
	return map[string]interface{}{
		"code":    201,
		"message": msg,
	}
}
