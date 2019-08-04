package routes

import (
	"gopher/models"
	"net/http"
	"os"
    "strconv"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func Init() (h http.Handler) {
	r := mux.NewRouter()

	h = recoverWrap(r)
	h = handlers.LoggingHandler(os.Stdout, h)

	r.Path("/").
		Methods("GET").
		HandlerFunc(serveIndex)

	r.Path("/submit").
		Methods("GET").
		HandlerFunc(serveSubmit)

	r.Path("/submit").
		Methods("POST").
		HandlerFunc(jokeCreateHandler)

	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./assets/")))

	return
}

func serveIndex(w http.ResponseWriter, r *http.Request) {
	joke := models.Joke{}
	if err := joke.Random(); err != nil {
		errRes(w, r, 500, "Error looking up joke", err)
		return
	}
	if err := Tmpl.ExecuteTemplate(w, "index.html", indexPageData{
		Joke: joke,
	}); err != nil {
		errRes(w, r, 500, "Problem with template", err)
		return
	}
}

func serveSubmit(w http.ResponseWriter, r *http.Request) {
	if err := Tmpl.ExecuteTemplate(w, "submit.html", nil); err != nil {
		errRes(w, r, 500, "Problem with template", err)
		return
	}
}

type indexPageData struct {
	Joke models.Joke
}

func jokeCreateHandler(w http.ResponseWriter, r *http.Request) {

    score, err := strconv.ParseInt(r.FormValue("score"), 10, 8); 
    
    if err != nil {
        errRes(w, r, 400, "Problem with joke score.", err)
		return
    }
    if score < 1 {
        failRes(w, r, 400, "Your jokes score is in the negatives, we have never seen this before.", nil)
        return;
    }    
    if score < 5 {
        failRes(w, r, 400, "Your jokes score is not high enough.", nil)
        return;
    }    
    if score > 10 {
        failRes(w, r, 400, "Your jokes score is too high, is this a joke to you?", nil)
        return;
    }
    
    name := r.FormValue("name")
    body := r.FormValue("body")
    
    joke := models.Joke{name, body, int(score)}
    joke.Save()
    
    submitter := r.FormValue("submitter")
    
    sucRes(w, r, 302, name, body, submitter, int(score), nil)
	return
}
