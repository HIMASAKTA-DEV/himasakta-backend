# HIMASAKTA Web API

## About

<p align="center">
  <img src="docs/api_flow.png" alt="API Architecture" />
  <br>
Official backend API for HIMASAKTA (Himpunan Mahasiswa Aktuaria) ITS Web Platform. Built with Go, Gin, GORM, and PostgreSQL.
</p>

## Features

- **Superadmin Authentication**: Secure login using JWT with auto-seeded admin credentials.
- **Hintful Validation**: Technical validation errors (e.g., `required`) are translated into user-friendly Indonesian messages.
- **Modular File Storage**: Supports local filesystem (`STORAGE_TYPE=local`) and S3-compatible storage (`STORAGE_TYPE=s3`). Gallery uploads capped at 10MB.
- **Database Export**: Export all database records as SQL via `go run main.go --export` for easy migration between environments.
- **Auto-Migration**: Tables and FK constraints are auto-created on startup via `--migrate` flag (handles circular FK dependencies).
- **Analytics**: Visitor tracking and statistics.
- **CMS Entities**:
  - `CabinetInfo`: Visi, Misi, and Cabinet details.
  - `Department`: HIMASAKTA departments with social links (Instagram, YouTube, Twitter, LinkedIn, TikTok).
  - `Role`: Member ranks (Kahima, Kadep, etc.).
  - `Member`: Management of members and their roles.
  - `Progenda`: Program Kerja and Agenda management with timelines.
  - `MonthlyEvent`: Calendar of events.
  - `News`: News and articles with tags/hashtags.
  - `NrpWhitelist`: NRP verification system.
  - `GlobalSetting`: Dynamic web settings.
  - `AcademicCalendar`: Academic calendar aggregation.

## Documentation

- **Interactive API Playground**: Navigate to `/` in your browser after starting the server.
- **OpenAPI 3.0 Spec**: See [`docs/openapi.yaml`](docs/openapi.yaml) — use with Swagger UI, Postman, or `openapi-typescript-codegen` to generate typed API clients.
- **API Flow Diagram**: See [`docs/api_flow.dot`](docs/api_flow.dot) — render with `dot -Tpng docs/api_flow.dot -o docs/api_flow.png`.

## Getting Started

### Environment Variables (.env)
<details markdown="1">  

```env
# APP
APP_PORT=8080
APP_HOST=localhost
APP_URL=http://localhost:8080
APP_MODE=dev # dev or production

# DATABASE — Full URL (takes priority)
# POSTGRES_URL=postgres://user:pass@host:port/dbname?sslmode=require

# DATABASE — Individual vars (fallback)
DB_HOST=localhost
DB_USER=postgres
DB_PASS=password
DB_NAME=himasakta
DB_PORT=5432

# AUTH
ADMIN_USERNAME=admin
ADMIN_PASSWORD=admin
JWT_SECRET=your_jwt_secret

# STORAGE — "local" (default) or "s3"
STORAGE_TYPE=local

# S3 (only needed if STORAGE_TYPE=s3)
# S3_ENDPOINT=https://your-project.supabase.co/storage/v1/s3
# AWS_REGION=ap-southeast-1
# S3_BUCKET=your-bucket
# AWS_ACCESS_KEY=your-access-key
# AWS_SECRET_KEY=your-secret-key
# S3_PUBLIC_URL_PREFIX=https://your-project.supabase.co/storage/v1/object/public/your-bucket/
```
</details>

### Running Locally

```bash
go run main.go
```

### Database Management

```bash
# Run Migrations (also auto-seeds admin if not exists)
go run main.go --migrate

# Run Seeders
go run main.go --seeder

# Export database to SQL file
go run main.go --export

# Run tests
go run main.go --test
```

## API Routes (v1)

All paginated endpoints accept `?page=`, `?limit=`, `?sort=`, `?sort_by=` query parameters.
Default sorting is by `created_at DESC` (newest first), except for Members which prioritize Rank/Index.

### Authentication

- `POST /api/v1/auth/login` — Superadmin login (returns JWT token)
- `POST /api/v1/auth/validate` — Validate JWT token
- `POST /api/v1/auth/update` — Update superadmin credentials

### Cabinet Info

- `GET /api/v1/cabinet-info` — List all (paginated)
- `GET /api/v1/current-cabinet` — Get current active cabinet
- `GET /api/v1/cabinet-info/:id` — Get by ID
- `POST /api/v1/cabinet-info` — Create (🔒 superadmin)
- `PUT /api/v1/cabinet-info/:id` — Update (🔒 superadmin)
- `DELETE /api/v1/cabinet-info/:id` — Delete (🔒 superadmin)

### Department

