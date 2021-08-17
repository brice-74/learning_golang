package main

import (
  "html/template"
  "path/filepath"
  "time"

  "brice.local/snippetbox/pkg/models"
  "brice.local/snippetbox/pkg/forms"
)

type templateData struct {
  CurrentYear int
  Form        *forms.Form
  Snippet     *models.Snippet
  Snippets    []*models.Snippet
}

// Create a humanDate function which returns a nicely formatted string
// representation of a time.Time object.
func humanDate(t time.Time) string {
  return t.Format(time.RFC822)
}

// Initialize a template.FuncMap object and store it in a global variable. This is
// essentially a string-keyed map which acts as a lookup between the names of our
// custom template functions and the functions themselves.
var functions = template.FuncMap{
  "humanDate": humanDate,
}

func newTemplateCache(dir string) (map[string]*template.Template, error) {

  cache := map[string]*template.Template{}

  pages, err := filepath.Glob(filepath.Join(dir + "page/", "*.html.tmpl"))
  if err != nil {
    return nil, err
  }

  for _, page := range pages {
    name := filepath.Base(page)

    ts, err := template.New(name).Funcs(functions).ParseFiles(page)
    if err != nil {
      return nil, err
    }

    ts, err = ts.ParseGlob(filepath.Join(dir + "layout/", "*.html.tmpl"))
    if err != nil {
      return nil, err
    }

    ts, err = ts.ParseGlob(filepath.Join(dir + "partial/", "*.html.tmpl"))
    if err != nil {
      return nil, err
    }

    cache[name] = ts
  }

  return cache, nil
}