package main

import (
	"fmt"
	"html/template" // Package for working with HTML templates
	"log"
	"net/http"
	"os"
	"strings"
)

// UserInput represents user input data.
type UserInput struct {
	UserText   string   // User's input text
	BannerType string   // Selected banner type
	OutputArr  []string // Array of generated ASCII art
}

var templates *template.Template

func init() {
	// Parse the HTML templates during initialization
	templates = template.Must(template.ParseGlob("templates/*.html"))
}

// WelcomeHandler handles requests to the root URL ("/").
func WelcomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		// If the URL path is not "/", handle it as a 404 error
		errorHandler(w, r, http.StatusNotFound)
		return
	}

	switch r.Method {
	case http.MethodGet:
		renderTemplate(w, r, "index.html", nil)
	case http.MethodPost:
		processForm(w, r)
	}
}

func processForm(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	userText := r.Form.Get("userText")     // Get the user's input text from the form
	userBanner := r.Form.Get("bannerType") // Get the selected banner type from the form
	
	if userText == "" {
		// If userText is empty, handle it as a 400 error
		errorHandler(w, r, http.StatusBadRequest)
		return
	}

	var bannerFile string
	switch userBanner {
	case "Standard":
		bannerFile = "standard.txt"
	case "Shadow":
		bannerFile = "shadow.txt"
	case "Thinkertoy":
		bannerFile = "thinkertoy.txt"
	default:
		bannerFile = "standard.txt"
	}

	// Read the banner file based on the selected banner type
	data, err := os.ReadFile(bannerFile)
	if err != nil {
		// If there's an error reading the banner file, handle it as a 500 error
		errorHandler(w, r, http.StatusInternalServerError)
		return
	}

	// Split the banner file into lines
	lines := strings.Split(string(data), "\n")
	// Split the user's input text into lines
	inputLines := strings.Split(userText, "\n")
	
	var asciiArr []string
	// Generate ASCII art for each line in the user's input text
	for _, line := range inputLines {
		words := strings.FieldsFunc(line, strSplit)
		for _, word := range words {
			if word == "" {
				continue
			}
			asciiArr = append(asciiArr, printWord(word, lines)...)
		}
	}

	myUser := UserInput{
		UserText:   userText,   // Set the UserText field of the UserInput struct to the userText variable
		BannerType: userBanner, // Set the BannerType field of the UserInput struct to the userBanner variable
		OutputArr:  asciiArr,   // Set the OutputArr field of the UserInput struct to the asciiArr variable
	}

	renderTemplate(w, r, "index.html", myUser) // Render the "index.html" template with myUser as the data
}

func renderTemplate(w http.ResponseWriter, r *http.Request, tmplName string, data interface{}) {
	// Parse and execute the specified template file with the provided data
	tmpl, err := template.ParseFiles("templates/" + tmplName)
	if err != nil {
		// If there's an error parsing the template, handle it as a 500 error
		tmpl, err = template.ParseFiles("templates/500.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		err = tmpl.Execute(w, nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		return
	}
	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func main() {

	// Register the WelcomeHandler function to handle requests to the root path "/"
	http.HandleFunc("/", WelcomeHandler)

	// Print a message indicating that the server is listening on port 8080
	fmt.Println("Server listening on :8000...")

	// Start the HTTP server on port 8080
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		// If an error occurs while starting the server, log the error and exit
		log.Fatal("Error starting server:", err)
	}
}

func printWord(word string, lines []string) []string {
	var strArray []string
	for j := 1; j < 9; j++ {
		str := ""
		// Iterate over each letter in the word
		for _, letter := range word {
			val := int(letter)
			// Calculate the line corresponding to the letter
			line := (val - 32) * 9
			// Check if the line is out of range
			if line+j >= len(lines) || line+j < 0 {
				// Break the loop if out of range to avoid accessing invalid indices
				break
			}
			// Append the line from the banner file to the current row of ASCII art
			str += lines[line+j]
		}
		// Append the row of ASCII art to the string array
		strArray = append(strArray, str)
	}
	// Return the generated ASCII art
	return strArray
}

func errorHandler(w http.ResponseWriter, r *http.Request, status int) {
	// Set the HTTP response status code
	w.WriteHeader(status)

	// Define the template filename based on the status code
	var templateFile string
	switch status {
	case http.StatusBadRequest:
		templateFile = "400.html"
	case http.StatusNotFound:
		templateFile = "404.html"
	default:
		templateFile = "500.html"
	}
	
	// Render the corresponding error template
	renderTemplate(w, r, templateFile, nil)
}

func strSplit(r rune) bool {
	return r == '\n' /*|| r == '\r'*/
}