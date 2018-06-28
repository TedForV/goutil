package configloader

const (
	EnvDev = "dev"
	EnvPro = "pro"
)

type BaseMSConfig struct {
	IP          string      `yaml:"ip"`
	Port        int         `yaml:"port"`
	Env         string      `yaml:"env"`
	ServiceInfo ServiceInfo `yaml:"serviceInfo"`
	EtcdInfo    ETCDInfo    `yaml:"etcdInfo"`
}

type ETCDInfo struct {
	EtcdServer string `yaml:"etcdServer"`
	Instance   string `yaml:"instance"`
}

type ServiceInfo struct {
	ServiceId     int    `yaml:"serviceId"`
	ServiceTypeId int    `yaml:"serviceTypeId"`
	ServicePrefix string `yaml:"servicePrefix"`
}
