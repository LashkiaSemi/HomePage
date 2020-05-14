package handler

import (
	"fmt"
	"homepage/pkg/infrastructure/server/response"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"strings"

	"github.com/pkg/errors"
)

// createInfo 描画時の必須データを作成。どの画面であれ、絶対必要
func createInfo(r *http.Request, pageType, studentID string) *response.Info {
	return &response.Info{
		PageType:  pageType,
		StudentID: studentID,
	}
}

// saveFile 指定したパスにファイルを保存します
func saveFile(fileName, saveDir string, file multipart.File) error {
	saveDir = strings.TrimSuffix(saveDir, "/")
	var saveImage *os.File
	saveImage, err := os.Create(fmt.Sprintf("%s/%s", saveDir, fileName))
	if err != nil {
		err = errors.Wrap(err, "failed to reserve file")
		return err
	}
	defer saveImage.Close()
	_, err = io.Copy(saveImage, file)
	if err != nil {
		err = errors.Wrap(err, "failed to copy to reserve file")
		return err
	}
	return nil
}

// createFormField フォームのフィールドを一個作る
func createFormField(name, value, label, formType string, props map[string]string) *FormField {
	return &FormField{
		Name:  name,
		Value: value,
		Label: label,
		Type:  formType,
		Props: props,
	}
}

// FormField adminサイトのフォーム
type FormField struct {
	Name  string            // htmlのname
	Value string            // 初期値
	Label string            // フォームの表示名
	Type  string            // htmlのtype
	Props map[string]string // select用
}
