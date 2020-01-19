package repository

// Transact トランザクションを作成する
func Transact(sh SQLHandler, txFunc func(Tx) error) (err error) {
	tx, err := sh.Begin()
	if err != nil {
		return
	}
	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p)
		} else if err != nil {
			tx.Rollback()
		} else {
			err = tx.Commit()
		}
	}()
	err = txFunc(tx)
	return err
}
