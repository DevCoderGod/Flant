package main

// Config contains application configuration.
type Config interface {
	GetQueueName() string
	GetRmqURI() string
	GetServerAddress() string
	GetPSQLURI() string

	// GetQueueNameStat() string
	// GetQueueNameProfiler() string
}

type ConfJSON struct {
	PSQLURI       string
	ServerAddress string
	RmqURI        string
	QueueName     string
}

func (c ConfJSON) GetQueueName() string {
	return c.QueueName
}
func (c ConfJSON) GetRmqURI() string {
	return c.RmqURI
}
func (c ConfJSON) GetServerAddress() string {
	return c.ServerAddress
}
func (c ConfJSON) GetPSQLURI() string {
	return c.PSQLURI
}
