package configs

type Server struct {
	Port string `json:"port"`
}

type Redis struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	Password string `json:"password"`
}

type Options struct {
	Schema string `json:"schema"`
	Prefix string `json:"prefix"`
	Mode   string `json:"mode"`
}

type Config struct {
	Server  `json:"server"`
	Redis   `json:"redis"`
	Options `json:"options"`
}
