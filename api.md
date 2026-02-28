# HIMASAKTA Backend API Reference

> **Base URL**: `https://himasakta-backend.vercel.app/api/v1`
> Tested live on 2026-02-17. All response structures below are from **real executed requests**, not stale examples.

---

## Common Response Envelope

Every response follows this structure:

```json
{
  "success": true,
  "message": "...",
  "data": T,            // single object or array
  "meta": { ... }       // only on paginated list endpoints
}
```

### Meta (Pagination)

```json
{
  "limit": 10,
  "page": 1,
  "total_data": 6,
  "total_page": 1,
  "sort": "asc",
  "sort_by": "id"
}
```

> [!IMPORTANT]
> Backend uses `total_data` / `total_page`, **not** `total` / `total_pages`.

### Common Base Fields (GORM)

All entities include these GORM standard fields:

| Field | Type | Notes |
|---|---|---|
| `created_at` | `string` (ISO 8601) | |
| `updated_at` | `string` (ISO 8601) | |
| `DeletedAt` | `string \| null` | Soft-delete timestamp |
| `id` | `string` (UUID) | |

---

## 🔑 Authentication

### `POST /auth/login`

🔒 Public

```json
// Request
{ "username": "admin", "password": "admin" }

// Response
{
  "success": true,
  "message": "login success",
  "data": {
    "token": "eyJhbGci..."
  }
}
```

> Token goes in `Authorization: Bearer <token>` header for protected endpoints.

---

## 🏛 Cabinet Info

### Real Structure

```json
{
  "id": "uuid",
  "visi": "string",
  "misi": "string",
  "description": "string",
  "tagline": "string",
  "period_start": "2024-01-01T00:00:00Z",
  "period_end": "2024-12-31T00:00:00Z",
  "logo_id": "uuid",
  "logo": { Media },
  "organigram_id": "uuid | null",
  "organigram": "{ Media } | null",
  "is_active": true
}
```

> [!WARNING]
> Uses `period_start` + `period_end` (ISO dates), **NOT** a single `period` string.
> Also has `description` and `organigram_id`/`organigram` fields.

### Endpoints

| Method | Path | Auth | Description |
|---|---|---|---|
| `GET` | `/cabinet-info` | ❌ | List all cabinet infos (paginated) |
| `GET` | `/cabinet-info/:id` | ❌ | Get single cabinet info |
| `GET` | `/current-cabinet` | ❌ | Get active cabinet (`is_active: true`) — returns single object |
| `POST` | `/cabinet-info` | ✅ | Create cabinet info |
| `PUT` | `/cabinet-info/:id` | ✅ | Update cabinet info |
| `DELETE` | `/cabinet-info/:id` | ✅ | Delete cabinet info |

---

## 🏢 Department

### Real Structure

```json
{
  "id": "uuid",
  "name": "Kaderisasi",
  "description": "...",
  "logo_id": "uuid",
  "logo": { Media },
  "social_media_link": "",
  "bank_soal_link": "",
  "silabus_link": "",
  "bank_ref_link": "https://linktr.ee/himasakta"
}
```

> [!WARNING]
> Has `social_media_link`, `bank_soal_link`, `silabus_link`, `bank_ref_link` — none of these are in the current frontend types.

### Endpoints

| Method | Path | Auth | Description |
|---|---|---|---|
| `GET` | `/department` | ❌ | List all departments (paginated) |
| `GET` | `/department/:name` | ❌ | Get by name (slug-like, e.g. `Kaderisasi`) |
| `POST` | `/department` | ✅ | Create department |
| `PUT` | `/department/:id` | ✅ | Update department |
| `DELETE` | `/department/:id` | ✅ | Delete department |

---

## 📰 News

### Real Structure

```json
{
  "id": "uuid",
  "title": "Penerimaan Anggota Baru",
  "slug": "penerimaan-anggota-baru",
  "tagline": "Ayo bergabung!",
  "hashtags": "OPREC,HIMASAKTA",
  "content": "# Selamat Datang\nKami membuka pendaftaran anggota baru.",
  "thumbnail_id": "uuid",
  "thumbnail": { Media },
  "published_at": "2026-02-12T14:12:11.272165Z"
}
```

### Endpoints

| Method | Path | Auth | Description |
|---|---|---|---|
| `GET` | `/news` | ❌ | List all news (paginated, supports `?page=&limit=&search=`) |
| `GET` | `/news/:slug` | ❌ | Get single news by slug |
| `POST` | `/news` | ✅ | Create news |
| `PUT` | `/news/:id` | ✅ | Update news |
| `DELETE` | `/news/:id` | ✅ | Delete news |

### `GET /news/autocompletion?search=<keyword>` *(UNDOCUMENTED)*

❌ No Auth required

```json
{
  "success": true,
  "message": "success get autocompletion",
  "data": ["Penerimaan Anggota Baru"]
}
```

> [!CAUTION]
> Returns `data: string[]` — an array of **plain strings** (titles only).
> Current `NewsAutocompletion` type `{id, title, thumbnail}` is **WRONG**.

---

## 📅 Monthly Event

### Real Structure

```json
{
  "id": "uuid",
  "title": "HIMASAKTA Cup",
  "thumbnail_id": "uuid",
  "thumbnail": { Media },
  "description": "",
  "month": "2026-02-12T00:00:00Z",
  "link": "https://youtube.com"
}
```

> [!NOTE]
> `month` is an ISO 8601 date string (not just `"February"`).
> Has `thumbnail_id`/`thumbnail` linking to Media.

### Endpoints

