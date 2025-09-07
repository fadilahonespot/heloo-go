# Go CRUD API (Netlify Functions + Echo + GORM + PostgreSQL)

Clean CRUD for `items` using Echo and GORM, deployed as a Netlify Function. Also includes a local runner.

## Routes
- `GET /api/items` → list items
- `GET /api/items/:id` → get item by id
- `POST /api/items` → create `{ "title": "..", "description": ".." }`
- `PUT /api/items` → update `{ "id": 1, "title": "..", "description": ".." }`
- `DELETE /api/items?id=1` → delete by id

## Setup
```bash
# Set your database URL
echo "export DATABASE_URL=postgres://user:pass@host:5432/dbname" >> ~/.bashrc # or zshrc
source ~/.bashrc # or zshrc

# Initialize deps
go mod tidy
```

## Local development
```bash
# Run local Echo server (not Netlify)
go run ./cmd/local
# Server: http://localhost:8080
# API:    http://localhost:8080/items
```

## Netlify
- Function: `netlify/functions/items/items.go` (Echo via AWS adapter)
- Config: `netlify.toml` maps `/api/*` → functions
- Ensure Netlify env var `DATABASE_URL` is set.

### Run Netlify dev
```bash
netlify dev
# Access under http://localhost:8888/api/items
```

## Schema
GORM auto-migrates `items` on start. `schema.sql` is optional.
