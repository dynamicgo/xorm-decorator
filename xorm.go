package decorator

import (
	"github.com/go-xorm/xorm"
)

// DuplicateKey check if error is a duplicate key error
func DuplicateKey(db *xorm.Engine, err error) bool {

	dialect := getDialect(string(db.Dialect().DBType()))

	return dialect.DuplicateKey(err)
}

// Dialect .
type Dialect interface {
	DuplicateKey(err error) bool
}
