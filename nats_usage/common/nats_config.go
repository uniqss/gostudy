package common

type NatsConfig struct {
	NatsAddr      string
	ServiceAddr   string
	BroadcastAddr string
}

func NewDefaultNatsConfig() *NatsConfig {
	return &NatsConfig{
		NatsAddr: "127.0.0.1:4222",
	}
}
