package postgres

import (
	"database/sql"
)

type Tx struct {
	*sql.Tx
}
