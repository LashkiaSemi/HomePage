package conf

const (
	CookieName   = "_l_semi_homepage_session"
	CookieMaxAge = 3600 * 1
)

const (
	DatetimeFormat    = "2006/01/02 15:04:05"
	DateFormat        = "2006/01/02"
	LogDatetimeFormat = "2006/01/02 15:04:05"
)

const (
	// DefaultTagID 備品を登録する時、タグの入力がなかったら、その他に誘導したい...
	DefaultTagID = 6
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
