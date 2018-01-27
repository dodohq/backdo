package helper

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/dodohq/backdo/models"
	"github.com/gorilla/schema"
)

// RenderJSON return json object in http response
func RenderJSON(w http.ResponseWriter, response []byte) {
	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}

// RenderErr return error in http response
func RenderErr(w http.ResponseWriter, err *models.HTTPError) {
	statusCode := err.StatusCode
	message, _ := json.Marshal(map[string]string{"message": err.Error()})
	http.Error(w, string(message), statusCode)
}

// ReadRequestBody read request body values and bind to the interface
func ReadRequestBody(r *http.Request, i interface{}) error {
	contentType := r.Header.Get("Content-Type")
	if contentType == "application/x-www-form-urlencoded" {
		if err := r.ParseForm(); err != nil {
			return err
		}
		decoder := schema.NewDecoder()
		if err := decoder.Decode(i, r.PostForm); err != nil {
			return err
		}
	} else if contentType == "application/json" {
		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(i); err != nil {
			return err
		}
	} else if strings.Contains(contentType, "multipart/form-data") {
		if err := r.ParseMultipartForm(32 << 20); err != nil {
			return err
		}
		decoder := schema.NewDecoder()
		if err := decoder.Decode(i, r.PostForm); err != nil {
			return err
		}
	} else {
		return errors.New("Content-Type Not Accepted")
	}
	return nil
}

// ReadFileUpload process file upload
// disk location on success
func ReadFileUpload(r *http.Request, fileFieldName string) (string, error) {
	file, _, err := r.FormFile(fileFieldName)
	if err != nil {
		return "", err
	}
	defer file.Close()

	filePath := "./tmp/" + strconv.FormatInt(time.Now().Unix(), 10)
	f, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0666)
	defer f.Close()
	if err != nil {
		return "", err
	}
	io.Copy(f, file)
	return filePath, nil
}
