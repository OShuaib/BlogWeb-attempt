package Post

import (
	_ "encoding/json"
	_ "fmt"
	_ "github.com/go-chi/chi/v5"
	_ "io/ioutil"
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


//var Blogs = make(map[string]Blog)


