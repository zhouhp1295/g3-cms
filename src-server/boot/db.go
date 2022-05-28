package boot

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/zhouhp1295/g3/crud"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"net/url"
	"strings"
	"time"
	"unknwon.dev/clog/v2"
)

const (
	MySQL      = "mysql"
	SQLite3    = "sqlite3"
	PostgreSQL = "postgres"
)

type GormLogger struct {
}

func (*GormLogger) Printf(format string, v ...interface{}) {
	clog.InfoTo("db", format, v...)
}

// parsePostgreSQLHostPort parses given input in various forms defined in
// https://www.postgresql.org/docs/current/static/libpq-connect.html#LIBPQ-CONNSTRING
// and returns proper host and port number.
func parsePostgreSQLHostPort(info string) (host, port string) {
	host, port = "127.0.0.1", "5432"
	if strings.Contains(info, ":") && !strings.HasSuffix(info, "]") {
		idx := strings.LastIndex(info, ":")
		host = info[:idx]
		port = info[idx+1:]
	} else if len(info) > 0 {
		host = info
	}
	return host, port
}

// parseDSN takes given database options and returns parsed DSN.
func parseDSN(cfg DatabaseConfig) (dsn string, err error) {
	// In case the database name contains "?" with some parameters
	concate := "?"
	if strings.Contains(cfg.Name, concate) {
		concate = "&"
	}

	switch cfg.Type {
	case MySQL:
		if cfg.Host[0] == '/' { // Looks like a unix socket
			dsn = fmt.Sprintf("%s:%s@unix(%s)/%s%scharset=utf8mb4&parseTime=true&loc=Local",
				cfg.User, cfg.Password, cfg.Host, cfg.Name, concate)
		} else {
			dsn = fmt.Sprintf("%s:%s@tcp(%s)/%s%scharset=utf8mb4&parseTime=true&loc=Local",
				cfg.User, cfg.Password, cfg.Host, cfg.Name, concate)
		}
	case PostgreSQL:
		host, port := parsePostgreSQLHostPort(cfg.Host)
		if host[0] == '/' { // looks like a unix socket
			dsn = fmt.Sprintf("postgres://%s:%s@:%s/%s%ssslmode=%s&host=%s",
				url.QueryEscape(cfg.User), url.QueryEscape(cfg.Password), port, cfg.Name, concate, cfg.SSLMode, host)
		} else {
			dsn = fmt.Sprintf("postgres://%s:%s@%s:%s/%s%ssslmode=%s",
				url.QueryEscape(cfg.User), url.QueryEscape(cfg.Password), host, port, cfg.Name, concate, cfg.SSLMode)
		}
	case SQLite3:
		dsn = "file:" + EnsureAbs(cfg.Name+".db") + "?cache=shared&mode=rwc&_busy_timeout=9999999"
	default:
		return "", errors.Errorf("unrecognized dialect: %s", cfg.Type)
	}

	return dsn, nil
}

func openDB(cfg DatabaseConfig, ormCfg *gorm.Config) (*gorm.DB, error) {
	dsn, err := parseDSN(cfg)
	if err != nil {
		return nil, errors.Wrap(err, "parse DSN")
	}

	var dialector gorm.Dialector
	switch cfg.Type {
	case MySQL:
		dialector = mysql.Open(dsn)
	case PostgreSQL:
		dialector = postgres.Open(dsn)
	case SQLite3:
		dialector = sqlite.Open(dsn)
	default:
		panic("未定义的数据库类型:" + cfg.Type)
	}

	return gorm.Open(dialector, ormCfg)
}

// AutoMigrateTables is the list of struct-to-table mappings.
//var AutoMigrateTables []interface{}

func TestDatabaseConn(cfg DatabaseConfig) (*gorm.DB, error) {
	return openDB(cfg, &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		NowFunc: func() time.Time {
			return time.Now().Local().Truncate(time.Microsecond)
		},
	})
}

func initGormDB(w logger.Writer) {
	//level := logger.Info
	level := logger.Warn
	if IsProdMode() {
		level = logger.Error
	}
	// NOTE: AutoMigrate does not respect logger passed in gorm.Config.
	logger.Default = logger.New(w, logger.Config{
		SlowThreshold: 100 * time.Millisecond,
		LogLevel:      level,
	})
	db, err := openDB(DatabaseCfg, &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		NowFunc: func() time.Time {
			return time.Now().Local().Truncate(time.Microsecond)
		},
	})
	if err != nil {
		panic(errors.Wrap(err, "数据库初始化失败"))
	}

	sqlDB, err := db.DB()
	if err != nil {
		panic(errors.Wrap(err, "数据库初始化失败"))
	}
	sqlDB.SetMaxOpenConns(DatabaseCfg.MaxOpenConns)
	sqlDB.SetMaxIdleConns(DatabaseCfg.MaxIdleConns)
	sqlDB.SetConnMaxLifetime(time.Minute)

	switch DatabaseCfg.Type {
	case MySQL:
		db = db.Set("gorm:table_options", "ENGINE=InnoDB").Session(&gorm.Session{})
	case PostgreSQL:
	case SQLite3:
	default:
		panic("未定义的数据库类型:" + DatabaseCfg.Type)
	}

	crud.InitDbEngine(db)
}
