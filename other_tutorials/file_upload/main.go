package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path"

	//"strings"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func UploadHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("inside upload func.")
	err := r.ParseMultipartForm(16777216) // 16MB grab the multipart form
	if err != nil {
		fmt.Fprintln(w, err)
		return
	}

	formdata := r.MultipartForm // ok, no problem so far, read the Form data

	//get the *fileheaders
	files := formdata.File["multiplefiles"] // grab the filenames

	for i, _ := range files { // loop through the files one by one
		file, err := files[i].Open()
		defer file.Close()
		if err != nil {
			fmt.Fprintln(w, err)
			return
		}
		pwd, _ := os.Getwd()
		out, err := os.Create(pwd + "/media/" + files[i].Filename)
		defer out.Close()
		if err != nil {
			fmt.Fprintf(w, "Unable to create the file for writing. Check your write access privilege")
			return
		}

		_, err = io.Copy(out, file) // file not files[i] !

		if err != nil {
			fmt.Fprintln(w, err)
			return
		}

		fmt.Fprintf(w, "Files uploaded successfully : ")
		fmt.Fprintf(w, files[i].Filename+"\n")

	}
}

func main() {
	router := NewRouter()
	// Serve static files from html dir.

	fs := http.FileServer(http.Dir("./media/"))
	//router.PathPrefix("/media/").Handler(fs)
	router.PathPrefix("/media/").Handler(http.StripPrefix("/media/", fs))
	//route := router.PathPrefix("/media").Handler(http.FileServer(http.Dir("./media/")))

	headers := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"})
	origins := handlers.AllowedOrigins([]string{"*"})

	//fmt.Println("Dir :" + http.Dir("/home/magesh/go/bin/html"))
	//router.PathPrefix("/html").Handler(http.FileServer(http.Dir("./html/")))
	//http.Handle("/html", http.FileServer(http.Dir("/home/magesh/go/bin/html")))
	log.Printf("Running in HTTP mode")
	log.Fatal(http.ListenAndServe(":8080", handlers.CORS(origins, headers, methods)(router)))

	//	http.HandleFunc("/upload", uploadHandler)
	//http.ListenAndServe(":8080", nil)
}

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
	Action      string
}

type Routes []Route

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {

		var handler http.Handler
		handler = route.HandlerFunc
		/*
			if strings.Contains(route.Action, "ValidationRequired") {
				handler = Validator(handler)
			}
			handler = Logger(handler, route.Name)
			handler = Recovery(handler)
		*/
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	return router
}

var routes = Routes{
	Route{
		"ChatLaunch",
		"GET",
		"/v2/html",
		ChatLaunch,
		//"SkipValidation",
		"ValidationRequired",
	},
	Route{
		"UPload",
		"POST",
		"/upload",
		UploadHandler,
		//"SkipValidation",
		"ValidationRequired",
	},
}

func ChatLaunch(w http.ResponseWriter, r *http.Request) {
	Filename := path.Base(r.URL.String())
	fmt.Printf("Attempting to serve :%v\n", Filename)
	//http.ServeFile(w, r, filepath.Join(".", "html", "index.html"))
	http.ServeFile(w, r, "./html/upload.html")
	//http.ServeFile(w, r, "index.html")
	fmt.Println("Done serving ", "upload.html")
}
