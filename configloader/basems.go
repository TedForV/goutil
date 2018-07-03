package configloader

//define the env model type
const (
	EnvDev = "dev"
	EnvPro = "pro"
)

// BaseMSConfig define a common used configuration model,
type BaseMSConfig struct {
	IP                 string      `yaml:"ip"`
	Port               int         `yaml:"port"`
	Env                string      `yaml:"env"`
	ServiceInfo        ServiceInfo `yaml:"serviceInfo"`
	EtcdServer         string      `yaml:"etcdServer"`
	LocalLogFolderPath string      `yaml:"localLogFolderPath"`
}

// ServiceInfo define service model
type ServiceInfo struct {
	ServiceId     int    `yaml:"serviceId"`
	ServiceTypeId int    `yaml:"serviceTypeId"`
	ServicePrefix string `yaml:"servicePrefix"`
}
