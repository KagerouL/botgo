package main

func initUserDB() {

	db.Exec(`
	CREATE TABLE IF NOT EXISTS users (
		jid TEXT PRIMARY KEY
	)
	`)
}

func isFirstTime(jid string) bool {

	var exists string

	err := db.QueryRow(
		"SELECT jid FROM users WHERE jid=?",
		jid,
	).Scan(&exists)

	if err != nil {

		db.Exec(
			"INSERT INTO users(jid) VALUES(?)",
			jid,
		)

		return true
	}

	return false
}