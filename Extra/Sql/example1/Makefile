

data.db: schema.sql
	sqlite3 data.db < schema.sql

.PHONY: clean
clean:
	rm -rf data.db
