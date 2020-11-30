package database

import (
	"context"
	"errors"
	"os"
	"strings"
	"sync"

	"github.com/prometheus/common/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	ErrNotFound  = errors.New("Resource not found")
	ErrDuplicate = errors.New("Duplicate resource")
	ErrDatabase  = errors.New("Error connecting to database")
)

type contextKey string

var (
	transactionContextKey = contextKey("transactionDB")
	poolNameContextKey    = contextKey("poolName")
	defaultConnString     = "host=localhost user=postgres dbname=postgres sslmode=disable password=password"
	defaultPoolMaxConns   = 10
	defaultPoolName       = "default"
	secondaryPoolName     = "subscribers"
	lock                  = &sync.Mutex{}
	conns                 = map[string]*gorm.DB{}
	poolSizes             = map[string]int{
		defaultPoolName:   10,
		secondaryPoolName: 10,
	}
)

func GetDB(ctx context.Context) (*gorm.DB, error) {
	trx, ok := ctx.Value(transactionContextKey).(*gorm.DB)
	if ok {
		return trx, nil
	}
	return GetDBFromPool(ctx)
}

func GetContextWithSecondaryDBPool(ctx context.Context) context.Context {
	return context.WithValue(ctx, poolNameContextKey, secondaryPoolName)
}

func GetDBFromPool(ctx context.Context) (*gorm.DB, error) {
	lock.Lock()
	defer lock.Unlock()
	pool, ok := ctx.Value(poolNameContextKey).(string)
	if !ok {
		pool = defaultPoolName
	}
	if c := conns[pool]; c != nil {
		return c, nil
	}
	connString := defaultConnString
	if os.Getenv("POSTGRES_URL") != "" {
		connString = os.Getenv("POSTGRES_URL")
	}
	var (
		poolSize int
	)
	if poolSize, ok = poolSizes[pool]; !ok {
		poolSize = 10
	}
	log.Debugf("Connecting to PG at %s for pool %s", connString, pool)
	var err error
	conn, err := connect(connString, poolSize)
	if err != nil {
		return nil, err
	}
	conns[pool] = conn
	return conns[pool], nil
}

func connect(connString string, poolSize int) (*gorm.DB, error) {
	conn, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  connString,
		PreferSimpleProtocol: true, // disables implicit prepared statement usage
	}), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// conn.Debug()
	return conn, nil
}

func ResetDB() error {
	db, err := GetDB(context.Background())
	if err != nil {
		return err
	}

	var tables []string
	err = db.Table("pg_tables").
		Where("schemaname = 'public' and tablename != 'schema_migrations'").
		Pluck("tablename", &tables).Error
	if err != nil {
		return err
	}

	err = db.Exec("TRUNCATE TABLE " + strings.Join(tables, ",") + " CASCADE").Error
	if err != nil {
		return err
	}
	return nil
}

// func WrapTransaction(ctx context.Context, fn func(context.Context, *gorm.DB) error) (err error) {
// 	span := opentracing.SpanFromContext(ctx)
// 	_, ok := ctx.Value(transactionContextKey).(*gorm.DB)
// 	if ok {
// 		return errors.New("cannot start new transaction as one is already in progress")
// 	}

// 	dbc, err := GetDB(ctx)
// 	if err != nil {
// 		return err
// 	}

// 	tx := dbc.BeginTx(ctx, nil)
// 	defer func() {
// 		if r := recover(); r != nil {
// 			tx.Rollback()
// 			switch r.(type) {
// 			case string:
// 				err = errors.New(r.(string))
// 			case error:
// 				err = r.(error)
// 			default:
// 				err = errors.New("panic occurred in transaction function")
// 			}
// 			if span != nil {
// 				span.SetTag("error", true)
// 				span.SetTag("stack", string(debug.Stack()))
// 			}
// 		}
// 	}()
// 	if tx.Error != nil {
// 		return tx.Error
// 	}
// 	err = fn(context.WithValue(ctx, transactionContextKey, tx), tx)
// 	if err != nil {
// 		tx.Rollback()
// 		return errors.Wrap(err, "db transaction error")
// 	}
// 	return tx.Commit().Error
// }

// TranslateErrors takes a pointer to a gorm database, gets the errors and transforms
// them into a single micro error which can be returned safely to a handler.
func TranslateErrors(err error) error {
	if err == nil {
		return nil
	}

	if err == gorm.ErrRecordNotFound {
		return ErrNotFound
	}
	if strings.Contains(err.Error(), "unique constraint") {
		return ErrDuplicate
	}

	return ErrDatabase
}
