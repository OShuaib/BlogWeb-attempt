package Post

import (
	_ "encoding/json"
	_ "fmt"
	"github.com/go-chi/chi/v5"
	_ "github.com/go-chi/chi/v5"
	"html/template"
	_ "io/ioutil"
	"log"
	_ "log"
	"net/http"
	"strconv"
	_ "strconv"
	_ "strings"
	"time"
)

type Blog struct {
	Id int
	Heading string
	Content string
	CreationTime time.Time
}


type EditBlog struct {
	Blog
	Id string
}


var Blogs = make(map[string]Blog)

var templates *template.Template
//Compile view templates
// Must is a helper that wraps a call to a function returning (*Template, error)
// and panics if the error is non-nil. It is intended for use in variable initializations
// such as
//	var t = template.Must(template.New("name").Parse("html"))
func init() {
  templates = template.Must(template.ParseGlob("html/*.html"))
}


func AddBlog(w http.ResponseWriter, r *http.Request) {
	err := templates.ExecuteTemplate(w, "add.html",Blogs)
	if err != nil {
        log.Println(err)
	}
}

func ViewBlog(w http.ResponseWriter, r *http.Request) {
	err := templates.ExecuteTemplate(w, "view.html",Blogs)
	if err != nil {
		log.Println(err)
	}
}

func GetBlog(w http.ResponseWriter, r *http.Request) {
	err := templates.ExecuteTemplate(w, "home.html", Blogs)
	if err != nil {
		log.Println(err)
	}
}
func SaveBlog(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Fatal(err)
	}

	heading := r.PostFormValue("heading")
	con := r.PostFormValue("content")
	id := len(Blogs)+1
	note := Blog{id,heading, con, time.Now()}
	//increment the value of id for generating key for the map

	//convert id value to string
	k := strconv.Itoa(id)
	Blogs[k] = note

	http.Redirect(w, r, "/home", 302)
}

func EditBlogs(w http.ResponseWriter, r *http.Request) {

	var viewModel EditBlog
	//Read value from route variable
	vars := chi.URLParam(r,"id")
	if blog, ok := Blogs[vars]; ok {

		viewModel = EditBlog{blog, vars}
	}else {
		http.Error(w, "Could not find the resource to edit.", http.StatusBadRequest)
	}

	err := templates.ExecuteTemplate(w, "edit.html",viewModel)
	if err != nil {
		log.Println(err)
	}
}

func UpdateBlog(w http.ResponseWriter, r *http.Request) {
	//Read value from route variable
	vars := chi.URLParam(r,"id")
	var blogToUpd Blog
	if blog, ok := Blogs[vars]; ok {
		r.ParseForm()
		blogToUpd.Heading = r.PostFormValue("heading")
		blogToUpd.Content = r.PostFormValue("description")
		blogToUpd.CreationTime = blog.CreationTime
		Blogs[vars] = blogToUpd
	} else {
		http.Error(w, "Could not find the resource to update.", http.StatusBadRequest)
	}
	http.Redirect(w, r, "/view", 302)
}

func DeleteBlog(w http.ResponseWriter, r *http.Request) {
	//Read value from route variable
	vars := chi.URLParam(r, "id")
	// Remove from Store
	if _, ok := Blogs[vars]; ok {
		//delete existing item
		delete(Blogs, vars)

	} else {
		http.Error(w, "Could not find the resource to delete.", http.StatusBadRequest)
	}
	http.Redirect(w, r, "/view", http.StatusMovedPermanently)
}