| Method | Path | Auth | Description |
|---|---|---|---|
| `GET` | `/monthly-event` | ❌ | List all monthly events (paginated) |
| `GET` | `/monthly-event/this-month` | ❌ | Get events for current month (no pagination/meta) |
| `GET` | `/monthly-event/:id` | ❌ | Get single event |
| `POST` | `/monthly-event` | ✅ | Create monthly event |
| `PUT` | `/monthly-event/:id` | ✅ | Update monthly event |
| `DELETE` | `/monthly-event/:id` | ✅ | Delete monthly event |

---

## 📋 Progenda

### Real Structure

```json
{
  "id": "uuid",
  "name": "Updated Name",
  "thumbnail_id": "uuid",
  "thumbnail": { Media },
  "goal": "Goal tes",
  "description": "",
  "website_link": "",
  "instagram_link": "",
  "twitter_link": "",
  "linkedin_link": "",
  "youtube_link": "",
  "department_id": "uuid",
  "department": { Department },
  "timelines": [
    {
      "id": "uuid",
      "progenda_id": "uuid",
      "date": "2026-01-19T00:00:00Z",
      "info": "dicoba",
      "link": "https"
    }
  ],
  "feeds": [
    {
      "id": "uuid",
      "image_url": "https://...",
      "caption": "",
      "category": "feeds",
      "department_id": null,
      "progenda_id": "uuid"
    }
  ]
}
```

> [!WARNING]
> **Major differences from current frontend types:**
> - Has `timelines: Timeline[]` (array of objects), **NOT** a `timeline: string`.
> - Has `feeds: Media[]` (array of gallery/media items).
> - Has `instagram_link`, `twitter_link`, `linkedin_link`, `youtube_link` social links.
> - `website_link` replaces the old simple `website_link`.
> - Includes nested `department` object.

### Endpoints

| `POST` | `/progenda` | ✅ | Create progenda |
| `PUT` | `/progenda/:id` | ✅ | Update progenda |
| `DELETE` | `/progenda/:id` | ✅ | Delete progenda |

#### Individual Timeline CRUD
| Method | Path | Auth | Description |
|---|---|---|---|
| `POST` | `/progenda/:id/timeline` | ✅ | Add individual timeline to progenda |
| `PUT` | `/progenda/timeline/:timelineId` | ✅ | Update individual timeline |
| `DELETE` | `/progenda/timeline/:timelineId` | ✅ | Delete individual timeline |

---

## 🖼 Gallery (Media)

### Real Structure

```json
{
  "id": "uuid",
  "image_url": "https://...",
  "caption": "Logo HIMASAKTA",
  "category": "logo",
  "department_id": "uuid | null",
  "progenda_id": "uuid | null"
}
```

> [!WARNING]
> No `title` or `description` fields — only `image_url`, `caption`, `category`.
> Has `progenda_id` to link media to progenda feeds.
> **Max file size is 10MB.**

### Endpoints

| Method | Path | Auth | Description |
|---|---|---|---|
| `GET` | `/gallery` | ❌ | List all gallery items (paginated) |
| `GET` | `/gallery/:id` | ❌ | Get single gallery item |
| `POST` | `/gallery` | ✅ | Upload image (multipart/form-data, **max 10MB**: `image`, `caption`, `category`) |
| `PUT` | `/gallery/:id` | ✅ | Update gallery item |
| `DELETE` | `/gallery/:id` | ✅ | Delete gallery item |

---

## 👥 Member

### Real Structure

```json
{
  "id": "uuid",
  "nrp": "12345678",
  "name": "John Doe",
  "role": "Ketua Departemen",
  "department_id": "uuid",
  "department": { Department },
  "photo_id": "uuid",
  "photo": { Media },
  "period": "2024"
}
```

### Request Body (POST/PUT)

```json
{
  "nrp": "123456",
  "name": "John Doe",
  "role": "Staff",
  "period": "2023-2024",
  "department_id": "uuid",
  "photo_id": "uuid"
}
```

### Endpoints

| Method | Path | Auth | Description |
|---|---|---|---|
| `GET` | `/member` | ❌ | List all members (paginated, supports `?search=&department_id=`) |
| `GET` | `/member/:id` | ❌ | Get single member |
| `POST` | `/member` | ✅ | Create member |
| `PUT` | `/member/:id` | ✅ | Update member |
| `DELETE` | `/member/:id` | ✅ | Delete member |

---

## Media Type Reference

The `Media` / `Gallery` entity is shared across all resources:

```typescript
interface Media {
  id: string;
  image_url: string;
  caption: string;
  category: string;        // "logo" | "thumbnail" | "feeds" | etc.
  department_id: string | null;
  progenda_id: string | null;
  created_at: string;
  updated_at: string;
  DeletedAt: string | null;
}
```

---

## Frontend Type Discrepancies Summary

| Type | Issue |
|---|---|
| `Meta` | Uses `total`/`total_pages` → should be `total_data`/`total_page` + `sort`/`sort_by` |
| `CabinetInfo` | Missing `period_start`/`period_end`/`description`/`organigram_id`; has wrong `period` field |
| `Department` | Missing `social_media_link`/`bank_soal_link`/`silabus_link`/`bank_ref_link` |
| `Gallery` | Has `title`/`description` → should be `image_url`/`caption`/`category`/`progenda_id` |
| `Progenda` | Has `timeline: string` → should be `timelines: Timeline[]` + `feeds: Media[]` + social links |
| `MonthlyEvent` | Missing `thumbnail_id`/`thumbnail`; `month` is ISO date not string |
| `NewsAutocompletion` | Is `{id,title,thumbnail}` → should be just `string` (plain title) |
| `Media` | Missing `department_id`/`progenda_id` |
