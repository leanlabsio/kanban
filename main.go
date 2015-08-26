package main

import (
	"github.com/shurcooL/github_flavored_markdown"
	"os"
	"io/ioutil"
	"log"
	"strings"
    "html/template"
)

type Page struct {
    Title string
    Url string
    Body  template.HTML
    Items []Page
}

func main() {
	files, _ := ioutil.ReadDir("kanban.wiki")
	var pages []Page
    for _, f := range files {
    	if strings.HasSuffix(f.Name(), ".md") && ! strings.HasPrefix(f.Name(), "_") {
    		markdown, err := ioutil.ReadFile("kanban.wiki/"+f.Name())

            if (err != nil) {
                log.Panic(err);
            }

    		title := strings.Replace(f.Name(), ".md", "", 1)
            var url string
            if title == "Home" {
                url = "documentation.html"
            } else {
                url = "documentation/" + title + ".html"  
            }

            page := Page{Title: title, Url:url, Body: template.HTML(github_flavored_markdown.Markdown(markdown))}
            pages = append(pages, page)
    	}
    }

    for _, p := range pages {
        p.Items = pages
        p.save()
    }
}

func (p *Page) save() error {
    file, err := os.Create(p.Url)
    if (err != nil) {
        log.Panic(err);
    }

    return renderTemplate(file, "documentation", p)
}

func renderTemplate(f *os.File, tmpl string, p *Page) error {
    t, err := template.ParseFiles(tmpl + ".tpl")

    if (err != nil) {
        log.Panic(err);
    }

    return t.Execute(f, p)
}