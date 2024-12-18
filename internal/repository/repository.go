package repository

import (
	"context"
	"database/sql"
	"fmt"
	"mini-stats-server/config"

	"github.com/rs/zerolog/log"

	pb "github.com/Loag/mini-stats-proto/gen/go"

	_ "github.com/tursodatabase/libsql-client-go/libsql"
	_ "modernc.org/sqlite"
)

type Metric struct {
	Name       string
	MetricType string
	Value      float64
	Timestamp  int64
}

const creatTable = "CREATE TABLE IF NOT EXISTS metrics(id INTEGER PRIMARY KEY AUTOINCREMENT, name TEXT, type TEXT, value REAL, time INT)"

type Repo struct {
	db  *sql.DB
	ctx context.Context
}

func New(conf *config.Config) *Repo {
	db, err := sql.Open("libsql", fmt.Sprintf("http://%s:%d?authToken=%s", conf.RepoPath, conf.RepoPort, conf.RepoToken))

	if err != nil {
		log.Fatal().Err(err).Msgf("failed to open db at path: %s", conf.RepoPath)
	}
	ctx := context.Background()

	exec(ctx, db, creatTable)

	return &Repo{
		db:  db,
		ctx: ctx,
	}
}

func (r *Repo) Set(input *pb.IngestRequest) error {
	stmt := "INSERT INTO metrics(name, type, value, time) VALUES(?, ?, ?, ?)"
	exec(r.ctx, r.db, stmt, input.GetName(), input.GetMetricType().String(), input.GetValue(), input.GetTime())
	return nil
}

func (r *Repo) Get(stmt string) ([]Metric, error) {
	var metrics []Metric
	rows := query(r.ctx, r.db, stmt) // "SELECT * FROM counter"
	for rows.Next() {
		var metric Metric
		if err := rows.Scan(&metric.Name, &metric.MetricType, &metric.Value, &metric.Timestamp); err != nil {
			log.Err(err).Msg("failed to scan row")
		}
		metrics = append(metrics, metric)
	}

	if err := rows.Err(); err != nil {
		log.Err(err).Msg("failed to scan row")
		return metrics, err
	}

	return metrics, nil
}

func exec(ctx context.Context, db *sql.DB, stmt string, args ...any) sql.Result {
	res, err := db.ExecContext(ctx, stmt, args...)
	if err != nil {
		log.Err(err).Msgf("failed to execute statement %s", stmt)
	}
	return res
}

func query(ctx context.Context, db *sql.DB, stmt string, args ...any) *sql.Rows {
	res, err := db.QueryContext(ctx, stmt, args...)
	if err != nil {
		log.Err(err).Msgf("failed to execute query %s", stmt)
	}
	return res
}
