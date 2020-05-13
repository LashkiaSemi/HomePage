package handler

import (
	"fmt"
	"homepage/pkg/infrastructure/auth"
	"homepage/pkg/infrastructure/server/response"
	"homepage/pkg/interface/controller"
	"homepage/pkg/interface/repository"
	"homepage/pkg/usecase/interactor"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type equipmentHandler struct {
	controller.EquipmentController
	controller.TagController
}

// EquipmentHandler 備品関連の入出力を受け付け
type EquipmentHandler interface {
	GetAll(w http.ResponseWriter, r *http.Request)

	Create(w http.ResponseWriter, r *http.Request)
	UpdateByID(w http.ResponseWriter, r *http.Request)

	// admin
	AdminGetAll(w http.ResponseWriter, r *http.Request)
	AdminGetByID(w http.ResponseWriter, r *http.Request)
	AdminDeleteByID(w http.ResponseWriter, r *http.Request)
}

// NewEquipmentHandler ハンドラの作成
func NewEquipmentHandler(sh repository.SQLHandler) EquipmentHandler {
	return &equipmentHandler{
		EquipmentController: controller.NewEquipmentController(
			interactor.NewEquipmentInteractor(
				repository.NewEquipmentRepository(sh),
			),
		),
		TagController: controller.NewTagController(
			interactor.NewTagInteractor(
				repository.NewTagRepository(sh),
			),
		),
	}
}

func (eh *equipmentHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	info := createInfo(r, "equipment", auth.GetStudentIDFromCookie(r))
	res, err := eh.EquipmentController.GetAll()
	if err != nil {
		response.InternalServerError(w, info)
		return
	}
	response.Success(w, "equipment/index.html", info, res)
}

func (eh *equipmentHandler) Create(w http.ResponseWriter, r *http.Request) {
	info := createInfo(r, "equipments", auth.GetStudentIDFromCookie(r))

	// tagmap
	tags, err := eh.TagController.GetAll()
	if err != nil {
		log.Println("failed to get tags")
		response.InternalServerError(w, info)
		return
	}
	tagsMap := map[string]string{}
	for _, tag := range tags.Tags {
		tagsMap[strconv.Itoa(tag.ID)] = tag.Name
	}

	body := []*FormField{
		createFormField("name", "", "品名", "text", nil),
		createFormField("stock", "0", "在庫", "number", nil),
		createFormField("comment", "", "コメント", "textarea", nil),
		createFormField("tagID", "", "タグ", "select", tagsMap),
	}
	if r.Method == "POST" {
		log.Println("equipment create: post request")
		name := r.PostFormValue("name")
		comment := r.PostFormValue("comment")
		stock, err := strconv.Atoi(r.PostFormValue("stock"))
		if err != nil {
			// TODO: 然るべき処理
			log.Println("stock: failed to parse string to int")
			response.AdminRender(w, "edit.html", info, body)
			return
		}
		tagID, err := strconv.Atoi(r.PostFormValue("tagID"))
		if err != nil {
			// TODO: 然るべき処理
			log.Println("tagID: failed to parse string to int")
			tagID = 6
			// response.AdminRender(w, "edit.html", info, body)
			// return
		}
		if name == "" {
			info.Errors = append(info.Errors, "品名は必須です")
			response.AdminRender(w, "edit.html", info, body)
			return
		}
		id, err := eh.EquipmentController.Create(name, comment, stock, tagID)
		if err != nil {
			log.Println(err)
			response.InternalServerError(w, info)
			return
		}
		log.Println("success create!")
		http.Redirect(w, r, fmt.Sprintf("/admin/equipments/%d", id), http.StatusSeeOther)

	}
	response.AdminRender(w, "edit.html", info, body)
}

func (eh *equipmentHandler) UpdateByID(w http.ResponseWriter, r *http.Request) {
	info := createInfo(r, "equipments", auth.GetStudentIDFromCookie(r))
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		log.Println("failed to parse path parameter", err)
		response.InternalServerError(w, info)
		return
	}
	// get initial value
	data, err := eh.EquipmentController.GetByID(id)
	if err != nil {
		log.Println("failed to get target: ", err)
		response.InternalServerError(w, info)
		return
	}

	// tagmap
	tags, err := eh.TagController.GetAll()
	if err != nil {
		log.Println("failed to get tags")
		response.InternalServerError(w, info)
		return
	}
	tagsMap := map[string]string{}
	for _, tag := range tags.Tags {
		tagsMap[strconv.Itoa(tag.ID)] = tag.Name
	}

	// create form
	body := []*FormField{
		createFormField("name", data.Name, "品名", "text", nil),
		createFormField("stock", strconv.Itoa(data.Stock), "在庫", "number", nil),
		createFormField("comment", data.Comment, "コメント", "textarea", nil),
		createFormField("tagID", strconv.Itoa(data.Tag.ID), "タグ", "select", tagsMap), // TODO; 選択肢
	}
	if r.Method == "POST" {
		log.Println("equipment update: post request")
		name := r.PostFormValue("name")
		comment := r.PostFormValue("comment")
		stock, err := strconv.Atoi(r.PostFormValue("stock"))
		if err != nil {
			// TODO: 然るべき処理
			log.Println("stock: failed to parse string to int")
			response.AdminRender(w, "edit.html", info, body)
			return
		}
		tagID, err := strconv.Atoi(r.PostFormValue("tagID"))
		if err != nil {
			// TODO: 然るべき処理
			log.Println("tagID: failed to parse string to int")
			tagID = 6
			// response.AdminRender(w, "edit.html", info, body)
			// return
		}
		if name == "" {
			info.Errors = append(info.Errors, "品名は必須です")
			response.AdminRender(w, "edit.html", info, body)
			return
		}

		err = eh.EquipmentController.UpdateByID(id, name, comment, stock, tagID)
		if err != nil {
			log.Println(err)
			response.InternalServerError(w, info)
			return
		}
		log.Println("success update!")
		http.Redirect(w, r, fmt.Sprintf("/admin/equipments/%d", id), http.StatusSeeOther)

	}
	response.AdminRender(w, "edit.html", info, body)
}

