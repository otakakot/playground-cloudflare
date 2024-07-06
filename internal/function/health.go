package function

import (
	"context"
	"net/http"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

func Health(w http.ResponseWriter, req *http.Request) {
	if _, err := GetPool(req.Context(), GetDSN()); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}

	w.Write([]byte("OK"))
}

func GetDSN() string {
	return "user=" + os.Getenv("POSTGRES_USER") + " password=" + os.Getenv("POSTGRES_PASSWORD") + " host=" + os.Getenv("POSTGRES_HOST") + " port=" + os.Getenv("POSTGRES_PORT") + " dbname=" + os.Getenv("POSTGRES_DATABASE")
}

func GetPool(
	ctx context.Context,
	dsn string,
) (*pgxpool.Pool, error) {
	conn, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		return nil, err
	}

	pool, err := pgxpool.NewWithConfig(ctx, conn)
	if err != nil {
		return nil, err
	}

	if err := pool.Ping(ctx); err != nil {
		return nil, err
	}

	return pool, nil
}
