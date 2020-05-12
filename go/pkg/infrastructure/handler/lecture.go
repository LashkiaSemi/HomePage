package handler

import (
	"fmt"
	"homepage/pkg/configs"
	"homepage/pkg/infrastructure/auth"
	"homepage/pkg/infrastructure/server/response"
	"homepage/pkg/interface/controller"
	"homepage/pkg/interface/repository"
	"homepage/pkg/usecase/interactor"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
)

type lectureHandler struct {
	controller.LectureController
}

// LectureHandler レクチャーの入出力の受付
type LectureHandler interface {
	GetAll(w http.ResponseWriter, r *http.Request)
	Create(w http.ResponseWriter, r *http.Request)
	UpdateByID(w http.ResponseWriter, r *http.Request)
	DeleteByID(w http.ResponseWriter, r *http.Request)
}

// NewLectureHandler ハンドラの作成
func NewLectureHandler(sh repository.SQLHandler) LectureHandler {
	return &lectureHandler{
		LectureController: controller.NewLectureController(
			interactor.NewLectureInteractor(
				repository.NewLectureRepository(sh),
			),
		),
	}
}

func (lh *lectureHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	info := createInfo(r, "lecture", auth.GetStudentIDFromCookie(r))
	res, err := lh.LectureController.GetAll()
	if err != nil {
		response.InternalServerError(w, info)
		return
	}
	response.Success(w, "lecture/index.html", info, res)
}

func (lh *lectureHandler) Create(w http.ResponseWriter, r *http.Request) {
	info := createInfo(r, "create", auth.GetStudentIDFromCookie(r))
	body := struct {
		Title   string
		Comment string
	}{
		Title:   "",
		Comment: "",
	}

	if r.Method == "POST" {
		log.Println("lecture create: post")
		title := r.PostFormValue("title")
		comment := r.PostFormValue("comment")
		file, fileHeader, err := r.FormFile("file")
		if title == "" {
			info.Errors = append(info.Errors, "タイトルは必須です")
			response.Success(w, "lecture/edit.html", info, body)
			return
		}
		if err != nil {
			log.Println("lectureHandler: Create: ", err)
			info.Errors = append(info.Errors, "ファイルは必須です")
			response.Success(w, "lecture/edit.html", info, body)
			return
		}
		var activation int
		if r.PostFormValue("activation") == "public" {
			activation = 1
		} else {
			activation = 0
		}

		// TODO: savefile
		// かぶった時用に、名前帰るとかした方が良さげ?
		fileName := fileHeader.Filename
		var saveImage *os.File
		saveImage, err = os.Create(fmt.Sprintf("%s/%s", configs.SaveLectureFileDir, fileName))
		if err != nil {
			log.Println("lectureHandler: Create: createFile: ", err)
			// TODO: 驚き最小じゃない気がする
			response.InternalServerError(w, info)
			return
		}
		defer saveImage.Close()
		defer file.Close()
		_, err = io.Copy(saveImage, file)
		if err != nil {
			log.Println("lectureHandler: Create: save file: ", err)
			// 驚き最小じゃない気がする
			response.InternalServerError(w, info)
			return
		}
		log.Println("lectureHandler: Create: success save file: ", fileName)

		_, err = lh.LectureController.Create(info.StudentID, title, fileName, comment, activation)
		if err != nil {
			info.Errors = append(info.Errors, "作成失敗")
			response.Success(w, "lecture/edit.html", info, body)
			return
		}
		http.Redirect(w, r, "/lectures", http.StatusSeeOther)
	}
	response.Success(w, "lecture/edit.html", info, body)
}

func (lh *lectureHandler) UpdateByID(w http.ResponseWriter, r *http.Request) {
	info := createInfo(r, "lecture", auth.GetStudentIDFromCookie(r))
	lectureID, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		log.Println("lectureHandler: UpdateByID: failed to parse path parameter")
		response.InternalServerError(w, info)
		return
	}
	body, err := lh.LectureController.GetByID(lectureID)
	if err != nil {
		response.InternalServerError(w, info)
		return
	}
	// レクチャーの作者じゃないとき
	if body.Author.StudentID != info.StudentID {
		log.Println("lectureHandler: UpdateByID: dont have permission. requestUser: ", info.StudentID)
		http.Redirect(w, r, "/lectures", http.StatusSeeOther)
	}

	if r.Method == "POST" {
		log.Println("lecture update: post")
		title := r.PostFormValue("title")
		comment := r.PostFormValue("comment")
		var activation int
		if r.PostFormValue("activation") == "public" {
			activation = 1
		} else {
			activation = 0
		}
		if title == "" {
			info.Errors = append(info.Errors, "タイトルは必須です")
			response.Success(w, "lecture/edit.html", info, body)
			return
		}
		_, err := lh.LectureController.UpdateByID(lectureID, title, comment, activation)
		if err != nil {
			info.Errors = append(info.Errors, "更新失敗")
			response.Success(w, "lecture/edit.html", info, body)
			return
		}
		http.Redirect(w, r, "/lectures", http.StatusSeeOther)

	}
	response.Success(w, "lecture/edit.html", info, body)
}

func (lh *lectureHandler) DeleteByID(w http.ResponseWriter, r *http.Request) {
	info := createInfo(r, "lecture", auth.GetStudentIDFromCookie(r))
	lectureID, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		log.Println("lectureHandler: DeleteByID: failed to parse path parameter")
		response.InternalServerError(w, info)
		return
	}
	body, err := lh.LectureController.GetByID(lectureID)
	if err != nil {
		log.Println("lectureHandler: DeleteByID: failed to get lecture: ", err)
		response.InternalServerError(w, info)
		return
	}
	if body.Author.StudentID != info.StudentID {
		log.Println("lectureHandler: DeleteByID: call wrong user: studentID is", info.StudentID)
		http.Redirect(w, r, "/lectures", http.StatusSeeOther)
		return
	}

	if r.Method == "POST" {
		log.Println("lectureHandler: DeleteByID: post")
		err = lh.LectureController.DeleteByID(lectureID)
		if err != nil {
			log.Println("lectureHandler: DeleteByID: ", err)
			response.InternalServerError(w, info)
			return
		}
		http.Redirect(w, r, "/lectures", http.StatusSeeOther)
	}

	response.Success(w, "lecture/delete.html", info, body)
}
