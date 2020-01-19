package conf

const (
	CookieName = "_l_semi_homepage_session"
)

// LoadServerConfig サーバの設定をmapで返す
func LoadServerConfig() map[string]string {
	return map[string]string{
		"addr": "127.0.0.1",
		"host": "localhost",
		"port": "8080",
	}
}

// LoadDatabaseConfig dbの設定をmapで返す
func LoadDatabaseConfig() map[string]string {
	return map[string]string{
		"driver":   "mysql",
		"addr":     "127.0.0.1",
		"host":     "localhost",
		"port":     "3307",
		"user":     "root",
		"password": "password",
		"db":       "homepage",
	}
}
