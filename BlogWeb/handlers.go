package main

import (
	"BlogPost/Post"
	"database/sql"
	"fmt"
	"github.com/go-chi/chi/v5"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"strings"
)

const (
	TABLE = "BLOGS"
)

var DB *sql.DB

var templates *template.Template

func init() {
	templates = template.Must(template.ParseGlob("html/*.html"))

	DB, _ = sql.Open("mysql", "root:rootroot@tcp(127.0.0.1:3306)/blog")

	defer func(DB *sql.DB) {
		err := DB.Ping()
		if err != nil {
			DB.Close()
		}
	}(DB)
}


func AddBlog(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		_ = templates.ExecuteTemplate(w, "add.html",nil)
		return
	}

	_ = r.ParseForm()
	heading:= r.PostFormValue("heading")
	content:= r.PostFormValue("content")


	var err error
	if  strings.Contains(heading," ") || strings.Contains(content," ") {
		fmt.Println("Error inserting row:", err)
		_ = templates.ExecuteTemplate(w, "add.html", "Error adding post, please check all fields.")
		return
	}

	_ = templates.ExecuteTemplate(w, "add.html","Blog Successfully Added")
}

func ViewBlog(w http.ResponseWriter, r *http.Request) {
	stmt := fmt.Sprintf("SELECT id, heading, content FROM %s", TABLE)
	rows,err := DB.Query(stmt)
	if err != nil {
		log.Printf("Error: %v", err.Error())
		return
	}
	defer rows.Close()
	var blogs []Post.Blog

	for rows.Next(){
		var b Post.Blog
		err= rows.Scan(&b.Id,&b.Heading,&b.Content)
		if err != nil {
			log.Printf("Error: %v", err.Error())
			return
		}
		blogs=append(blogs,b)
	}
	err = templates.ExecuteTemplate(w, "view.html",blogs)
	if err != nil {
		log.Println(err)
	}
}

func GetBlog(w http.ResponseWriter, r *http.Request) {

	err := templates.ExecuteTemplate(w, "home.html", nil)
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
	content := r.PostFormValue("content")

	stmt := fmt.Sprintf("INSERT INTO %s (heading, content) VALUES(?, ?)", TABLE)
	ad, err := DB.Prepare(stmt)

	if err != nil {
		panic(err)
	}
	defer ad.Close()

	resp, err := ad.Exec(heading,content)

	rowsAffec, _ := resp.RowsAffected()
	if err != nil || rowsAffec != 1 {
		fmt.Println("Error adding row: ", err)
	}

	http.Redirect(w, r, "/home", 302)
}


func EditBlogs(w http.ResponseWriter, r *http.Request) {
	//Read value from route variable
	var Id = chi.URLParam(r, "id")
	id, _ := strconv.Atoi(Id)
	row := DB.QueryRow(fmt.Sprintf("SELECT * FROM %s WHERE id = ? ",TABLE),id)
	var post Post.Blog
	err := row.Scan(&post.Id, &post.Heading, &post.Content)
	if err != nil {
		fmt.Println(err)
		http.Redirect(w, r, "/view", 307)
		return
	}
	_ = templates.ExecuteTemplate(w, "edit.html", post)
}

func UpdateBlog(w http.ResponseWriter, r *http.Request) {
	//Read value from route variable
	vars := chi.URLParam(r,"id")
	id, _ := strconv.Atoi(vars)

	heading := r.FormValue("heading")
	content := r.FormValue("description")
	stm := fmt.Sprintf("UPDATE %s SET `heading`= ?, `content`= ? WHERE `id` = ? ",TABLE)

	upstm,_ := DB.Prepare(stm)

	defer upstm.Close()

	_, err := upstm.Exec(heading,content,id)
	if err != nil {
		fmt.Println(err)
		http.Redirect(w, r, "/view", 307)
		return
	}
	http.Redirect(w, r, "/view", 302)

}


func DeleteBlog(w http.ResponseWriter, r *http.Request) {
	//	//Read value from route variable
	var Id = chi.URLParam(r, "id")
	id, _ := strconv.Atoi(Id)

	stm := fmt.Sprintf("DELETE FROM %s WHERE id = ? ;",TABLE)
	del, err:= DB.Prepare(stm)

	if err != nil {
		log.Printf(err.Error())
	}
	defer del.Close()

	res, err := del.Exec(id)

	if err != nil {
		log.Printf(err.Error())
	}
	rowsAff, err := res.RowsAffected()

	fmt.Println("rowsAff: ", rowsAff)

	if err != nil || rowsAff != 1 {
		fmt.Fprint(w, "Error deleting product")
		return
	}

	http.Redirect(w, r,"/view", 302)

}
