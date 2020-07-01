package config

type Config struct {
	Storage struct{
		Redis struct{
			URL string `mapstructure:"url"`
		} `mapstructure:"redis"`

		Mysql struct{
			DSN string `mapstructure:"dsn"`
		} `mapstructure:"mysql"`

	} `mapstructure:"storage"`

	NetworkServer struct{
		Api struct{
			Bind string `mapstructure:"bind"`
		} `mapstructure:"api"`

	} `mapstructure:"network_server"`
}

// C holds the global configuration.
var C Config