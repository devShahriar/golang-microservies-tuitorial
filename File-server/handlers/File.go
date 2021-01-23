package handlers

import (
	"log"
	"net/http"

	"github.com/Shahriar-shudip/golang-microservies-tuitorial/tree/main/File-server/fileStore"
)

type File struct {
	log   *log.Logger
	store fileStore.Storage
}

func NewFile(s fileStore.Storage, l *log.Logger) *File {
	return &File{log: l, store: s}
}

func (f *File) UploadRest(rw http.ResponseWriter, r *http.Request) {

}
