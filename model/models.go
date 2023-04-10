package model

type CronJob struct {
	RepoUrl   string `json:"repositaryLink"`
	Location  string `json:"backupLocation"`
	Frequency int    `json:"backupFrequency"`
}
