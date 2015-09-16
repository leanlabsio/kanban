package gitlab

import _ "encoding/json"

type Project struct {
	Id                int64      `json:"id"`
	Name              string     `json:"name"`
	NamespaceWithName string     `json:"name_with_namespace"`
	PathWithNamespace string     `json:"path_with_namespace"`
	Namespace         *Namespace `json:"namespace,nil,omitempty"`
	Description       string     `json:"description"`
	LastModified      string     `json:"last_modified"`
	CreatedAt         string     `json:"created_at"`
	Owner             *User      `json:"owner,nil,omitempty"`
	AvatarUrl         string     `json:"avatar_url,nil,omitempty"`
}

type Label struct {
	Color string `json:"color"`
	Name  string `json:"name"`
}

type Namespace struct {
	Id     int64   `json:"id"`
	Name   string  `json:"name,omitempty"`
	Avatar *Avatar `json:"avatar,nil,omitempty"`
}

type Avatar struct {
	Url string `json:"url"`
}

type ProjectSingleResponse struct {
	Data Project  `json:"data"`
	Meta []string `json:"meta"`
}

type ProjectListResponse struct {
	Data []Project `json:"data"`
	Meta []string  `json:"meta"`
}

type LabelListResponse struct {
	Data []Label  `json:"data"`
	Meta []string `json:"meta"`
}
