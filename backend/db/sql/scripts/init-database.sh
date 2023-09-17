#run the setup script to create the DB and the schema in the DB
mysql -u admin -padmin mdet < "/docker-entrypoint-initdb.d/001-create-tables.sql"
