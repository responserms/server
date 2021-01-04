package config

func applyDefaults(cfg *Config) *Config {
	// apply `events` defaults
	cfg = applyEventsDefaults(cfg)

	// apply `database` defaults
	cfg = applyDatabaseDefaults(cfg)

	// apply `cluster` defaults
	cfg = applyClusterDefaults(cfg)

	// apply `http` defaults
	cfg = applyHTTPDefaults(cfg)

	return cfg
}

func applyClusterDefaults(cfg *Config) *Config {
	if cfg.Cluster.Environment == "" {
		cfg.Cluster.Environment = "local"
	}

	if cfg.Cluster.BindAddress == "" {
		cfg.Cluster.BindAddress = "0.0.0.0"
	}

	if cfg.Cluster.BindPort == 0 {
		cfg.Cluster.BindPort = 3320
	}

	if cfg.Cluster.BindMemberAddress == "" {
		cfg.Cluster.BindMemberAddress = "0.0.0.0"
	}

	if cfg.Cluster.BindMemberPort == 0 {
		cfg.Cluster.BindMemberPort = 3322
	}

	return cfg
}

func applyEventsDefaults(cfg *Config) *Config {
	if cfg.Events.Type == "" {
		cfg.Events.Type = "embedded"
	}

	if cfg.Events.Subject == "" {
		cfg.Events.Subject = "response-events"
	}

	switch cfg.Events.Type {
	case "nats":
		if cfg.Events.URL == "" {
			cfg.Events.URL = "nats://localhost:4222"
		}
	case "embedded":
		cfg.Events.URL = "mem://"
	}

	return cfg
}

func applyDatabaseDefaults(cfg *Config) *Config {
	if cfg.Database.Type == "" {
		cfg.Database.Type = "embedded"
	}

	if cfg.Database.Host == "" {
		cfg.Database.Host = "localhost"
	}

	// apply database defaults based on the type
	switch cfg.Database.Type {
	case "postgresql":
		if cfg.Database.Port == 0 {
			cfg.Database.Port = 5432
		}

		if cfg.Database.Username == "" {
			cfg.Database.Username = "postgres"
		}
	case "mysql":
		if cfg.Database.Port == 0 {
			cfg.Database.Port = 3306
		}

		if cfg.Database.Username == "" {
			cfg.Database.Username = "root"
		}
	}

	return cfg
}

func applyHTTPDefaults(cfg *Config) *Config {
	if cfg.HTTP.BindAddress == "" {
		cfg.HTTP.BindAddress = "0.0.0.0"
	}

	if cfg.HTTP.Port == 0 {
		cfg.HTTP.Port = 8080
	}

	// if cfg.HTTP.MaxUploadSize == 0 {
	// 	cfg.HTTP.MaxUploadSize = 10000
	// }

	return cfg
}
