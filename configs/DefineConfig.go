package configs

const (
	SERVER_LIVE_TYPE  = "local"
	SERVICE_NAME      = "fcm-sender"
	SERVICE_HOST_NAME = "localhost"
	SERVICE_PORT      = "4100"
	DEBUG_LEVEL       = "ERROR"
	DEV_DEBUG_LEVEL   = "TRACE"
	LOG_STYLE         = "CONSOLE"
)

const (
	CodeRdbConfigDatabaseTypeMaster = "master"
	CodeRdbConfigDatabaseTypeLog    = "log"

	//# RDB

	RDB_HOST   = ""
	RDB_PORT   = "3306"
	RDB_USER   = ""
	RDB_PW     = ""
	RDB_DBNAME = ""

	RDB_HOST_TEST   = ""
	RDB_PORT_TEST   = ""
	RDB_USER_TEST   = ""
	RDB_PW_TEST     = ""
	RDB_DBNAME_TEST = ""

	RDB_USE_MASTER_DB = "true"
	RDB_TYPE          = "mariadb"

	RDB_TIMEOUT       = "600"
	RDB_MAX_IDLE_CNT  = "1000"
	RDB_MAX_OPEN_CONN = "1000"
	RDB_CHARSET       = "utf8"
	RDB_TIMEZONE      = "UTC"
)

const (
	RdbTimeout              string = "600"
	RdbMaxIdleCnt           string = "1000"
	RdbMaxOpenConn          string = "1000"
	RdbCharset              string = "utf8mb4"
	RdbTimezone             string = "UTC"
	RdbDatabasesType        string = "mariadb" // mariadb or mysql
	RdbDatabasesTypeMariaDb string = "mariadb"
	RdbDatabasesTypeMysql   string = "mysql"
	RdbDebugLevel           string = "TRACE"
)

var (
	RdbMasterDatabaseNameLive string = "order"
	RdbMasterDatabaseNameTest string = "order_test"
)
