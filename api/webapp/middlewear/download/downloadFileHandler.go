package download

import (
	"io"
	"microseviceAdmin/domain/store"
	"microseviceAdmin/webapp/middlewear"
	"microseviceAdmin/webapp/session"
	"net/http"
	"os"
	"strings"
)

func DownloadFileHandler(s *store.Store) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		session.CheckSession(w, r)
		err := session.CheckRigths(w, r)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			s.Logger.Errorf("Bad request. Err msg:%v. ", err)
			return
		}

		path, err := middlewear.CtxFile(r.Context())
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			s.Logger.Errorf("Cannot parse file: %v", err)
			return
		}

		file, err := os.Open(path)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			s.Logger.Errorf("Cannot parse file: %v", err)
			return
		}

		arr := strings.Split(file.Name(), "/")
		name := arr[len(arr)-1]

		w.Header().Set("Accept-ranges", "bytes")
		w.Header().Set("Content-Type", "application/octet-stream")
		w.Header().Set("Content-Disposition", "attachment; filename="+name+"")
		w.WriteHeader(http.StatusOK)

		_, err = io.Copy(w, file)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			s.Logger.Errorf("Cannot send file: %v", err)
			return
		}
	}
}
