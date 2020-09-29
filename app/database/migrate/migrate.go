package migrate

import (
	"database/sql"
	"io/ioutil"
	"log"
)

// DropAllTables will empty the database
func DropAllTables(db *sql.DB) {
	query := `
		DO $$ DECLARE
			r RECORD;
		BEGIN
				-- if the schema you operate on is not "current", you will want to
				-- replace current_schema() in query with 'schematodeletetablesfrom'
				-- *and* update the generate 'DROP...' accordingly.
				FOR r IN (SELECT tablename FROM pg_tables WHERE schemaname = current_schema()) LOOP
						EXECUTE 'DROP TABLE IF EXISTS ' || quote_ident(r.tablename) || ' CASCADE';
				END LOOP;
		END $$;
	`
	if _, err := db.Exec(query); err != nil {
		log.Fatal(err)
	}
}

// ImportTables will import all tables from app/database/data.sql
func ImportTables(db *sql.DB) {
	content, err := ioutil.ReadFile("app/database/data.sql")
	if err != nil {
		log.Fatal(err)
	}

	query := string(content)
	if _, err := db.Exec(query); err != nil {
		log.Fatal(err)
	}
}
