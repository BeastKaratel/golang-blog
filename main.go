package main

import (
	"Web_1/models"
	"fmt"
	"html/template"
	"net/http"
)

// PORT - is listening port
const PORT = "3000"

// HOST - is web host
const HOST = ":"

//All posts in var posts
var posts map[string]*models.Post

func indexHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/header.html", "templates/index.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
		fmt.Println(err.Error())
	} else {
		t.ExecuteTemplate(w, "index", posts)
	}
}

func writeHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/header.html", "templates/write.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
		fmt.Println(err.Error())
	} else {
		t.ExecuteTemplate(w, "write", nil)
	}
}

func editPostHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/header.html", "templates/write.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
		fmt.Println(err.Error())
	} else {
		id := r.FormValue("id")

		post, found := posts[id]
		if !found {
			http.NotFound(w, r)
		} else {
			t.ExecuteTemplate(w, "write", post)
		}
	}
}

func savePostHandler(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")
	title := r.FormValue("title")
	content := r.FormValue("content")
	var post *models.Post

	if id == "" {
		id = GenerateID()
		post = models.NewPost(id, title, content)
		posts[post.Id] = post
		fmt.Println("Add new post with ID =", id)
	} else {
		post = posts[id]
		post.Content = content
		post.Title = title
	}

	http.Redirect(w, r, "/", 302)
}

func deletePostHandler(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")
	if id == "" {
		http.NotFound(w, r)
	} else {
		delete(posts, id)
		http.Redirect(w, r, "/", 302)
	}
}

func main() {
	posts = make(map[string]*models.Post, 0)

	http.Handle("/bootstrap/", http.FileServer(http.Dir("./bootstrap/")))

	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/write", writeHandler)
	http.HandleFunc("/savePost", savePostHandler)
	http.HandleFunc("/edit", editPostHandler)
	http.HandleFunc("/delete", deletePostHandler)

	fmt.Println("Host has bean start on port", PORT)
	err := http.ListenAndServe(HOST+PORT, nil)
	if err != nil {
		fmt.Errorf(err.Error())
	}
}
