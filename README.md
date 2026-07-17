<picture>
  <source media="(prefers-color-scheme: dark)" srcset="https://img.shields.io/badge/TubeDex-3b82f6?style=for-the-badge&logo=youtube&logoColor=white">
  <img alt="TubeDex" src="https://img.shields.io/badge/TubeDex-3b82f6?style=for-the-badge&logo=youtube&logoColor=white">
</picture>

**Your personal YouTube library.** Import, organize, and search your subscriptions — free from algorithmic feeds.

[![Go](https://img.shields.io/badge/Go-1.26+-00ADD8?style=flat&logo=go)](https://go.dev)
[![SvelteKit](https://img.shields.io/badge/SvelteKit-5-ff3e00?style=flat&logo=svelte)](https://svelte.dev)
[![PostgreSQL](https://img.shields.io/badge/PostgreSQL-17-4169E1?style=flat&logo=postgresql)](https://postgresql.org)
[![Tailwind CSS](https://img.shields.io/badge/Tailwind_CSS-4-06B6D4?style=flat&logo=tailwindcss)](https://tailwindcss.com)
[![Chi](https://img.shields.io/badge/chi-router-5B5BD6?style=flat)](https://go-chi.io)
[![License](https://img.shields.io/badge/license-MIT-green?style=flat)](LICENSE)
[![PRs](https://img.shields.io/badge/PRs-welcome-brightgreen?style=flat)](https://github.com/anomalyco/tubedex/pulls)

---

## Overview

TubeDex transforms YouTube subscriptions into a personal library. Think Letterboxd for YouTube — a fast, searchable, organized collection of the channels you actually care about.

**Core philosophy:** Your subscriptions are a library, not a feed. Own your data, organize it your way, and find what you need instantly.

### Features

- **Google Login** — One-click sign in with your Google account
- **Subscription Import** — Fetch all your YouTube subscriptions with a single sync
- **Library View** — Browse subscriptions sorted A–Z or by recently subscribed
- **Collections** — Organize channels into color-coded groups (e.g. Programming, Cooking, Design)
- **Global Search** — Instant search across channel names, descriptions, collections, and notes
- **Channel Details** — Dedicated pages with metadata, notes, ratings, and collection membership
- **Ratings** — Rate channels 1–5 stars
- **Personal Notes** — Attach private notes to any channel
- **Manual Sync** — One-click re-sync to pull the latest subscription data

---

## Tech Stack

| Layer | Technology |
|---|---|
| **Frontend** | SvelteKit 5, TypeScript, Tailwind CSS 4, TanStack Query, Lucide |
| **Backend** | Go 1.26+, chi router, pgx, sqlc |
| **Database** | PostgreSQL 17 |
| **Auth** | Google OAuth 2.0 + secure session cookies |
| **Infrastructure** | Docker Compose |

### Architecture

```
┌─────────────┐     ┌──────────────┐     ┌────────────┐
│  SvelteKit   │────▶│  Go (chi)    │────▶│ PostgreSQL │
│  Frontend    │◀────│  REST API    │◀────│            │
└─────────────┘     └──────┬───────┘     └────────────┘
                           │
                    ┌──────▼───────┐
                    │  YouTube     │
                    │  Data API    │
                    └──────────────┘
```

The API follows a clean architecture with dependency injection. Business logic lives in services; HTTP handlers stay thin. The frontend uses TanStack Query for server state management with automatic caching and background refetching.

---

## Getting Started

### Prerequisites

- [Go](https://go.dev/dl/) 1.26+
- [Bun](https://bun.sh) 1.3+
- [Docker](https://docker.com) (for PostgreSQL)
- [sqlc](https://sqlc.dev) (for database code generation)
- Google OAuth credentials ([setup guide](https://console.cloud.google.com/apis/credentials))
- YouTube Data API key ([enable here](https://console.cloud.google.com/apis/library/youtube.googleapis.com))

### 1. Clone and configure

```bash
git clone https://github.com/anomalyco/tubedex.git
cd tubedex

cp .env.example .env
```

Edit `.env` with your credentials:

```env
GOOGLE_CLIENT_ID=your-client-id
GOOGLE_CLIENT_SECRET=your-client-secret
YOUTUBE_API_KEY=your-api-key
SESSION_SECRET=generate-a-random-secret
```

### 2. Start PostgreSQL

```bash
docker compose up -d db
```

### 3. Start the API

```bash
cd apps/api
go run ./cmd/api
```

The API starts on `http://localhost:8080`. Migrations run automatically on startup.

### 4. Start the frontend

```bash
cd apps/web
bun install
bun run dev
```

The frontend starts on `http://localhost:5173`.

### 5. Sign in

Open `http://localhost:5173`, click **Sign in with Google**, and authorize the application. Navigate to **Sync** and click **Sync Now** to import your subscriptions.

---

## Docker Compose (full stack)

```bash
# Start everything
docker compose up -d

# View logs
docker compose logs -f api web

# Stop
docker compose down
```

---

## Project Structure

```
tubedex/
├── apps/
│   ├── api/                    # Go backend
│   │   ├── cmd/api/            # Entry point + route setup
│   │   ├── internal/
│   │   │   ├── auth/           # Google OAuth, sessions
│   │   │   ├── channel/        # Channel CRUD
│   │   │   ├── collection/     # Collections CRUD
│   │   │   ├── config/         # Environment config
│   │   │   ├── db/             # Migrations, sqlc queries + generated code
│   │   │   ├── middleware/     # Auth, CORS, logging, recovery
│   │   │   ├── notes/          # Channel notes
│   │   │   ├── ratings/        # Channel ratings
│   │   │   ├── search/         # Global search
│   │   │   ├── services/       # Dependency injection container
│   │   │   ├── subscription/   # Subscription management
│   │   │   ├── sync/           # YouTube subscription import
│   │   │   ├── user/           # User profile
│   │   │   └── youtube/        # YouTube Data API client
│   │   └── sqlc.yaml           # sqlc configuration
│   └── web/                    # SvelteKit frontend
│       └── src/
│           ├── lib/            # API client, stores, utilities
│           └── routes/         # Pages (dashboard, search, collections, etc.)
├── docker-compose.yml
├── Makefile
└── .env.example
```

---

## API Reference

All endpoints are versioned under `/api/v1` and return JSON.

### Authentication

| Method | Path | Description |
|---|---|---|
| GET | `/api/v1/auth/login` | Redirect to Google OAuth |
| GET | `/api/v1/auth/callback` | OAuth callback — sets session cookie |
| POST | `/api/v1/auth/logout` | Clear session |

### Protected Endpoints

| Method | Path | Description |
|---|---|---|
| GET | `/me` | Current user profile |
| GET | `/subscriptions?sort=name&limit=50&offset=0` | List subscriptions |
| DELETE | `/subscriptions?channel_id=1` | Unsubscribe |
| GET | `/channels` | List subscribed channels |
| GET | `/channels/:id` | Channel details |
| GET | `/collections` | List collections |
| POST | `/collections` | Create collection |
| PATCH | `/collections/:id` | Update collection |
| DELETE | `/collections/:id` | Delete collection |
| GET | `/collections/:id/channels` | Channels in collection |
| POST | `/collections/:id/channels` | Add channel to collection |
| DELETE | `/collections/:id/channels` | Remove channel from collection |
| POST | `/sync` | Start subscription sync |
| GET | `/sync/status` | Latest sync status |
| GET | `/search?q=query` | Global search |
| GET | `/notes?channel_id=1` | Get note |
| PUT | `/notes` | Create/update note |
| DELETE | `/notes?channel_id=1` | Delete note |
| GET | `/ratings?channel_id=1` | Get rating |
| PUT | `/ratings` | Create/update rating |
| DELETE | `/ratings?channel_id=1` | Delete rating |

---

## Development

### Regenerate sqlc code

```bash
cd apps/api
sqlc generate
```

### Run linters

```bash
# Go
cd apps/api && golangci-lint run

# TypeScript
cd apps/web && bun run check
```

### Database migrations

Migrations run automatically on API startup. They are located in `apps/api/internal/db/migrations/`.

```bash
# Create a new migration
cd apps/api
touch internal/db/migrations/000011_new_feature.up.sql
touch internal/db/migrations/000011_new_feature.down.sql
```

---

## Design Principles

- **Mobile-first** — Bottom navigation, large touch targets, responsive cards
- **Fast** — Cursor pagination, debounced search, optimized queries
- **Clean** — No animations, no glassmorphism, no gradients. Just typography, spacing, and one accent color
- **Maintainable** — Dependency injection, thin handlers, feature-based packages, no globals

---

## Future Extensions

The provider interface is designed to support additional sources beyond YouTube:

- RSS/Atom feeds
- Nebula
- Podcasts
- Twitch

Planned features: AI categorization, upload feeds, watch queues, analytics.

---

## License

[MIT](LICENSE)
