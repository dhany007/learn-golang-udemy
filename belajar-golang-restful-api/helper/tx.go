package helper

import "database/sql"

// fungsi commit atau rollback
// setiap commit atau rollback bisa saja menimbulkan error
func CommitOrCallback(tx *sql.Tx) {
	err := recover()

	if err != nil {
		errorRollback := tx.Rollback()
		PanicIfError(errorRollback)
		panic(err)
	} else {
		errorCommit := tx.Commit()
		PanicIfError(errorCommit)
	}
}
