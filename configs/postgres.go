package configs

type PostgresConf struct {
	User     string `env:"DATABASE_USER,required"`
	Password string `env:"DATABASE_PASSWORD,required"`
	Host     string `env:"DATABASE_HOST,required"`
	Port     string `env:"DATABASE_PORT,required"`
	HostPort string `env:"HOST_DATABASE_PORT,required"`
	Name     string `env:"DATABASE_NAME,required"`
}
