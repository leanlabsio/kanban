package gitlab

import (
	"encoding/json"
	"regexp"
	"strings"
	"time"
)

type Comment struct {
	Id        int64     `json:"id"`
	Author    *User     `json:"author"`
	Body      string    `json:"body"`
	CreatedAt time.Time `json:"created_at"`
	IsInfo    bool      `json:"is_info"`
}

type CommentListResponse struct {
	Data []Comment `json:"data"`
	Meta []string  `json:"meta"`
}

var reg = string(strings.Join([]string{
	`Reassigned to .*?`,
	`Milestone changed to .*?`,
	`Title changed from .*? to .*?`,
	`Added .*? label(s)?`,
	`mentioned in commit .*?`,
	`mentioned in merge request .*?`,
	`Status changed to closed`,
	`Status changed to reopened`,
	`moved issue from .*? to .*?`,
	`Marked as \*\*blocked\*\*: .*?`,
	`Marked as \*\*ready\*\* for next stage`,
	`Marked as \*\*unblocked\*\*`,
	`mentioned in issue .*?`,
	`Assignee removed`,
}, "|"))

var regInfo = regexp.MustCompile("^" + reg + "$")

func (c *Comment) MarshalJSON() ([]byte, error) {
	type Alias Comment
	return json.Marshal(struct {
		CreatedAt int64 `json:"created_at"`
		*Alias
	}{
		CreatedAt: c.CreatedAt.Unix(),
		Alias:     (*Alias)(c),
	})
}

func (c *Comment) UnmarshalJSON(b []byte) error {
	type Alias Comment
	aux := (*Alias)(c)

	if err := json.Unmarshal(b, &aux); err != nil {
		return err
	}

	c.IsInfo = regInfo.MatchString(c.Body)

	return nil
}
