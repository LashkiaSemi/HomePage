package handler

import "strings"

import "strconv"

// getIntParameter パスの中に含まれるparamを取得する
func getIntParameter(path string, prefix string, suffix string) (int, error) {
	tmp := strings.TrimPrefix(path, prefix)
	tmp = strings.TrimSuffix(tmp, suffix)
	param, err := strconv.Atoi(tmp)
	if err != nil {
		return 0, err
	}
	return param, nil
}

// getParameterInPath パスの中に含まれるparamを取得する
func getParameterInPath(path string, prefix string, suffix string) string {
	param := strings.TrimPrefix(path, prefix)
	param = strings.TrimSuffix(param, suffix)
	return param
}
