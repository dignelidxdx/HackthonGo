package config

// RestClientConfiguration URL, logging mode.
type RestClientConfiguration struct {
	URL     string `yaml:"url"`
	Timeout int64  `yaml:"timeout"`
}

// DBConfiguration connectionns, lifetime.
type DBConfiguration struct {
	HostEnv      string `yaml:"host_env"`
	UserR        string `yaml:"user_r"`
	PasswordREnv string `yaml:"password_r_env"`
	UserW        string `yaml:"user_w"`
	PasswordWEnv string `yaml:"password_w_env"`
	Schema       string `yaml:"schema"`
	MaxOpenConns int    `yaml:"max_open_conns"`
	MaxIdleConns int    `yaml:"max_idle_conns"`
	MaxLifetime  int    `yaml:"max_lifetime_min"`
}
