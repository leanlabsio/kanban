package models

import (
	"encoding/json"
	"gitlab.com/kanban/kanban/modules/gitlab"
	"regexp"
)

type Card struct {
	Id          int64       `json:"id"`
	Iid         int64       `json:"iid"`
	Assignee    *User       `json:"assignee"`
	Author      *User       `json:"author"`
	Description string      `json:"description"`
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

var (
	regTodo = regexp.MustCompile(`[-\*]{1}\s(?P<checked>\[.\])(?P<body>.*)`)
	regProp = regexp.MustCompile(`<!--\s@KB:(.*?)\s-->`)
)

// ListCards returns list card
func ListCards(u *User, provider, board_id string) ([]*Card, error) {
	var b []*Card
	switch provider {
	case "gitlab":
		c := gitlab.NewContext(u.Credential["gitlab"].Token)
		op := &gitlab.IssueListOptions{
			State: "opened",
		}
		op.Page = "1"
		op.PerPage = "200"

		r, err := c.ListIssues(board_id, op)

		if err != nil {
			return nil, err
		}

		for _, d := range r {
			b = append(b, mapCardFromGitlab(d))
		}
	}

	return b, nil
}

// mapCardFromGitlab mapped gitlab issue to kanban card
func mapCardFromGitlab(c *gitlab.Issue) *Card {
	return &Card{
		Id:          c.Id,
		Iid:         c.Iid,
		Title:       c.Title,
		State:       c.State,
		Assignee:    mapUserFromGitlab(c.Assignee),
		Author:      mapUserFromGitlab(c.Author),
		Description: mapCardDescriptionFromGitlab(c.Description),
		Labels:      c.Labels,
		ProjectId:   c.ProjectId,
		Properties:  mapCardPropertiesFromGitlab(c.Description),
		Todo:        mapCardTodoFromGitlab(c.Description),
	}
}

// mapCardTodoFromGitlab tranforms gitlab todo to kanban todo
func mapCardTodoFromGitlab(d string) []*Todo {
	var i []*Todo
	m := regTodo.MatchString(d)

	if m {
		n := regTodo.SubexpNames()
		res := regTodo.FindAllStringSubmatch(d, -1)

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

			i = append(i, t)
		}
	}

	return i
}

// mapCardDescriptionFromGitlab clears gitlab description to card description
func mapCardDescriptionFromGitlab(d string) string {
	var r string
	r = regTodo.ReplaceAllString(d, "")
	r = regProp.ReplaceAllString(r, "")
	return r
}

// mapCardPropertiesFromGitlab transforms gitlab description to card properties
func mapCardPropertiesFromGitlab(d string) *Properties {
	m := regProp.MatchString(d)
	var p Properties

	if m {
		an := regProp.FindStringSubmatch(d)
		json.Unmarshal([]byte(an[1]), &p)
	}

	return &p
}
