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

## Refinements

- **Unique Identifiers**: Entities like Departments, Members, Progenda, MonthlyEvents, and News have unique `Name` or `Title` fields.
- **Specialized Queries**: All `GET` routes support exact filtering by `name`, `title`, or `period` via query parameters.
- **Supabase S3 Integration**: File uploads are automatically stored in Supabase S3 buckets with public URL generation.

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

### Authentication

- `POST /api/v1/auth/login`: Superadmin login.

### Uploads

- `POST /api/v1/uploads`: Upload file to S3. Returns `url` and `path`.

### CMS Entities (All support standard CRUD)

- `/api/v1/gallery`: Centralized media assets. Query by `?caption=`. Get by ID: `/api/v1/gallery/:id`.
- `/api/v1/cabinet-info`: Cabinet details. Query by `?period=`. Get by ID: `/api/v1/cabinet-info/:id`.
- `/api/v1/departments`: Departments. Query by `?name=`. Get by ID: `/api/v1/departments/:id`.
- `/api/v1/members`: Members. Query by `?name=`. Get by ID: `/api/v1/members/:id`.
- `/api/v1/progenda`: Programs/Agenda. Query by `?name=`, `?search=`, `?department_id=`. Get by ID: `/api/v1/progenda/:id`.
- `/api/v1/monthly-events`: Events. Query by `?title=`. Get by ID: `/api/v1/monthly-events/:id`.
- `/api/v1/news`: News. Query by `?title=`, `?search=`, `?category=`. Get by ID: `/api/v1/news/:id`.

## Data Structure (Core Entities)

- All entities use **UUID** as Primary Key.
- Most entities include a `Timestamp` (Created, Updated, Deleted).
- Media assets are referenced via `Gallery` (e.g., `LogoId`, `ThumbnailId`).

---

HIMASAKTA Developer Team
