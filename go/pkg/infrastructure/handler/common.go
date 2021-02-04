package handler

import (
	"fmt"
	"homepage/pkg/configs"
	"homepage/pkg/infrastructure/server/response"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/pkg/errors"
)

// createInfo 描画時の必須データを作成。どの画面であれ、絶対必要
func createInfo(r *http.Request, pageType, studentID string) *response.Info {
	return &response.Info{
		PageType:  pageType,
		StudentID: studentID,
	}
}

// createFileName ファイル名を出来るだけ一意にするためのもの
func createFileName(fileName string) string {
	return fmt.Sprintf("%s-%s", time.Now().Format(configs.DateForFileName), fileName)
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
func createFormField(name, value, label, formType string, options []*SelectFormOptions) *FormField {
	return &FormField{
		Name:    name,
		Value:   value,
		Label:   label,
		Type:    formType,
		Options: options,
	}
}

// FormField adminサイトのフォーム
type FormField struct {
	Name    string               // htmlのname
	Value   string               // 初期値
	Label   string               // フォームの表示名
	Type    string               // htmlのtype
	Options []*SelectFormOptions // select用
}

// SelectFormOptions formにselectを作成する
type SelectFormOptions struct {
	Value  string // value
	Label  string // 選択肢として表示する値
	Select bool
}

func createRoleOptions(role string) []*SelectFormOptions {
	return []*SelectFormOptions{
		&SelectFormOptions{
			Value:  "member",
			Label:  "member",
			Select: role == "member",
		},
		&SelectFormOptions{
			Value:  "admin",
			Label:  "admin",
			Select: role == "admin",
		},
		&SelectFormOptions{
			Value:  "owner",
			Label:  "owner",
			Select: role == "owner",
		},
	}
}

// createGradeOptions 初期値を設定しない場合、-1を指定してください
func createGradeOptions(grade int) []*SelectFormOptions {
	return []*SelectFormOptions{
		&SelectFormOptions{
			Value:  "2",
			Label:  "学部2年",
			Select: grade == 2,
		},
		&SelectFormOptions{
			Value:  "3",
			Label:  "学部3年",
			Select: grade == 3,
		},
		&SelectFormOptions{
			Value:  "4",
			Label:  "学部4年",
			Select: grade == 4,
		},
		&SelectFormOptions{
			Value:  "5",
			Label:  "大学院1年",
			Select: grade == 5,
		},
		&SelectFormOptions{
			Value:  "6",
			Label:  "大学院2年",
			Select: grade == 6,
		},
		&SelectFormOptions{
			Value:  "0",
			Label:  "卒業生",
			Select: grade == 0,
		},
	}

}
