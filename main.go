package main

import(
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main(){

	router := httprouter.New()
	router.GET("/", HomeHandler)

	// Posts collection
    router.GET("/posts", PostsIndexHandler)
    router.POST("/posts", PostsCreateHandler)

    // Posts singular
    router.GET("/posts/:id", PostShowHandler)
    router.PUT("/posts/:id", PostUpdateHandler)
    router.GET("/posts/:id/edit", PostEditHandler)

    fmt.Println("Starting server on :8080")
    http.ListenAndServe(":8080", router)

}

func HomeHandler(rw http.ResponseWriter, r *http.Request, p httprouter.Params){
	fmt.Fprintln(rw, "Home")
}

func PostsIndexHandler(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
    fmt.Fprintln(rw, "posts index")
}

func PostsCreateHandler(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
    fmt.Fprintln(rw, "posts create")
}

func PostShowHandler(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
    id := p.ByName("id")
    fmt.Fprintln(rw, "showing post", id)
}

func PostUpdateHandler(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
    fmt.Fprintln(rw, "post update")
}

func PostDeleteHandler(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
    fmt.Fprintln(rw, "post delete")
}

func PostEditHandler(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
    fmt.Fprintln(rw, "post edit")
}