package core

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
)

// Respond ...
func Respond(w http.ResponseWriter, r *http.Request, status int, data interface{}) {
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(data); err != nil {
		log.Println("Erro encode:", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(status)
	if _, err := io.Copy(w, &buf); err != nil {
		log.Println("Erro copy buf:", err)
	}
	log.Println(r.URL, "status:", status)
}

// RespondFile ...
func RespondFile(w http.ResponseWriter, r *http.Request, status int, filename string, data []byte) {

	w.Header().Set("Content-Type", "text/plain")
	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=%s", filename))

	w.Write(data)

	log.Println(r.URL, "status:", status)
}

// RespondErro ...
func RespondErro(w http.ResponseWriter, r *http.Request, status int, errMsg *ErrMessage) {
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(errMsg); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(status)
	w.Write(buf.Bytes())
}

// handleZen ...
func handleZen(w http.ResponseWriter, r *http.Request) {
	data := SuccessMessage{
		Message: "Keep it logically awesome.",
	}
	Respond(w, r, http.StatusOK, data)
	return
}

// handleVersion ...
func handleVersion(w http.ResponseWriter, r *http.Request) {
	data := VersionMessage{
		AppID:          os.Getenv("HEROKU_APP_ID"),
		AppName:        os.Getenv("HEROKU_APP_NAME"),
		ServerID:       os.Getenv("HEROKU_DYNO_ID"),
		CreatedAt:      os.Getenv("HEROKU_RELEASE_CREATED_AT"),
		ReleaseVersion: os.Getenv("HEROKU_RELEASE_VERSION"),
		Commit:         os.Getenv("HEROKU_SLUG_COMMIT"),
		Description:    os.Getenv("HEROKU_SLUG_DESCRIPTION"),
	}
	Respond(w, r, http.StatusOK, data)
	return
}

// handleNotFound ...
func handleNotFound(w http.ResponseWriter, r *http.Request) {
	body := ErrMessage{Message: "URL não encontrada",
		Code: strconv.Itoa(http.StatusNotFound)}
	Respond(w, r, http.StatusNotFound, body)
	return
}
