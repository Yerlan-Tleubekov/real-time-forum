package repository

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func OpenDB(driver, path string) (*sql.DB, error) {
	db, err := sql.Open(driver, path)
	if err != nil {
		return nil, err
	}

	db.SetMaxIdleConns(100)

	if err = db.Ping(); err != nil {
		return nil, err

	}

	if err = createAllTables(db); err != nil {
		return nil, err
	}

	return db, nil
}

func createAllTables(database *sql.DB) error {

	tx, err := database.Begin()
	if err != nil {
		tx.Rollback()
		return err
	}

	if _, err := tx.Exec(`
	CREATE TABLE IF NOT EXISTS user(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		nickname TEXT UNIQUE NOT NULL,
		age INTEGER NOT NULL,
		gender TEXT NOT NULL,
		first_name TEXT NOT NULL,
		last_name TEXT NOT NULL,
		email TEXT UNIQUE NOT NULL,
		password TEXT NOT NULL,
		created_date DATE NOT NULL,
		role TEXT NOT NULL
	);`); err != nil {
		tx.Rollback()
		return err
	}

	if _, err := tx.Exec(`
	CREATE TABLE IF NOT EXISTS post(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		user_id INTEGER,
		title TEXT NOT NULL,
		content TEXT NOT NULL,
		created_date DATE,
		FOREIGN KEY (user_id) REFERENCES user (id)
	);`); err != nil {
		tx.Rollback()
		return err
	}

	if _, err := tx.Exec(`
	CREATE TABLE IF NOT EXISTS category(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL UNIQUE 
	);`); err != nil {
		tx.Rollback()
		return err
	}

	if _, err := tx.Exec(`
	CREATE TABLE IF NOT EXISTS posts_categories(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		post_id INTEGER,
		category_id INTEGER,
		FOREIGN KEY (category_id) REFERENCES category (id),
		FOREIGN KEY (post_id) REFERENCES post (id)
	);`); err != nil {
		tx.Rollback()
		return err
	}

	if _, err := tx.Exec(`
	CREATE TABLE IF NOT EXISTS comment(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		post_id INTEGER,
		user_id INTEGER,
		text TEXT,
		created_date DATE,
		FOREIGN KEY (user_id) REFERENCES user (id),
		FOREIGN KEY (post_id) REFERENCES post (id)
	);`); err != nil {
		tx.Rollback()
		return err
	}

	if _, err := tx.Exec(`
	CREATE TABLE IF NOT EXISTS online_user(
		user_id INTEGER NOT NULL,
		FOREIGN KEY (user_id) REFERENCES user (id)
	);`); err != nil {
		tx.Rollback()
		return err
	}

	if _, err := tx.Exec(`
		CREATE TABLE IF NOT EXISTS dialog_room(
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			first_user_id INTEGER NOT NULL,
			first_user_name TEXT NOT NULL,
			second_user_id INTEGER NOT NULL,
			second_user_name TEXT NOT NULL,
			created_date DATE,
			last_message_date DATE,
			FOREIGN KEY (first_user_id) REFERENCES user (id),
			FOREIGN KEY (second_user_id) REFERENCES user (id),
			UNIQUE (first_user_id, second_user_id)
		);
	`); err != nil {
		tx.Rollback()
		return err
	}

	if _, err := tx.Exec(`
	CREATE TABLE IF NOT EXISTS message_history(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		dialog_room_id INTEGER NOT NULL,
		message TEXT NOT NULL,
		user_id INTEGER NOT NULL,
		created_date DATE,
		FOREIGN KEY (dialog_room_id) REFERENCES dialog_room (id)
		FOREIGN KEY (user_id) REFERENCES user (id)
	);`); err != nil {
		tx.Rollback()
		return err
	}

	if _, err := tx.Exec(`
	CREATE TABLE IF NOT EXISTS queue_notifications(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		message_history_id INTEGER NOT NULL,
		FOREIGN KEY (message_history_id) REFERENCES message_history (id)
	);`); err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}
