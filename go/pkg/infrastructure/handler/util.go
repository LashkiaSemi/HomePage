package handler

import (
	"homepage/pkg/domain"
	"io"
	"mime/multipart"
	"os"
	"strconv"
	"strings"
)

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

// saveFile 指定した場所にfile(img, pdf etc...)を保存する
func saveFile(file multipart.File, path, fileName string) error {
	saveFile, err := os.Create(path + fileName)
	if err != nil {
		return domain.InternalServerError(err)
	}
	defer saveFile.Close()

	_, err = io.Copy(saveFile, file)
	return domain.InternalServerError(err)
}
