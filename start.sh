go mod tidy
go get -v github.com/rubenv/sql-migrate/...@f842348935589e4563be545226d465178bd439cf
sql-migrate up
