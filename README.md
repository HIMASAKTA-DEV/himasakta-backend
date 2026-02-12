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

### Authentication & General
- `POST /api/v1/auth/login`: Superadmin login.
- `POST /api/v1/nrp-whitelist`: Check if NRP is allowed (Demo/Frontend helper).
- `POST /api/v1/uploads`: General file upload to S3. Returns `url` and `path`.

### Cabinet Info
- `GET /api/v1/cabinet-info`: List all cabinet info. Support `?period=`, `?page=`, `?limit=`.
- `GET /api/v1/cabinet-info/:id`: Get specific cabinet details.
- `GET /api/v1/current-cabinet`: Get active cabinet.
- `POST /api/v1/cabinet-info`: Create new cabinet info (Requires Auth).
- `PUT /api/v1/cabinet-info/:id`: Update cabinet info (Requires Auth).
- `DELETE /api/v1/cabinet-info/:id`: Delete cabinet info (Requires Auth).

### Departments
- `GET /api/v1/department`: List departments. Support `?name=`, `?page=`, `?limit=`.
- `GET /api/v1/department/:name`: Get department details by name (e.g., `/PTI`).
- `POST /api/v1/department`: Create department (Requires Auth).
- `PUT /api/v1/department/:id`: Update department (Requires Auth - uses UUID).
- `DELETE /api/v1/department/:id`: Delete department (Requires Auth - uses UUID).

### Members
- `GET /api/v1/member`: List members. Support `?name=`, `?page=`, `?limit=`.
- `GET /api/v1/member/:id`: Get member details.
- `POST /api/v1/member`: Create member (Requires Auth).
- `PUT /api/v1/member/:id`: Update member (Requires Auth).
- `DELETE /api/v1/member/:id`: Delete member (Requires Auth).

### Progenda (Program Kerja & Agenda)
- `GET /api/v1/progenda`: List progenda. Support `?search=`, `?department_id=`, `?name=`, `?page=`, `?limit=`.
- `GET /api/v1/progenda/:id`: Get progenda details.
- `POST /api/v1/progenda`: Create progenda (Requires Auth).
- `PUT /api/v1/progenda/:id`: Update progenda (Requires Auth).
- `DELETE /api/v1/progenda/:id`: Delete progenda (Requires Auth).

### Monthly Events
- `GET /api/v1/monthly-event`: List events. Support `?title=`, `?page=`, `?limit=`.
- `GET /api/v1/monthly-event/this-month`: Get events for the current month.
- `GET /api/v1/monthly-event/:id`: Get event details.
- `POST /api/v1/monthly-event`: Create event (Requires Auth).
- `PUT /api/v1/monthly-event/:id`: Update event (Requires Auth).
- `DELETE /api/v1/monthly-event/:id`: Delete event (Requires Auth).

### News Articles
- `GET /api/v1/news`: List news. Support `?search=`, `?page=`, `?limit=`.
- `GET /api/v1/news/:slug`: Get news article by slug (e.g., `/judul-berita`).
- `POST /api/v1/news`: Create news article (Requires Auth - generates slug).
- `PUT /api/v1/news/:id`: Update news article (Requires Auth - uses UUID).
- `DELETE /api/v1/news/:id`: Delete news article (Requires Auth - uses UUID).

### Gallery (Media Management)
- `GET /api/v1/gallery`: List gallery items. Support `?caption=`, `?page=`, `?limit=`, `?filter_by=category`.
- `GET /api/v1/gallery/:id`: Get gallery item details.
- `POST /api/v1/gallery`: Upload image and create gallery entry (Requires Auth).
- `PUT /api/v1/gallery/:id`: Update gallery metadata (Requires Auth).
- `DELETE /api/v1/gallery/:id`: Delete gallery entry (Requires Auth).

> [!TIP]
> All list endpoints (`GET`) support generic filtering via `?filter_by=field&filter=value`.

