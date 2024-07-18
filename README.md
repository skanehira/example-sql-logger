# example-sql-logger
Example for go-sql-proxy + sqlc

```sh
$ go run .
{"time":"2024-07-18T13:45:11.5368+09:00","level":"INFO","msg":"DB Open"}
{"time":"2024-07-18T13:45:11.537219+09:00","level":"INFO","msg":"DB Exec","query":"CREATE TABLE IF NOT EXISTS t1 (id INTEGER PRIMARY KEY, name TEXT NOT NULL)","args":[]}
{"time":"2024-07-18T13:45:11.537444+09:00","level":"INFO","msg":"DB Exec","query":"INSERT INTO t1 (id, name) VALUES (?, ?)","args":[{"Name":"","Ordinal":1,"Value":1},{"Name":"","Ordinal":2,"Value":"foo"}]}
{"time":"2024-07-18T13:45:11.537521+09:00","level":"INFO","msg":"DB Query","query":"SELECT id, name FROM t1","args":[]}
{"time":"2024-07-18T13:45:11.537606+09:00","level":"INFO","msg":"DB Exec","query":"UPDATE t1 SET name = ? WHERE id = ?","args":[{"Name":"","Ordinal":1,"Value":"bar"},{"Name":"","Ordinal":2,"Value":1}]}
{"time":"2024-07-18T13:45:11.53763+09:00","level":"INFO","msg":"DB Exec","query":"DELETE FROM t1 WHERE id = ?","args":[{"Name":"","Ordinal":1,"Value":1}]}
{"time":"2024-07-18T13:45:11.537643+09:00","level":"INFO","msg":"DB Begin"}
{"time":"2024-07-18T13:45:11.537664+09:00","level":"INFO","msg":"DB Exec","query":"INSERT INTO t1 (id, name) VALUES (?, ?)","args":[{"Name":"","Ordinal":1,"Value":2},{"Name":"","Ordinal":2,"Value":"baz"}]}
{"time":"2024-07-18T13:45:11.53768+09:00","level":"INFO","msg":"DB Query","query":"SELECT id, name FROM t1","args":[]}
{"time":"2024-07-18T13:45:11.537748+09:00","level":"INFO","msg":"DB Begin"}
```
