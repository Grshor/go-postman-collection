package postman

type Request struct {
	URL         interface{} `json:"url"`
	Auth        interface{} `json:"auth,omitempty"`
	Proxy       interface{} `json:"proxy,omitempty"`
	Certificate interface{} `json:"certificate,omitempty"`
	Method      string      `json:"method"`
	Description interface{} `json:"description,omitempty"`
	Header      interface{} `json:"header,omitempty"`
	Body        interface{} `json:"body,omitempty"`
}

func NewRequest(URL string, method string) *Request {
	return &Request{
		URL:    URL,
		Method: method,
	}
}
