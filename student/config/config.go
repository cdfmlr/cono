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

// ConfServe is a struct for serving configures.
type ConfServe struct {
	// StudentRPCAddress is the address gRPC serves StudentRPC
	StudentRPCAddress string
}

// Conf is a struct wraps all configures.
// field XXX -> <type ConfXXX struct>
type Conf struct {
	Database ConfDatabase
	Serve    ConfServe
}
