package handlers

import (
	"fmt"
	"log"
	"net/http"
	"path/filepath"
	"sync"
	"time"

	"github.com/Shahriar-shudip/golang-microservies-tuitorial/File-server/fileStore"
	"github.com/gorilla/mux"
)

var wg sync.WaitGroup

type File struct {
	log   *log.Logger
	store fileStore.Storage
}

func NewFile(s fileStore.Storage, l *log.Logger) *File {
	return &File{log: l, store: s}
}

func (f *File) UploadRest(rw http.ResponseWriter, r *http.Request) {
	start := time.Now()
	vars := mux.Vars(r)

	id := vars["id"]
	filename := vars["filename"]
	wg.Add(1)
	go f.saveFile(id, filename, rw, r)

	elapsed := time.Since(start)
	fmt.Println(elapsed)
	wg.Wait()
}

func (f *File) saveFile(id, path string, rw http.ResponseWriter, r *http.Request) {
	fp := filepath.Join(id, path)
	err := f.store.Save(fp, r.Body)
	if err != nil {
		f.log.Println("Unable to save file", "error", err)
		http.Error(rw, "Unable to save file", http.StatusInternalServerError)
	}
}