func (eh *equipmentHandler) AdminGetAll(w http.ResponseWriter, r *http.Request) {
	info := createInfo(r, "equipments", auth.GetStudentIDFromCookie(r))
	res, err := eh.EquipmentController.AdminGetAll()
	if err != nil {
		log.Println("EquipmentHandler: ", err)
		response.InternalServerError(w, info)
		return
	}
	response.AdminRender(w, "list.html", info, res)
}

func (eh *equipmentHandler) AdminGetByID(w http.ResponseWriter, r *http.Request) {
	info := createInfo(r, "equipments", auth.GetStudentIDFromCookie(r))
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		log.Println("equipmentHandler: AdminGetByID: failed to parse path param: ", err)
		response.InternalServerError(w, info)
		return
	}
	res, err := eh.EquipmentController.AdminGetByID(id)
	if err != nil {
		log.Println("equipmentHandler: AdminGetByID: ", err)
		response.InternalServerError(w, info)
		return
	}
	res.ID = id
	response.AdminRender(w, "detail.html", info, res)
}

func (eh *equipmentHandler) AdminDeleteByID(w http.ResponseWriter, r *http.Request) {
	info := createInfo(r, "equipments", auth.GetStudentIDFromCookie(r))
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		log.Println("failed to parse path parameter", err)
		response.InternalServerError(w, info)
		return
	}
	body, err := eh.EquipmentController.AdminGetByID(id)
	if err != nil {
		log.Println("AdminDeleteByID: ", err)
		response.InternalServerError(w, info)
		return
	}

	if r.Method == "POST" {
		log.Println("post request: delete equipment")
		err = eh.EquipmentController.DeleteByID(id)
		if err != nil {
			log.Println("failed to delete")
			info.Errors = append(info.Errors, "削除に失敗しました")
			response.AdminRender(w, "delete.html", info, body)
			return
		}
		log.Println("success to delete equipment")
		http.Redirect(w, r, "/admin/equipments", http.StatusSeeOther)
	}
	response.AdminRender(w, "delete.html", info, body)
}
