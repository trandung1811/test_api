package main

import (
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func getRequest(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello world"))
}
func getUploadFile(w http.ResponseWriter, r *http.Request) {
	 token:= r.Header.Get("x-access-token")
	 nameToken :=  token[0:5]
     if checkCurrentDir(nameToken) {
		removeContents(nameToken)
		 getFile(w,r,proxy_key, nameToken)
		 getFile(w,r,cookie_key, nameToken)
	 } else {
		 err:= os.Mkdir(nameToken, 0755)
		 if err != nil {
			 log.Fatal(err)
		 }
		 getFile(w,r,proxy_key, nameToken)
		 getFile(w,r,cookie_key, nameToken)
	 }

}

func getFile(w http.ResponseWriter, r *http.Request, key string, token string) {

	file, _, err := r.FormFile(key)

	if err != nil {
		panic(err)
	}
	defer file.Close()
	name:= ""
    if key == proxy_key {
    	name = "proxy.xlsx"
	}
	if key == cookie_key {
		name = "cookie.xlsx"
	}
	path:= filepath.Join(token,name)
	f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	_, _ = io.WriteString(w, "File "+key+" Uploaded successfully\n")
	_, _ = io.Copy(f, file)
}

func checkCurrentDir(token string) bool {
	file, err := os.Open(".")
	if err != nil {
		log.Fatalf("failed opening directory: %s", err)
	}
	defer file.Close()

	list,_ := file.Readdirnames(0) // 0 to read all files and folders
	for _, name := range list {
		if name == token {
			return true
		}
	}
	return false
}
func removeContents(dir string) error {
	d, err := os.Open(dir)
	if err != nil {
		return err
	}
	defer d.Close()
	names, err := d.Readdirnames(-1)
	if err != nil {
		return err
	}
	for _, name := range names {
		err = os.RemoveAll(filepath.Join(dir, name))
		if err != nil {
			return err
		}
	}
	return nil
}
