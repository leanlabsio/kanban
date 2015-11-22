package models

import (
	"encoding/json"
	"fmt"
	"gitlab.com/leanlabsio/kanban/modules/gitlab"
	"regexp"
	"strconv"
	"strings"
)

// Card represents an card in kanban board
type Card struct {
	Id          int64       `json:"id"`
	Iid         int64       `json:"iid"`
	Assignee    *User       `json:"assignee"`
	Milestone   *Milestone  `json:"milestone"`
	Author      *User       `json:"author"`
	Description string      `json:"description"`
	Labels      *[]string   `json:"labels"`
	ProjectId   int64       `json:"project_id"`
	Properties  *Properties `json:"properties"`
	State       string      `json:"state"`
	Title       string      `json:"title"`
	Todo        []*Todo     `json:"todo"`
}

// CardRequest represents a card request for create, update, delete card on kanban
type CardRequest struct {
	CardId      int64             `json:"issue_id"`
	ProjectId   int64             `json:"project_id"`
	Title       string            `json:"title"`
	Description string            `json:"description"`
	AssigneeId  *int64            `json:"assignee_id"`
	MilestoneId *int64            `json:"milestone_id"`
	Labels      string            `json:"labels"`
	Properties  *Properties       `json:"properties"`
	Stage       map[string]string `json:"stage"`
	Todo        []*Todo           `json:"todo"`
}

// Properties represents a card properties
type Properties struct {
	Andon string `json:"andon"`
}

// Todo represents an todo an card
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
		c := gitlab.NewContext(u.Credential["gitlab"].Token, u.Credential["gitlab"].PrivateToken)
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

// CreateCard create new card on board
func CreateCard(u *User, provider string, form *CardRequest) (*Card, int, error) {
	var cr *Card
	var code int
	switch provider {
	case "gitlab":
		c := gitlab.NewContext(u.Credential["gitlab"].Token, u.Credential["gitlab"].PrivateToken)
		r, res, err := c.CreateIssue(strconv.FormatInt(form.ProjectId, 10), mapCardRequestToGitlab(form))
		if err != nil {
			return nil, res.StatusCode, err
		}

		cr = mapCardFromGitlab(r)
	}

	return cr, code, nil
}

// UpdateCard updates existing card on board
func UpdateCard(u *User, provider string, form *CardRequest) (*Card, int, error) {
	var cr *Card
	var code int
	switch provider {
	case "gitlab":
		c := gitlab.NewContext(u.Credential["gitlab"].Token, u.Credential["gitlab"].PrivateToken)
		r, res, err := c.UpdateIssue(
			strconv.FormatInt(form.ProjectId, 10),
			strconv.FormatInt(form.CardId, 10),
			mapCardRequestToGitlab(form),
		)
		if err != nil {
			return nil, res.StatusCode, err
		}

		cr = mapCardFromGitlab(r)
	}

	return cr, code, nil
}

// DeleteCard removes card from board
func DeleteCard(u *User, provider string, form *CardRequest) (*Card, int, error) {
	var cr *Card
	var code int
	switch provider {
	case "gitlab":
		c := gitlab.NewContext(u.Credential["gitlab"].Token, u.Credential["gitlab"].PrivateToken)
		foru := mapCardRequestToGitlab(form)
		foru.StateEvent = "close"
		r, res, err := c.UpdateIssue(
			strconv.FormatInt(form.ProjectId, 10),
			strconv.FormatInt(form.CardId, 10),
			foru,
		)
		if err != nil {
			return nil, res.StatusCode, err
		}

		cr = mapCardFromGitlab(r)
	}

	return cr, code, nil
}

// mapCardRequestToGitlab transforms card to gitlab issue request
func mapCardRequestToGitlab(c *CardRequest) *gitlab.IssueRequest {
	return &gitlab.IssueRequest{
		Title:       c.Title,
		Description: mapCardDescriptionToGitlab(c.Description, c.Todo, c.Properties),
		AssigneeId:  c.AssigneeId,
		MilestoneId: c.MilestoneId,
		Labels:      c.Labels,
	}
}

// mapCardDescriptionToGitlab Transforms card description to gitlab description
func mapCardDescriptionToGitlab(desc string, t []*Todo, p *Properties) string {
	var d string
	var chek string
	d = strings.TrimSpace(desc)
	for _, v := range t {
		if v.Checked {
			chek = "x"
		} else {
			chek = " "
		}
		d = fmt.Sprintf("%s\n- [%s] %s", d, chek, v.Body)
	}

	pr, err := json.Marshal(p)

	if err == nil {
		d = fmt.Sprintf("%s\n\n<!-- @KB:%s -->", strings.TrimSpace(d), string(pr))
	}

	return strings.TrimSpace(d)
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
		Milestone:   mapMilestoneFromGitlab(c.Milestone),
		Labels:      removeDuplicates(c.Labels),
		ProjectId:   c.ProjectId,
		Properties:  mapCardPropertiesFromGitlab(c.Description),
		Todo:        mapCardTodoFromGitlab(c.Description),
	}
}

// removeDuplicates removed duplicates
func removeDuplicates(xs *[]string) *[]string {
	found := make(map[string]bool)
	j := 0
	for i, x := range *xs {
		if !found[x] {
			found[x] = true
			(*xs)[j] = (*xs)[i]
			j++
		}
	}
	*xs = (*xs)[:j]

	return xs
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
	} else {
		i = make([]*Todo, 0)
	}

	return i
}

func (c *Card) RoutingKey() string {
	return fmt.Sprintf("kanban.%d", c.ProjectId)
}

// mapCardDescriptionFromGitlab clears gitlab description to card description
func mapCardDescriptionFromGitlab(d string) string {
	var r string
	r = regTodo.ReplaceAllString(d, "")
	r = regProp.ReplaceAllString(r, "")
	return strings.TrimSpace(r)
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
