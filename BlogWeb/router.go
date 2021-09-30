package main

import (
	_ "database/sql/driver"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
)

func RegisterHandlers(newrouter *chi.Mux){
	newrouter.Use(middleware.Logger)

	newrouter.Get("/home", GetBlog)
	newrouter.Get("/add", AddBlog)
	newrouter.Post("/save", SaveBlog)
	newrouter.Get("/view", ViewBlog)

	newrouter.Route("/{id}", func(r chi.Router) {
		r.Get("/blogs/delete", DeleteBlog)
		r.Get("/edit", EditBlogs)
		r.Post("/update", UpdateBlog)
	})
}


func main(){
    newrouter := chi.NewRouter()

    RegisterHandlers(newrouter)


	log.Printf("Server Started at Localhost:1000")
	log.Fatalf("%v", http.ListenAndServe(":1000", newrouter))
}
