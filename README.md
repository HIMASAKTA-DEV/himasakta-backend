# Flexoo Academy Golang Template

## About

This is the standard backend template for Flexoo Academy projects, built with Golang. It includes setups for Docker, Database seeding, Migrations, and standard API structures.

## Getting Started

### Environment Setup

Setup environment variables for development and production.

**Development:**
```bash
cp .env.example .env
```

**Production:**
```bash
cp .env.example .env.prod
```

### Running the Application

**Local Development:**
```bash
go run main.go
```

**With Auto-Reload (Air):**
```bash
go run main.go --watch
```

## Database Management

### Migrations
```bash
go run main.go --migrate
```

### Seeding
```bash
go run main.go --seeder
```

## Docker Deployment

This project uses `Makefile` to simplify Docker operations.

### Initialization

Initialize development environment:
```bash
make init-dev
```

Initialize production environment:
```bash
make init-prod
```

### Useful Commands

| Command | Description |
|---------|-------------|
| `make up-dev` | Start development containers |
| `make down-dev` | Stop development containers |
| `make logs-dev` | View development logs |
| `make rebuild-dev` | Rebuild and restart development containers |
| `make help` | Show all available commands |

## Project Structure

- `cmd/`: Application entrypoints
- `internal/`: Private application code
- `db/`: Database migrations and seeds

## License

Flexoo Academy
