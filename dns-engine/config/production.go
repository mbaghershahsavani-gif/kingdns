package config

type ProductionConfig struct {
	Environment string
	Port        int
}

func Load() ProductionConfig {
	return ProductionConfig{
		Environment: "production",
		Port:        53,
	}
}
