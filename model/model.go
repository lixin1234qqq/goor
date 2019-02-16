package model

import "net/url"

type URL struct {
	Full     string `json:"full,omitempty"`
	Protocol string `json:"protocol,omitempty"`
	Hostname string `json:"hostname,omitempty"`
	Port     string `json:"port,omitempty"`
	Path     string `json:"pathname,omitempty"`
	Search   string `json:"search,omitempty"`
	Hash     string `json:"hash,omitempty"`
}

type Request struct {
	Method       string `json:"request_method"`
	UrlFull      string `json:"url"`
	UrlHostname  string `json:"target"`
	UrlPath      string `json:"path"`
	Referer      string `json:"referer"`
	UserAgent    string `json:"user_agent"`
	AttackSource string `json:"attack_source"`
	ClientIp     string `json:"client_ip"`
	RequestId    string `json:"request_id"`
	*RequestBody
}

type RequestBody struct {
	Raw  string     `json:"body"`
	Form url.Values `json:"-"`
}

type StacktraceFrame struct {
	Path     string `json:"path"`
	Line     int    `json:"lineno"`
	Function string `json:"function"`
}
