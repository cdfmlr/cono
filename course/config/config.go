package config

// ConfDatabase is a struct for Database configures. Provides DSN (Data Source Name):
//    [username[:password]@][protocol[(address)]]/dbname
// Refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details.
type ConfDatabase struct {
	Username string
	Password string
	Protocol string
	Address  string
	DBName   string
}

// ConfWxOfficialAccount 是微信公众号服务的配置
type ConfWxOfficialAccount struct {
	AppID     string
	AppSecret string
	Token     string
}

// Conf is a struct wraps all configures.
// field XXX -> <type ConfXXX struct>
type Conf struct {
	// 数据库配置
	Database ConfDatabase
	// 强智教务系统地址的学校字段
	QzSchool string
	// 微信公众号服务配置
	WxOfficialAccount ConfWxOfficialAccount
	// StudentRPC 服务地址
	StudentRPCAddress string
}
