# Pairswap's backend

1. Run docker:
```bash
docker compose up -d
```

2. Run database migration:
```bash
migrate -path migrations -database "mysql://root:password@tcp(127.0.0.1:3306)/proxy_api" -verbose up
```

3. Start server:
```bash
go run cmd/server/main.go
```
