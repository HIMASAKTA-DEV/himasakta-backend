# HIMASAKTA Web API

## About

The official backend API for HIMASAKTA (Himpunan Mahasiswa Teknik Aktuaria) ITS Web Platform. Built with Go, Gin, GORM, and PostgreSQL (Supabase).

## Features

- **Superadmin Authentication**: Secure login using JWT and environment variables.
- **Centralized Gallery**: All media assets (images, logos, photos) are managed via the Gallery component and stored in Supabase S3.
- **CMS Entities**:
  - `CabinetInfo`: Visi, Misi, and Cabinet details.
  - `Department`: HIMASAKTA departments.
  - `Member`: Management of members and their roles.
  - `Progenda`: Program Kerja and Agenda management.
  - `MonthlyEvent`: Calendar of events.
  - `News`: News and articles.

## Documentation

- **Interactive API Playground**: Navigate to `/` in your browser after starting the server.
- **OpenAPI 3.0 Spec**: See [`docs/openapi.yaml`](docs/openapi.yaml) — use with Swagger UI, Postman, or `openapi-typescript-codegen` to generate typed API clients.
- **API Flow Diagram**: See [`docs/api_flow.dot`](docs/api_flow.dot) — render with `dot -Tpng docs/api_flow.dot -o docs/api_flow.png`.

## Getting Started

### Environment Variables (.env)

```env
# APP
APP_PORT=8080
APP_HOST=localhost
APP_URL=http://localhost:8080
APP_MODE=dev # dev or production

# DATABASE (Vercel/Supabase style)
POSTGRES_URL=postgres://user:pass@host:port/dbname?sslmode=require

# DATABASE (Traditional - fallback)
DB_HOST=localhost
DB_USER=postgres
DB_PASS=password
DB_NAME=himasakta
DB_PORT=5432

# AUTH
ADMIN_USERNAME=admin
ADMIN_PASSWORD=admin
JWT_SECRET=your_jwt_secret

# STORAGE (Supabase S3)
S3_ENDPOINT=https://your-project.supabase.co/storage/v1/s3
AWS_REGION=ap-southeast-1
S3_BUCKET=your-bucket
AWS_ACCESS_KEY=your-access-key
AWS_SECRET_KEY=your-secret-key
S3_PUBLIC_URL_PREFIX=https://your-project.supabase.co/storage/v1/object/public/your-bucket/
```

### Running Locally

```bash
go run main.go
```

### Database Management

```bash
# Run Migrations
go run main.go --migrate

# Run Seeders
go run main.go --seeder

# Run tests
go run main.go --test
```

## API Routes (v1)

All paginated endpoints accept `?page=`, `?limit=`, `?sort=`, `?sort_by=` query parameters.

### Authentication

- `POST /api/v1/auth/login` — Superadmin login (returns JWT token)

### Uploads

- `POST /api/v1/uploads` — Upload file to S3 (multipart/form-data, field: `file`)

### Cabinet Info

- `GET /api/v1/cabinet-info` — List all (paginated)
- `GET /api/v1/current-cabinet` — Get current active cabinet
- `GET /api/v1/cabinet-info/:id` — Get by ID
- `POST /api/v1/cabinet-info` — Create (fields: `visi`, `misi`, `description`, `tagline`, `period_start`, `period_end`, `logo_id`, `organigram_id`, `is_active`)
- `PUT /api/v1/cabinet-info/:id` — Update
- `DELETE /api/v1/cabinet-info/:id` — Delete

### Department

- `GET /api/v1/department` — List all (paginated, filter: `?name=`)
- `GET /api/v1/department/:name` — Get by name or ID
- `POST /api/v1/department` — Create (fields: `name`, `description`, `logo_id`, `social_media_link`, `bank_soal_link`, `silabus_link`, `bank_ref_link`)
- `PUT /api/v1/department/:id` — Update
- `DELETE /api/v1/department/:id` — Delete

### News

- `GET /api/v1/news` — List all (paginated, filters: `?search=`, `?category=`, `?title=`)
- `GET /api/v1/news/autocompletion` — Title autocompletion (`?search=`)
- `GET /api/v1/news/:slug` — Get by slug
- `POST /api/v1/news` — Create (fields: `title`, `tagline`, `hashtags`, `content`, `thumbnail_id`, `published_at`)
- `PUT /api/v1/news/:id` — Update
- `DELETE /api/v1/news/:id` — Delete

### Monthly Event

- `GET /api/v1/monthly-event` — List all (paginated, filter: `?title=`)
- `GET /api/v1/monthly-event/this-month` — Get events for current month
- `GET /api/v1/monthly-event/:id` — Get by ID
- `POST /api/v1/monthly-event` — Create (fields: `title`, `thumbnail_id`, `description`, `month`, `link`)
- `PUT /api/v1/monthly-event/:id` — Update
- `DELETE /api/v1/monthly-event/:id` — Delete

### Progenda

- `GET /api/v1/progenda` — List all (paginated, filters: `?search=`, `?department_id=`, `?name=`)
- `GET /api/v1/progenda/:id` — Get by ID
- `POST /api/v1/progenda` — Create (fields: `name`, `thumbnail_id`, `goal`, `description`, social links, `department_id`, `timelines[]`)
- `PUT /api/v1/progenda/:id` — Update
- `DELETE /api/v1/progenda/:id` — Delete

### Gallery

- `GET /api/v1/gallery` — List all (paginated, filter: `?caption=`)
- `GET /api/v1/gallery/:id` — Get by ID
- `POST /api/v1/gallery` — Upload image (multipart/form-data: `image`, `caption`, `category`, `department_id`, `progenda_id`, `cabinet_id`)
- `PUT /api/v1/gallery/:id` — Update metadata
- `DELETE /api/v1/gallery/:id` — Delete

### Member

- `GET /api/v1/member` — List all (paginated, filter: `?name=`)
- `GET /api/v1/member/:id` — Get by ID
- `POST /api/v1/member` — Create (fields: `nrp`, `name`, `role`, `period`, `department_id`, `photo_id`)
- `PUT /api/v1/member/:id` — Update
- `DELETE /api/v1/member/:id` — Delete

### NRP Whitelist

- `POST /api/v1/nrp-whitelist` — Check NRP (field: `nrp`)
- `POST /api/v1/nrp-whitelist/add` — Add NRP to whitelist (admin)
- `DELETE /api/v1/nrp-whitelist/:id` — Remove from whitelist (admin)

## Data Structure

- All entities use **UUID** as Primary Key.
- All entities include `Timestamp` (created_at, updated_at, deleted_at).
- Media assets are referenced via `Gallery` FKs (e.g., `logo_id`, `thumbnail_id`, `photo_id`, `organigram_id`, `cabinet_id`).

---

HIMASAKTA Developer Team
