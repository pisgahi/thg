package hndl

import (
	"log"
	"net/http"
	"text/template"
	"time"

	"github.com/pisgahi/xcerpted-admin/lib"
	"github.com/pisgahi/xcerpted-admin/mdl"
)

var homeTmpl = template.Must(template.ParseFiles(
	"src/tmpl/home.html",
	"src/components/card.html"),
)

func HomePageHdlr(w http.ResponseWriter, r *http.Request) {

	createAt := time.Now()
	timetamp := lib.GetTimeStamp(createAt)

	var data = mdl.Card{
		PostId:    "1",
		Title:     "THG",
		Content:   "This is the thg stack.",
		TimeStamp: timetamp,
		CreatedAt: createAt,
	}

	err := homeTmpl.ExecuteTemplate(w, "home.html", data)
	if err != nil {
		log.Println("Error rendering home page:", err)
		http.Error(w, "Error rendering home page", http.StatusInternalServerError)
	}
}
