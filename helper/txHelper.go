package helper

import "database/sql"

func TxErrHandle(tx *sql.Tx) {
	err := recover()
	if err != nil {
		errorRollback := tx.Rollback()
		PanicErrorIf(errorRollback)
		panic(err)
	} else {
		errorCommit := tx.Commit()
		PanicErrorIf(errorCommit)
	}
}
