package model

type RegisterInfo struct {
	*Server
	*Language
	Id                string `json:"id"`
	Version           string `json:"version"`
	Hostname          string `json:"hostname"`
	RegisterIp        string `json:"register_ip"`
	HeartbeatInterval uint64 `json:"heartbeat_interval"`
	RaspHome          string `json:"rasp_home"`
}
