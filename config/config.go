package config

var Conf config

type config struct {
	Project   string
	Module    string
	Host      string
	HttpPort  int               `toml:"http_port"`
	Log       logConf           `toml:"log_conf"`
	EnvConfig environmentConfig `toml:"env_config"`
}

type environmentConfig struct {
	UserName      string `toml:"user_name" env:"SILLYHAT.DB.USER.USERNAME"`
	Password      string `toml:"password" env:"SILLYHAT.DB.USER.PASSWORD"`
	Host          string `toml:"host" env:"SILLYHAT.DB.USER.HOST"`
	Port          string `toml:"port" env:"SILLYHAT.DB.USER.PORT"`
	Schema        string `toml:"schema" env:"SILLYHAT.DB.USER.SCHEMA"`
	DDLPath       string `toml:"ddl_path" env:"SILLYHAT.DB.USER.DDL.PATH"`
	LogstashURL   string `toml:"logstash_url" env:"SILLYHAT.LOGSTASH.URL"`
	ConsulAddress string `toml:"consul_address" env:"SILLYHAT.HOST.CONSUL"`
}

type logConf struct {
	OpenLogstash bool   `toml:"open_logstash"`
	OpenLogfile  bool   `toml:"open_logfile"`
	FilePath     string `toml:"file_path"`
}
