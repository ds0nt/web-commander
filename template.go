package main

import (
  "sync"
  "net/http"
  "text/template"
  "path/filepath"
)

type templateHandler struct {
  once     sync.Once
  filename string
  templ    *template.Template
}

func (t *templateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
  t.once.Do(func() {
    t.templ = template.Must(template.ParseFiles(filepath.Join("templates", t.filename)))
  })
  t.templ.Execute(w, r)
}
