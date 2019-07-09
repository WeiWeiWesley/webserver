package orm

//Host Host info
type Host struct {
	DB       string `toml:"db"`
	Host     string `toml:"host"`
	Port     string `toml:"port"`
	Username string `toml:"username"`
	Password string `toml:"password"`
	MaxConn  int    `toml:"max_conn"`
	LogMode  bool   `toml:"log_mode"`
}
