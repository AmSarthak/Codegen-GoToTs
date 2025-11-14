package models

type Item struct {
	ID           string   `json:"id"`
	ServerType   string   `json:"serverType"`
	Manufacturer string   `json:"manufacturer"`
	ModelNo      string   `json:"modelNo"`
	CpuModel     string   `json:"cpuModel"`
	CpuMemory    int      `json:"cpuMemory"`
	Gpus         []string `json:"gpus"`
}
