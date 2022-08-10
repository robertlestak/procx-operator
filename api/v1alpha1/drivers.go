package v1alpha1

type SqlQuery struct {
	Query  string   `json:"query"`
	Params []string `json:"params"`
}
