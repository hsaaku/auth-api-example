package migrate

import (
	"io/ioutil"
	"log"

	"github.com/jmoiron/sqlx"
)

// DropAllTables will empty the database
func DropAllTables(db *sqlx.DB) {
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

	db.MustExec(query)
}

// ImportTables will import all tables from app/database/data.sql
func ImportTables(db *sqlx.DB) {
	content, err := ioutil.ReadFile("app/database/data.sql")
	if err != nil {
		log.Fatal(err)
	}

	query := string(content)
	db.MustExec(query)
}
