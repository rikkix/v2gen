package vmess

import "fmt"

type Link struct {
	Ps   string      `json:"ps"`
	Add  string      `json:"add"`
	Port interface{} `json:"port"`
	Id   string      `json:"id"`
	Aid  interface{} `json:"aid"`
	Net  string      `json:"net"`
	Type string      `json:"type"`
	Host string      `json:"host"`
	Path string      `json:"path"`
	TLS  string      `json:"tls"`
}

func (node Link) Parse() map[string]string {
	Settings := make(map[string]string)

	// set node settings
	Settings["address"] = node.Add
	Settings["serverPort"] = fmt.Sprintf("%v", node.Port)
	Settings["uuid"] = node.Id
	Settings["aid"] = fmt.Sprintf("%v", node.Aid)
	Settings["streamSecurity"] = node.TLS
	Settings["network"] = node.Net
	Settings["tls"] = node.TLS
	Settings["type"] = node.Type
	Settings["host"] = node.Host
	Settings["type"] = node.Type
	Settings["path"] = node.Path
	return Settings
}
