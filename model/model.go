package model

type URL struct {
	Full     string `json:"full,omitempty"`
	Protocol string `json:"protocol,omitempty"`
	Hostname string `json:"hostname,omitempty"`
	Port     string `json:"port,omitempty"`
	Path     string `json:"pathname,omitempty"`
	Search   string `json:"search,omitempty"`
	Hash     string `json:"hash,omitempty"`
}

type StacktraceFrame struct {
	Path     string `json:"path"`
	Line     int    `json:"lineno"`
	Function string `json:"function"`
}