- `GET /api/v1/department` — List all (paginated, filter: `?name=`)
- `GET /api/v1/department/:slug` — Get by slug or ID
- `POST /api/v1/department` — Create (🔒 superadmin)
- `PUT /api/v1/department/:id` — Update (🔒 superadmin)
- `DELETE /api/v1/department/:id` — Delete (🔒 superadmin)

### Role

- `GET /api/v1/role` — List all (paginated, filter: `?name=`)
- `GET /api/v1/role/:id` — Get by ID
- `POST /api/v1/role` — Create (🔒 superadmin)
- `PUT /api/v1/role/:id` — Update (🔒 superadmin)
- `DELETE /api/v1/role/:id` — Delete (🔒 superadmin)

### Member

- `GET /api/v1/member` — List all (paginated, filters: `?name=`, `?groupby=rank`)
- `GET /api/v1/member/:id` — Get by ID
- `POST /api/v1/member` — Create (🔒 superadmin)
- `PUT /api/v1/member/:id` — Update (🔒 superadmin)
- `DELETE /api/v1/member/:id` — Delete (🔒 superadmin)

### News

- `GET /api/v1/news` — List all (paginated, filters: `?search=`, `?category=`, `?title=`, `?tags=#baru,#its`)
- `GET /api/v1/news/autocompletion` — Title autocompletion (`?search=`)
- `GET /api/v1/news/:slug` — Get by slug
- `POST /api/v1/news` — Create (🔒 superadmin)
- `PUT /api/v1/news/:id` — Update (🔒 superadmin)
- `DELETE /api/v1/news/:id` — Delete (🔒 superadmin)

### Monthly Event

- `GET /api/v1/monthly-event` — List all (paginated, filter: `?title=`)
- `GET /api/v1/monthly-event/this-month` — Get events for current month
- `GET /api/v1/monthly-event/:id` — Get by ID
- `POST /api/v1/monthly-event` — Create (🔒 superadmin)
- `PUT /api/v1/monthly-event/:id` — Update (🔒 superadmin)
- `DELETE /api/v1/monthly-event/:id` — Delete (🔒 superadmin)

### Progenda

- `GET /api/v1/progenda` — List all (paginated, filters: `?search=`, `?department_id=`, `?name=`)
- `GET /api/v1/progenda/:id` — Get by ID
- `POST /api/v1/progenda` — Create with timelines (🔒 superadmin)
- `PUT /api/v1/progenda/:id` — Update with timelines (🔒 superadmin)
- `DELETE /api/v1/progenda/:id` — Delete (🔒 superadmin)
- `POST /api/v1/progenda/:id/timeline` — Add timeline entry (🔒 superadmin)
- `PUT /api/v1/progenda/timeline/:timelineId` — Update timeline (🔒 superadmin)
- `DELETE /api/v1/progenda/timeline/:timelineId` — Delete timeline (🔒 superadmin)

### Gallery

- `GET /api/v1/gallery` — List all (paginated, filter: `?caption=`)
- `GET /api/v1/gallery/:id` — Get by ID
- `POST /api/v1/gallery` — Upload image (multipart: `image`, `caption`, `category`, `department_id`, `progenda_id`, `cabinet_id`) (🔒 superadmin)
- `PUT /api/v1/gallery/:id` — Update metadata (🔒 superadmin)
- `DELETE /api/v1/gallery/:id` — Delete (🔒 superadmin)

### NRP Whitelist

- `POST /api/v1/nrp-whitelist` — Check NRP
- `GET /api/v1/nrp-whitelist` — List all (🔒 superadmin)
- `POST /api/v1/nrp-whitelist/add` — Add NRP (🔒 superadmin)
- `PUT /api/v1/nrp-whitelist/:id` — Update (🔒 superadmin)
- `DELETE /api/v1/nrp-whitelist/:nrp` — Remove by NRP (🔒 superadmin)

### Analytics

- `POST /api/v1/analytics/visit` — Record page visit
- `GET /api/v1/analytics` — Get visitor statistics

### Global Settings

- `GET /api/v1/settings/web` — Get web settings
- `PUT /api/v1/settings/web` — Update web settings (🔒 superadmin)

### Academic Calendar

- `GET /api/v1/kalender-akademik` — Get academic calendar

### Static Files

- `GET /api/static/*` — Serve uploaded files from `assets/uploads/`

## Data Structure

- All entities use **UUID** as Primary Key.
- All entities include `Timestamp` (created_at, updated_at, deleted_at).
- Media assets are referenced via `Gallery` FKs (e.g., `logo_id`, `thumbnail_id`, `photo_id`, `organigram_id`, `cabinet_id`).
- Members are linked to `Role`, `Department`, `CabinetInfo`, and `Gallery` via FKs.
- News supports many-to-many tags via `news_tags` join table.

---

:copyright: HIMASAKTA Developer Team 2026  
