package main

import (
	"context"
	"database/sql"
	"database/sql/driver"
	"log"
	"os"
	"strings"

	"github.com/mattn/go-sqlite3"
	proxy "github.com/shogo82148/go-sql-proxy"
	"golang.org/x/exp/slog"
)

var logger = slog.New(slog.NewJSONHandler(os.Stdout, nil))

func must(_r any, err error) {
	if err != nil {
		log.Fatalf("Exec failed: %v", err)
	}
}

func trimQueryComment(query string) string {
	if strings.Contains(query, "--") {
		return strings.TrimRight(strings.SplitN(query, "\n", 2)[1], "\n")
	}
	return query
}

func main() {
	sql.Register("sqlite3-proxy", proxy.NewProxyContext(&sqlite3.SQLiteDriver{}, &proxy.HooksContext{
		Open: func(ctx context.Context, _ any, conn *proxy.Conn) error {
			logger.InfoContext(ctx, "DB Open")
			return nil
		},
		Exec: func(ctx context.Context, _ any, stmt *proxy.Stmt, args []driver.NamedValue, result driver.Result) error {
			logger.InfoContext(ctx, "DB Exec", slog.String("query", trimQueryComment(stmt.QueryString)), slog.Any("args", args))
			return nil
		},
		Query: func(ctx context.Context, _ any, stmt *proxy.Stmt, args []driver.NamedValue, rows driver.Rows) error {
			logger.InfoContext(ctx, "DB Query", slog.String("query", trimQueryComment(stmt.QueryString)), slog.Any("args", args))
			return nil
		},
		Begin: func(ctx context.Context, _ any, conn *proxy.Conn) error {
			logger.InfoContext(ctx, "DB Begin")
			return nil
		},
		Commit: func(ctx context.Context, _ any, tx *proxy.Tx) error {
			logger.InfoContext(ctx, "DB Begin")
			return nil
		},
		Rollback: func(ctx context.Context, _ any, tx *proxy.Tx) error {
			logger.InfoContext(ctx, "DB Rollback")
			return nil
		},
	}))

	db, err := sql.Open("sqlite3-proxy", ":memory:")
	if err != nil {
		log.Fatalf("Open filed: %v", err)
	}
	defer db.Close()

	q := New(db)

	ctx := context.Background()
	must(q.db.ExecContext(ctx, "CREATE TABLE IF NOT EXISTS t1 (id INTEGER PRIMARY KEY, name TEXT NOT NULL)"))
	must(q.CreateT1(ctx, CreateT1Params{
		ID:   1,
		Name: "foo",
	}))
	must(q.GetT1(ctx))
	must(q.UpdateT1(ctx, UpdateT1Params{
		ID:   1,
		Name: "bar",
	}))
	must(q.DeleteT1(ctx, 1))

	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		log.Fatalf("BeginTx failed: %v", err)
	}

	q = q.WithTx(tx)

	must(q.CreateT1(ctx, CreateT1Params{
		ID:   2,
		Name: "baz",
	}))
	must(q.GetT1(ctx))

	if err := tx.Commit(); err != nil {
		log.Fatalf("Commit failed: %v", err)
	}
}
