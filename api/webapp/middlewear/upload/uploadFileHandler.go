package upload

import (
	"fmt"
	"io/ioutil"
	"microseviceAdmin/domain/store"
	"microseviceAdmin/webapp/session"
	"net/http"
	"os"
	"path/filepath"

	"github.com/julienschmidt/httprouter"
)

const maxUploadSize = 2 * 1024 * 1024 // 2 mb
const uploadPath = "./files"

func UploadFileHandler(s *store.Store) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

		session.CheckSession(w, r)
		err := session.CheckRigths(w, r)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			s.Logger.Errorf("Bad request. Err msg:%v. ", err)
			return
		}

		/*files := []string{
			"/api/webapp/tamplates/upload.html",
			"/api/webapp/tamplates/base.html",
		}

		tmpl, err := template.ParseFiles(files...)
		if err != nil {
			http.Error(w, err.Error(), 400)
			s.Logger.Errorf("Can not parse template: %v", err)
			return
		}

		err = tmpl.Execute(w, nil)
		if err != nil {
			http.Error(w, err.Error(), 400)
			s.Logger.Errorf("Can not parse template: %v", err)
			return
		}*/

		/*	if r.Method == "GET" {
			t, _ := template.ParseFiles("/api/webapp/tamplates/upload.html")
			t.Execute(w, nil)
			return
		}*/

		if err := r.ParseMultipartForm(maxUploadSize); err != nil {
			fmt.Printf("Could not parse multipart form: %v\n", err)
			w.WriteHeader(http.StatusBadRequest)
			s.Logger.Errorf("Cant pasrse from. Err msg:%v. ", err)
			return
		}

		// parse and validate file and post parameters
		file, fileHeader, err := r.FormFile("uploadFile")
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			s.Logger.Errorf("Invalid file. Err msg:%v. ", err)
		}
		defer file.Close()

		// Get and print out file size
		fileSize := fileHeader.Size
		fmt.Printf("File size (bytes): %v\n", fileSize)

		// validate file size
		if fileSize > maxUploadSize {
			w.WriteHeader(http.StatusBadRequest)
			s.Logger.Errorf("File is too big. Err msg:%v. ", err)
			return
		}

		fileBytes, err := ioutil.ReadAll(file)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			s.Logger.Errorf("Invalid file. Err msg:%v. ", err)
			return

		}

		// check file type, detectcontenttype only needs the first 512 bytes
		detectedFileType := http.DetectContentType(fileBytes)
		fileName := fileHeader.Filename

		newPath := filepath.Join(uploadPath, fileName)
		fmt.Printf("FileType: %s, File: %s\n", detectedFileType, newPath)

		// write file
		newFile, err := os.Create(newPath)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			s.Logger.Errorf("Cant write file. Err msg:%v. ", err)
			return
		}

		defer newFile.Close() // idempotent, okay to call twice
		if _, err := newFile.Write(fileBytes); err != nil || newFile.Close() != nil {
			w.WriteHeader(http.StatusBadRequest)
			s.Logger.Errorf("Cant write file. Err msg:%v. ", err)
			return
		}

		w.Write([]byte("SUCCESS"))
	}

}
