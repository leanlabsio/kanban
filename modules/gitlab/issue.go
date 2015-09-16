package gitlab

import (
	"encoding/json"
	"regexp"
)

type Issue struct {
	Assignee    *User       `json:"assignee"`
	Author      *User       `json:"author"`
	Description string      `json:"description"`
	Id          int64       `json:"id"`
	Iid         int64       `json:"iid"`
	Labels      []string    `json:"labels"`
	ProjectId   int64       `json:"project_id"`
	Properties  *Properties `json:"properties"`
	State       string      `json:"state"`
	Title       string      `json:"title"`
	Todo        []*Todo     `json:"todo"`
}

type Properties struct {
	Andon string `json:"andon"`
}

type Todo struct {
	Body    string `json:"body"`
	Checked bool   `json:"checked"`
}

type IssueListResponse struct {
	Data []Issue  `json:"data"`
	Meta []string `json:"meta"`
}

var regTodo = regexp.MustCompile(`[-\*]{1}\s(?P<checked>\[.\])(?P<body>.*)`)
var regProp = regexp.MustCompile(`<!--\s@KB:(.*?)\s-->`)

func (i *Issue) UnmarshalJSON(b []byte) error {
	type Alias Issue
	aux := (*Alias)(i)

	if err := json.Unmarshal(b, &aux); err != nil {
		return err
	}

	m := regTodo.MatchString(i.Description)

	if m {
		n := regTodo.SubexpNames()
		res := regTodo.FindAllStringSubmatch(i.Description, -1)
		i.Description = regTodo.ReplaceAllString(i.Description, "")

		for _, r1 := range res {
			t := &Todo{}
			for i, r2 := range r1 {
				switch n[i] {
				case "checked":
					if r2 == "[x]" {
						t.Checked = true
					} else {
						t.Checked = false
					}
				case "body":
					t.Body = r2
				}
			}

			i.Todo = append(i.Todo, t)
		}
	}

	m1 := regProp.MatchString(i.Description)

	if m1 {
		an := regProp.FindStringSubmatch(i.Description)
		i.Description = regProp.ReplaceAllString(i.Description, "")
		json.Unmarshal([]byte(an[1]), &i.Properties)
	}

	return nil
}
