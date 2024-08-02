package gcpfunction

import (
	"net/http"

	"github.com/ahmedsameha1/ccjsonparser/cmd/server"
)

var Serverr *server.Server

func init() {
	Serverr = server.New()
	Serverr.Run()
}

func ServeJSONParser(w http.ResponseWriter, r *http.Request) {
	Serverr.ServeHTTP(w, r)
}
