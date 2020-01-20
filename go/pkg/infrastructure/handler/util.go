package handler

import "strings"

func getParameterInPath(path string, prefix string, suffix string) string {
	param := strings.TrimPrefix(path, prefix)
	param = strings.TrimSuffix(param, suffix)
	return param
}
