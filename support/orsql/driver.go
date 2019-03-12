package orsql

import (
	"database/sql"
	"database/sql/driver"
	"errors"
	"sync"

	"github.com/baidu/openrasp"
	"github.com/baidu/openrasp/model"
)

var (
	driversMu sync.RWMutex
	drivers   = make(map[string]*wrapDriver)
)

type DSNParserFunc func(dsn string) DSNInfo
type ErrorInterceptorFunc func(err *error) (bool, string, string)

func genericDSNParser(string) DSNInfo {
	return DSNInfo{}
}

func genericErrorInterceptor(err *error) (bool, string, string) {
	return false, "", ""
}

func Register(name string, driver driver.Driver, opts ...WrapOption) {
	driversMu.Lock()
	defer driversMu.Unlock()

	wrapped := newWrapDriver(driver, opts...)
	sql.Register(wrapDriverName(name), wrapped)
	drivers[name] = wrapped
}

func wrapDriverName(origin string) string {
	return "openrasp/" + origin
}

func Open(driverName, dataSourceName string) (*sql.DB, error) {
	return sql.Open(wrapDriverName(driverName), dataSourceName)
}

func Wrap(driver driver.Driver, opts ...WrapOption) driver.Driver {
	return newWrapDriver(driver, opts...)
}

func newWrapDriver(driver driver.Driver, opts ...WrapOption) *wrapDriver {
	d := &wrapDriver{
		Driver: driver,
	}
	for _, opt := range opts {
		opt(d)
	}
	if d.driverName == "" {
		d.driverName = ExtractName(driver)
	}
	if d.dsnParser == nil {
		d.dsnParser = genericDSNParser
	}
	if d.errorInterceptor == nil {
		d.errorInterceptor = genericErrorInterceptor
	}
	return d
}

func DriverDSNParser(driverName string) DSNParserFunc {
	driversMu.RLock()
	driver := drivers[driverName]
	defer driversMu.RUnlock()
	return driver.dsnParser
}

type WrapOption func(*wrapDriver)

func DriverNameWrap(name string) WrapOption {
	return func(d *wrapDriver) {
		d.driverName = name
	}
}

func DSNParserWrap(f DSNParserFunc) WrapOption {
	return func(d *wrapDriver) {
		d.dsnParser = f
	}
}

func ErrorInterceptorWrap(f ErrorInterceptorFunc) WrapOption {
	return func(d *wrapDriver) {
		d.errorInterceptor = f
	}
}

type wrapDriver struct {
	driver.Driver
	driverName       string
	dsnParser        DSNParserFunc
	errorInterceptor ErrorInterceptorFunc
}

func (d *wrapDriver) Open(name string) (driver.Conn, error) {
	dsnInfo := d.dsnParser(name)
	dbConnParam := NewDbConnectionParam(&dsnInfo, d.driverName)
	interceptCode, _ := dbConnParam.PolicyCheck()
	//TODO log
	if interceptCode == model.Block {
		panic(openrasp.ErrBlock)
	}
	conn, err := d.Driver.Open(name)
	if err != nil {
		return nil, err
	} else {
		if interceptCode == model.Log {
			//TODO log
		}
	}
	return newConn(conn, d, dsnInfo), nil
}

func namedValueToValue(named []driver.NamedValue) ([]driver.Value, error) {
	dargs := make([]driver.Value, len(named))
	for n, param := range named {
		if len(param.Name) > 0 {
			return nil, errors.New("sql: driver does not support the use of Named Parameters")
		}
		dargs[n] = param.Value
	}
	return dargs, nil
}

type namedValueChecker interface {
	CheckNamedValue(*driver.NamedValue) error
}

func checkNamedValue(nv *driver.NamedValue, next namedValueChecker) error {
	if next != nil {
		return next.CheckNamedValue(nv)
	}
	return driver.ErrSkip
}
