package main

import (
	"log"
	"net/http"
)

func indexPage(w http.ResponseWriter, r *http.Request) {
	// Render the main html page
	http.ServeFile(w, r, "static/index.html")
}

func projectsPage(w http.ResponseWriter, r *http.Request) {
	// Render the projects html page
	http.ServeFile(w, r, "static/projects.html")
}

func aboutPage(w http.ResponseWriter, r *http.Request) {
	// Render the about html page
	http.ServeFile(w, r, "static/about.html")
}

func contactPage(w http.ResponseWriter, r *http.Request) {
	// Render the contact html page
	http.ServeFile(w, r, "static/contact.html")
}

func main() {
	// Serve static files (images, CSS, JavaScript)
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// Route handlers for the different pages
	http.HandleFunc("/projects", projectsPage)
	http.HandleFunc("/about", aboutPage)
	http.HandleFunc("/contact", contactPage)
	http.HandleFunc("/", indexPage)

	// Start the server
	err := http.ListenAndServe("0.0.0.0:8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
