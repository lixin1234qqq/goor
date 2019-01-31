package model

type RegisterInfo struct {
	Id                string `json:"id"`
	Version           string `json:"version"`
	Hostname          string `json:"hostname"`
	RegisterIp        string `json:"register_ip"`
	Language          string `json:"language"`
	LanguageVersion   string `json:"language_version"`
	ServerType        string `json:"server_type"`
	ServerVersion     string `json:"server_version"`
	HeartbeatInterval uint64 `json:"heartbeat_interval"`
	RaspHome          string `json:"rasp_home"`
}
