package models

type ConsulAgentResponse []struct {
	ID       string `json:"ID"`
	Node     string `json:"Node"`
	Address  string `json:"Address"`
	NodeMeta struct {
		Enabled string `json:"enabled"`
	} `json:"NodeMeta"`
	ServiceName string `json:"ServiceName"`
	ServicePort int    `json:"ServicePort"`
}
