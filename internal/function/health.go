package function

import (
	"context"
	"database/sql"
	"net/http"
	"os"

	_ "github.com/lib/pq"
)

func Health(w http.ResponseWriter, req *http.Request) {
	if _, err := GetDB(req.Context(), GetDSN()); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}

	w.Write([]byte("OK"))
}

func GetDSN() string {
	return "user=" + os.Getenv("POSTGRES_USER") + " password=" + os.Getenv("POSTGRES_PASSWORD") + " host=" + os.Getenv("POSTGRES_HOST") + " port=" + os.Getenv("POSTGRES_PORT") + " dbname=" + os.Getenv("POSTGRES_DATABASE")
}

func GetDB(
	ctx context.Context,
	dsn string,
) (*sql.DB, error) {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	if db.Ping() != nil {
		return nil, err
	}

	return db, nil
}
