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
    Body template.HTML
    Items []Page
    Sidebar template.HTML
}

const MdDir = "kanban.wiki"
const MdSidebar = MdDir + "/_Sidebar.md"
const HtmlDir = "documentation"
const HtmlTemplare = "build/templates/documentation.tpl"

func main() {
    var pages []Page
    var page Page
    var sidebar template.HTML

    files, err := ioutil.ReadDir(MdDir)

    if (err != nil) {
        log.Panic(err);
    }

    for _, f := range files {
        if strings.HasSuffix(f.Name(), ".md") && ! strings.HasPrefix(f.Name(), "_") {
            markdown, err := ioutil.ReadFile(MdDir + "/"+f.Name())
            log.Printf("%s generated", f.Name());

            if (err != nil) {
                log.Panic(err);
            }

            title := strings.Replace(f.Name(), ".md", "", 1)
            var url string
            if title == "Home" {
                url = HtmlDir + "/index.html"
            } else {
                url = HtmlDir + "/" + title + ".html"  
            }

            page = Page{}
            page.Title = title
            page.Url = url
            page.Body = template.HTML(github_flavored_markdown.Markdown(markdown))

            pages = append(pages, page)
        }
    }

    file, err := ioutil.ReadFile(MdSidebar);

    if err == nil {
        sidebar = template.HTML(github_flavored_markdown.Markdown(file))
    }

    for _, p := range pages {
        p.Items = pages
        p.Sidebar = sidebar
        p.save()
    }
}

func (p *Page) save() error {
    file, err := os.Create(p.Url)

    if (err != nil) {
        log.Panic(err);
    }

    return renderTemplate(file, p)
}

func renderTemplate(f *os.File, p *Page) error {
    t, err := template.ParseFiles(HtmlTemplare)

    if (err != nil) {
        log.Panic(err);
    }

    return t.Execute(f, p)
}