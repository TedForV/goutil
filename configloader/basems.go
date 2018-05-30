package configloader

const (
	EnvDev = "dev"
	EnvPro = "pro"
)

type BaseMSConfig struct {
	IP       string   `yaml:"ip"`
	Port     int      `yaml:"port"`
	Env      string   `yaml:"env"`
	EtcdInfo ETCDInfo `yaml:"etcdInfo"`
}

type ETCDInfo struct {
	EtcdServer    string `yaml:"etcdServer"`
	ServicePrefix string `yaml:"servicePrefix"`
	Instance      string `yaml:"instance"`
}
