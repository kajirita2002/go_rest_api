package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// User はユーザーです。
type User struct {
	FullName string `json:"fullname"`
	Username string `json:"username"`
	Email string `json:"email"`
}

// Post は投稿です
type Post struct {
	Title  string `json:"title"`
	Body   string `json:"body"`
	Author User `json:"author"`
}

var posts []Post = []Post{} // Itemのスライスを初期化

func main() {
	// ルーター作成
	router := mux.NewRouter()

	//　pathとhandlerを紐付けをする
	//  methodを決める
	router.HandleFunc("/posts", addItem).Methods("POST")

	router.HandleFunc("/posts", getAllPosts).Methods("GET")

	router.HandleFunc("/posts/{id}", getPost).Methods("GET")

	router.HandleFunc("/posts/{id}", updatePost).Methods("PUT")

	router.HandleFunc("/posts/{id}", patchPost).Methods("PATCH")

	router.HandleFunc("/posts/{id}",deletePost).Methods("DELETE")

	// ポート起動
	http.ListenAndServe(":5000", router)
}

func getPost(w http.ResponseWriter, r *http.Request) {
	// route parameterを取得する
	var idParam string = mux.Vars(r)["id"]
	id, err := strconv.Atoi(idParam)
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte("ID could not be converted to integer"))
		return
	}
	// エラーチェック
	if id >= len(posts) {
		w.WriteHeader(404)
		w.Write([]byte("No post found with specified ID"))
		return
	}

	post := posts[id]

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(post)
}

func getAllPosts(w http.ResponseWriter, r *http.Request) {
	// json型にResponseを変更する
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(posts)
}

func addItem(w http.ResponseWriter, r *http.Request) {
	var newPost Post
	// jsonでpostでリクエストされた値をnewItemに入力
	json.NewDecoder(r.Body).Decode(&newPost)

	posts = append(posts, newPost)
	w.Header().Set("Content-Type", "application/json")
	// jsonでdataを出力
	json.NewEncoder(w).Encode(posts)
}

func updatePost(w http.ResponseWriter, r *http.Request) {
	// get the ID of the post from the route parameters
	var idParam string = mux.Vars(r)["id"]
	id, err := strconv.Atoi(idParam)
	if (err != nil) {
		w.WriteHeader(400)
		w.Write([]byte("ID could not be converted to integer"))
		return
	}

	if id >= len(posts) {
		w.WriteHeader(404)
		w.Write([]byte("No post found with specified Id"))
		return
	}

	// get the value from JSON body
	var updatedPost Post
	json.NewDecoder(r.Body).Decode(&updatedPost)

	posts[id] = updatedPost

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updatedPost)
}

func patchPost(w http.ResponseWriter, r *http.Request) {
	// get the ID of the post from the route parameters
	var idParam string = mux.Vars(r)["id"]
	id, err := strconv.Atoi(idParam)
	if (err != nil) {
		w.WriteHeader(400)
		w.Write([]byte("ID could not be converted to integer"))
		return
	}

	if id >= len(posts) {
		w.WriteHeader(404)
		w.Write([]byte("No post found with specified Id"))
		return
	}

	post := &posts[id]
	json.NewDecoder(r.Body).Decode(post)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(post)
}

func deletePost(w http.ResponseWriter, r *http.Request) {
	// get the ID of the post from the route parameters
	var idParam string = mux.Vars(r)["id"]
	id, err := strconv.Atoi(idParam)
	if (err != nil) {
		w.WriteHeader(400)
		w.Write([]byte("ID could not be converted to integer"))
		return
	}

	if id >= len(posts) {
		w.WriteHeader(404)
		w.Write([]byte("No post found with specified Id"))
		return
	}

	// Delete the post from the slice
	posts = append(posts[:id], posts[id+1:]...)

	w.WriteHeader(200)
}
