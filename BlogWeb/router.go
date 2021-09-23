package main

import (
	"BlogPost/Post"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
)


func main(){
    newrouter := chi.NewRouter()

	newrouter.Get("/home", Post.GetBlog)
	newrouter.Get("/add", Post.AddBlog)
	newrouter.Post("/save", Post.SaveBlog)
	newrouter.Get("/view", Post.ViewBlog)

	newrouter.Route("/{id}", func(r chi.Router) {
		r.Get("/blogs/delete", Post.DeleteBlog)
		r.Get("/edit", Post.EditBlogs)
		r.Post("/blogs/update", Post.UpdateBlog)
	})


	log.Printf("Server Started at Localhost:1000")
	log.Fatalf("%v", http.ListenAndServe(":1000", newrouter))
}