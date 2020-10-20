package element

type API struct {
	Title          string          `json:"title"`
	Description    string          `json:"description"`
	Host           string          `json:"host"`
	Metadata       []Metadata      `json:"metadata"`
	ResourceGroups []ResourceGroup `json:"resource_groups"`
	Resources      []Resource      `json:"resources"`
}

type Metadata struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type ResourceGroup struct {
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Resources   []Resource `json:"resources"`
}

type Resource struct {
	Title       string       `json:"title"`
	Description string       `json:"description"`
	Href        Href         `json:"href"`
	Transitions []Transition `json:"transitions"`
}

type Href struct {
	Path      string         `json:"path"`
	Variables []HrefVariable `json:"variables"`
}

type HrefVariable struct {
	Title       string `json:"title"`
	Required    bool   `json:"required"`
	Description string `json:"description"`
	Key         string `json:"key"`
	Value       string `json:"value"`
}

type Transition struct {
	Title        string        `json:"title"`
	Description  string        `json:"description"`
	Method       string        `json:"method"`
	Href         Href          `json:"href"`
	Transactions []Transaction `json:"transactions"`
}

type Transaction struct {
	Request  Request  `json:"request"`
	Response Response `json:"response"`
}

type Request struct {
	Description string   `json:"description"`
	Method      string   `json:"method"`
	Headers     []Header `json:"headers"`
	Body        Asset    `json:"body"`
	Schema      Asset    `json:"schema"`
}

type Response struct {
	Description string   `json:"description"`
	StatusCode  int      `json:"status_code"`
	Headers     []Header `json:"headers"`
	Body        Asset    `json:"body"`
	Schema      Asset    `json:"schema"`
}

type Header struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type Asset struct {
	ContentType string `json:"content_type"`
	Body        string `json:"body"`
}
