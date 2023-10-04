package controllers

import (
	"api/src/db"
	"api/src/models"
	"api/src/repository"
	"api/src/response"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		response.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	var user models.User
	err = json.Unmarshal(requestBody, &user)
	if err != nil {
		response.Error(w, http.StatusBadRequest, err)
		return
	}

	err = user.Prepare()
	if err != nil {
		response.Error(w, http.StatusBadRequest, err)
		return
	}

	db, err := db.Connect()
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	userRepo := repository.NewUserRepo(db)
	userID, err := userRepo.Create(user)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}

	user.ID = userID
	response.JSON(w, http.StatusCreated, user)
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	nameOrNickname := strings.ToLower(r.URL.Query().Get("user"))

	db, err := db.Connect()
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err)
	}

	defer db.Close()

	repo := repository.NewUserRepo(db)
	users, err := repo.GetUsers(nameOrNickname)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err)
	}
	response.JSON(w, http.StatusOK, users)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	userID, err := strconv.ParseUint(params["userID"], 10, 64)
	if err != nil {
		response.Error(w, http.StatusBadRequest, err)
	}

	db, err := db.Connect()
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err)
	}
	repo := repository.NewUserRepo(db)
	user, err := repo.GetUser(userID)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err)
	}
	response.JSON(w, http.StatusOK, user)
	defer db.Close()

}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("update one user"))
}

func DeletUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("delete one user"))
}
