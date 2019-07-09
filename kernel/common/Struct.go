package common

//Config 設定檔
type Config struct {
	ENV          string `toml:"env"`
	RedisDefault Redis  `toml:"redis_default"`
}

//Redis Redis config
type Redis struct {
	Name    string `toml:"name"`
	Host    string `toml:"host"`
	MaxConn int    `toml:"max_conn"`
}

//MySQL MySQL DB config
type MySQL struct {
	DB       string `toml:"db"`
	Host     string `toml:"host"`
	Port     string `toml:"port"`
	Username string `toml:"username"`
	Password string `toml:"password"`
	MaxConn  int    `toml:"max_conn"`
	LogMode  bool   `toml:"log_mode"`
}