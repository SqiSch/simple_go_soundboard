package main

import (
    "github.com/ant0ine/go-json-rest/rest"
    "encoding/json"
    "log"
    "net/http"
    "net/url"
    "os/exec"
    "path/filepath"
    "os"
    "fmt"
    "strings"
)

type Message struct {
    Body string
}

func exe_cmd(title string) {
    cmd := "play"
    newTitle := fmt.Sprint("./sounds/favs/", title)
    args := []string{"-q", newTitle}
   err := exec.Command(cmd, args...).Run()
    if err != nil {
        log.Fatal(err)
    }
}


func list_songs()([]string, error) {
    searchDir := "sounds/favs/"

    fileList := []string{}
    filepath.Walk(searchDir, func(path string, f os.FileInfo, err error) error {
        path = strings.Replace(path, searchDir, "", -1)
        if len(path) > 0 {
          fileList = append(fileList, path)
        }
        return nil
    })

    js, _ := json.Marshal(fileList)
    fmt.Println(string(js))
   return fileList, nil
}


func main() {
    api := rest.NewApi()
    api.Use(rest.DefaultDevStack...)

    router, err := rest.MakeRouter(
        rest.Get("/play/#title", func(w rest.ResponseWriter, req *rest.Request) {
          dec_title, _ := url.QueryUnescape(req.PathParam("title"))
	          go exe_cmd(dec_title)
        }),
        rest.Get("/list/", func(w rest.ResponseWriter, req *rest.Request) {
  	      retValue, _ := list_songs()
              w.WriteJson(&retValue)
        }),
    )
    if err != nil {
        log.Fatal(err)
    }
    api.SetApp(router)



    http.Handle("/api/", http.StripPrefix("/api", api.MakeHandler()))
    http.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir("./web/"))))

    log.Fatal(http.ListenAndServe(":8080", nil))
}
