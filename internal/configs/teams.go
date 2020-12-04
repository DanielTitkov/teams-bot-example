package configs

type TeamsConfig struct {
	AppID       string `yaml:"appID"`
	AppPassword string `yaml:"appPassword"`
	Port        int    `yaml:"port"`
}
