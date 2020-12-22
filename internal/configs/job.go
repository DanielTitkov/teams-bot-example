package configs

type JobConfig struct {
	SampleJobPeriod           int `yaml:"sampleJobPeriod"`
	ProjectNotificationPeriod int `yaml:"projectNotificationPeriod"`
}
