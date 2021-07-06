package web

import (
	"encoding/json"
	"fibonnaci-spiral/service"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type App struct {
	handlers map[string]http.HandlerFunc
}

func NewApp(cors bool) App {
	app := App{
		handlers: make(map[string]http.HandlerFunc),
	}
	spiralHandler := app.GetSpiral
	if !cors {
		spiralHandler = disableCors(spiralHandler)
	}
	app.handlers["/spiral"] = spiralHandler
	app.handlers["/"] = http.FileServer(http.Dir("/webapp")).ServeHTTP
	return app
}

func (a *App) Serve() error {
	for path, handler := range a.handlers {
		http.Handle(path, handler)
	}
	log.Println("Web server is available on port 8080")
	return http.ListenAndServe(":8080", nil)
}

func (a *App) GetSpiral(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	query := r.URL.Query()
	rowStr := query["rows"]
	rows := parseQueryParam(rowStr)
	colStr := query["cols"]
	cols := parseQueryParam(colStr)
	technologies, err := service.FibonnaciSpiral(rows, cols)
	if err != nil {
		sendErr(w, http.StatusInternalServerError, err.Error())
		return
	}
	err = json.NewEncoder(w).Encode(technologies)
	if err != nil {
		sendErr(w, http.StatusInternalServerError, err.Error())
	}
}

func sendErr(w http.ResponseWriter, code int, message string) {
	resp, _ := json.Marshal(map[string]string{"error": message})
	http.Error(w, string(resp), code)
}

// Needed in order to disable CORS for local development
func disableCors(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "*")
		w.Header().Set("Access-Control-Allow-Headers", "*")
		h(w, r)
	}
}

func parseQueryParam(s []string) int {
	aux := strings.Join(s, "")
	r, err := strconv.Atoi(aux)

	if err != nil {
		return 1
	}
	return r
}
