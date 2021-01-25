package utils

func CopyParams(source, target map[string]interface{}, keys ...string) {
	for _, key := range keys {
		if source[key] != nil {
			target[key] = source[key]
		}
	}
}

func GetPage(params map[string]interface{}) (pageNo, pageSize int) {
	pageNo = 1
	pageSize = 20
	if val, ok := params["pageNo"].(float64); ok {
		pageNo = int(val)
	}
	if val, ok := params["pageSize"].(float64); ok {
		pageSize = int(val)
	}
	return
}