## Tutorial (Nextjs) Frontend  
> Cara pakai fetch() di nextjs: [klik ygy](https://nextjs.org/docs/app/api-reference/functions/fetch)
1. Cara mendapatkan informasi kabinet sekarang:
```javascript
const res = await fetch(`${API_URL}/api/v1/current-cabinet`);
const { data } = await res.json();
// data berisi: visi, misi, tagline, logo { url }, description, dll.
```
2. Cara mendapatkan gambar organigram sekarang:
```javascript
// Organigram terhubung dengan cabinet-info dan sedang dikerjakan
```
3. Cara mendapatkan acara bulanan (get to know):
```javascript
const res = await fetch(`${API_URL}/api/v1/monthly-event/this-month`);
const { data } = await res.json();
```
4. Cara mendapatkan 12 berita terbaru:
```javascript
const res = await fetch(`${API_URL}/api/v1/news?limit=12`);
const { data } = await res.json();
```
5. Cara mendapatkan informasi departemen x (nama):
```javascript
// Contoh departemen Kaderisasi
const res = await fetch(`${API_URL}/api/v1/department/Kaderisasi`);
const { data } = await res.json();
```
6. Cara mendapatkan list progenda dari departemen x:
```javascript
// Step 1: Dapatkan id dari departemen
const deptRes = await fetch(`${API_URL}/api/v1/department/Kaderisasi`);
const { data: dept } = await deptRes.json();

// Step 2: Dapatkan progenda berdasarkan department_id
const res = await fetch(`${API_URL}/api/v1/progenda?department_id=${dept.id}`);
const { data: progendus } = await res.json();
```
7. Cara mendapatkan galeri dari departemen x:
```javascript
// deptId itu dept.id atau id departemen
const res = await fetch(`${API_URL}/api/v1/gallery?department_id=${deptId}`);
const { data } = await res.json();
```
8. Cara mendapatkan informasi berita / detail berita dari slug:
```javascript
// slug itu judul berita yang sudah diubah menjadi format url (e.g. /berita/judul-berita)
const res = await fetch(`${API_URL}/api/v1/news/${slug}`);
const { data } = await res.json();
```
9. Cara mendapatkan informasi progenda dari id progenda:
```javascript
// id progenda didapat dari di nomor 6 step 2, misal res[0].id
const res = await fetch(`${API_URL}/api/v1/progenda/${progendaId}`);
const { data } = await res.json();
```
10. Cara verifikasi NRP:
```javascript
// untuk create, delete nrp sedang dikerjakan
const res = await fetch(`${API_URL}/api/v1/nrp-whitelist`, {
    method: 'POST',
    headers: {
        'Content-Type': 'application/json',
    },
    body: JSON.stringify({
        nrp: '1234567890',
    }),
});
const { data } = await res.json();
```
11. Cara login admin dan mendapatkan JWT:
```javascript
// usn sama pw nya admin semua
const res = await fetch(`${API_URL}/api/v1/auth/login`, {
    method: 'POST',
    headers: {
        'Content-Type': 'application/json',
    },
    body: JSON.stringify({
        username: 'admin',
        password: 'admin',
    }),
});
const { data } = await res.json();
```
12. Cara upload ke galeri:
```javascript
// token adalah jwt yang didapat dari nomor 11 step 1
const res = await fetch(`${API_URL}/api/v1/gallery`, {
    method: 'POST',
    headers: {
        'Content-Type': 'application/json',
        'Authorization': `Bearer ${token}`,
    },
    body: JSON.stringify({
        caption: 'caption',
        department_id: 'department_id',
    }),
});
const { data } = await res.json();
```
13. Cara membuat berita:
```javascript
// department_id didapat dari nomor 5 step 1
// thumbnail_id didapat dari nomor 12 step 1 (galery_id)
const res = await fetch(`${API_URL}/api/v1/news`, {
    method: 'POST',
    headers: {
        'Content-Type': 'application/json',
        'Authorization': `Bearer ${token}`,
    },
    body: JSON.stringify({
        title: 'title',
        content: 'content',
        department_id: '1234',
        thumbnail_id: '1234',
        hashtags: "tag1,tag2"
    }),
});
const { data } = await res.json();
```

### LIST FULL API DI: 
> https://himasakta-backend.vercel.app

## Data Structure (Core Entities)
- All entities use **UUID** as Primary Key.
- Most entities include a `Timestamp` (Created, Updated, Deleted).
- Media assets are referenced via `Gallery` (e.g., `LogoId`, `ThumbnailId`).

---
HIMASAKTA Developer Team
