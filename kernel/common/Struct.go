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
