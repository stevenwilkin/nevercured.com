package main

import (
	"database/sql"
	"flag"
	"fmt"
	"html/template"
	"net/http"
	"os"
	"path"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

type Entry struct {
	Month   string
	Day     int
	Title   template.HTML
	Excerpt template.HTML
	Source  template.HTML
	Content template.HTML
	Summary template.HTML
}

var (
	db   *sql.DB
	port int
	tmpl *template.Template
)

func init() {
	flag.IntVar(&port, "p", 8080, "Port to listen on")
	flag.Parse()

	pwd, _ := os.Getwd()
	filename := path.Join(pwd, "templates", "index.tmpl")
	tmpl, _ = template.ParseFiles(filename)

	var err error
	db, err = sql.Open("sqlite3", "./db/nevercured.db")
	if err != nil {
		fmt.Println(err)
	}
}

func getEntry() Entry {
	sql := "SELECT * FROM jft WHERE month = ? AND day = ?"
	now := time.Now()

	var id int
	var month int
	var day int
	var title string
	var excerpt string
	var source string
	var content string
	var summary string

	err := db.QueryRow(sql, int(now.Month()), now.Day()).Scan(
		&id, &month, &day, &title, &excerpt, &source, &content, &summary)
	if err != nil {
		fmt.Println(err)
	}

	return Entry{
		Month:   time.Month(month).String(),
		Day:     day,
		Title:   template.HTML(title),
		Excerpt: template.HTML(excerpt),
		Source:  template.HTML(source),
		Content: template.HTML(content),
		Summary: template.HTML(summary)}
}

func handler(w http.ResponseWriter, r *http.Request) {
	tmpl.Execute(w, getEntry())
}

func main() {
	fmt.Printf("> Starting on http://0.0.0.0:%d\n", port)
	http.HandleFunc("/", handler)
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		fmt.Println("Error starting!")
	}
}
