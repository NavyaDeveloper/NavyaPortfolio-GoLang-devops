package main

import (
	"log"
	"net/http"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	// Render the home html page from static folder
	http.ServeFile(w, r, "src/index.html")
}

func portfolioPage(w http.ResponseWriter, r *http.Request) {
	// Render the blog html page
	http.ServeFile(w, r, "src/portfolio.html")
}

func aboutPage(w http.ResponseWriter, r *http.Request) {
	// Render the about html page
	http.ServeFile(w, r, "src/about.html")
}

func contactPage(w http.ResponseWriter, r *http.Request) {
	// Render the contact html page
	http.ServeFile(w, r, "src/contact.html")
}

func certificationsPage(w http.ResponseWriter, r *http.Request) {
	// Render the certifications html page
	http.ServeFile(w, r, "src/services.html")
}

func resumePage(w http.ResponseWriter, r *http.Request) {
	// Render the resume html page
	http.ServeFile(w, r, "src/resume.html")
}
func main() {
	fs := http.FileServer(http.Dir("src/assets/"))
	http.Handle("/assets/", http.StripPrefix("/assets", fs))

	fs1 := http.FileServer(http.Dir("src/forms/"))
	http.Handle("/forms/", http.StripPrefix("/forms", fs1))

	http.HandleFunc("/home", homePage)
	http.HandleFunc("/portfolioPage", portfolioPage)
	http.HandleFunc("/about", aboutPage)
	http.HandleFunc("/contact", contactPage)
	http.HandleFunc("/certifications", certificationsPage)
	http.HandleFunc("/resume", resumePage)

	err := http.ListenAndServe("0.0.0.0:8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
