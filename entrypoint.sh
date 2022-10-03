wait-for "${DATABASE_HOST}:${DATABASE_PORT}" -- "$@"

go run cmd/alert/main.go