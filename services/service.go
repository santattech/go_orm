package services

import (
	"encoding/json"
	"log"
	"net/http"
	"post/models"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

var dbconn *gorm.DB

type Response struct {
	Data    []models.Post `json:"data"`
	Message string        `json:"message"`
}

func GetAllPosts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var posts = models.GetPosts()
	var resp Response

	err := dbconn.Find(&posts).Error

	if err == nil {
		log.Println(posts)
		resp.Data = posts
		resp.Message = "200"
		json.NewEncoder(w).Encode(resp)
	} else {
		log.Println(err)
		http.Error(w, err.Error(), 400)
	}
}

func GetPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	var resp Response
	var post = models.GetPost()
	log.Println(id)

	err := dbconn.Find(&post, "id = ?", id).Error

	if err == nil {
		log.Println(post)
		resp.Data = append(resp.Data, post)
		resp.Message = "200"
		json.NewEncoder(w).Encode(resp)
	} else {
		log.Println(err)
		http.Error(w, err.Error(), 400)
	}
}

func CreatePost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var resp Response
	var post = models.GetPost()
	_ = json.NewDecoder(r.Body).Decode(&post)
	log.Println(post)

	result := dbconn.Create(&post)
	err := result.Error

	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), 400)
		return
	}

	log.Println(result.RowsAffected, " rows created")

	resp.Message = "201"
	json.NewEncoder(w).Encode(resp)
}

func UpdatePost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var resp Response
	var post = models.GetPost()
	_ = json.NewDecoder(r.Body).Decode(&post)
	id, _ := strconv.Atoi(params["id"])

	err := dbconn.Model(&post).Where("id = ?", id).Updates(&post).Error

	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	dbconn.Find(&post, "id = ?", id)
	resp.Data = append(resp.Data, post)
	resp.Message = "200"
	json.NewEncoder(w).Encode(resp)
}

func DeletePost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	var resp Response
	var post = models.GetPost()
	err := dbconn.Delete(&post, id).Error

	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	resp.Message = "DELETED"
	json.NewEncoder(w).Encode(resp)
}

func SetDB(db *gorm.DB) {
	dbconn = db
	var post = models.GetPost()
	dbconn.AutoMigrate(&post)
}
