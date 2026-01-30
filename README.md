# CRS Backend

## Getting Started

setup env for prod and dev

dev

```bash
.env
```

prod

```bash
.env.prod
```

### Run program

```bash
go run main.go
```

### Run migrate

```bash
go run main.go --migrate
```

### Run seeder

```bash
go run main.go --seeder
```

### Run Air (auto loading for development)

```bash
go run main.go --watch
```

### Deploy Using Docker

You can run it using the Makefile. For instructions, refer to the 'run' guide.

#### Deploy

initialize dev deployment

```bash
make init-[prod/dev]
```

#### Redeploy

```bash
make rebuild-[prod/dev]
```

### Additional

You can run it using the Makefile. For instructions, refer to the 'run' guide.

```bash
make help
```
