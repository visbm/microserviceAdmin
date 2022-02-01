package download

import (
	"microseviceAdmin/domain/store"
	"microseviceAdmin/webapp/middlewear"
	"net/http"
)

func DownloadFileHandler(s *store.Store) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		path, err := middlewear.CtxFile(r.Context())
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			s.Logger.Errorf("Cannot parse path: %v", err)
			return
		}
		fs := http.FileServer(http.Dir(path))
		http.StripPrefix("url", fs)
		http.StripPrefix("path", fs)
	}
}
