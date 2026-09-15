package config

type Config struct {
	UDPAddress string
	TCPAddress string
}

func Default() Config {
	return Config{
		UDPAddress: ":53",
		TCPAddress: ":53",
	}
}
