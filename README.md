# RSS Aggregator

This project is an RSS Aggregator built using Go. It provides a RESTful API for managing users, feeds, posts, and feed follows. The application uses PostgreSQL as the database and includes features such as user authentication and periodic scraping of RSS feeds.

## Features

- User management (create, retrieve users)
- Feed management (create, retrieve feeds)
- Post management (retrieve posts for users)
- Feed follows (follow/unfollow feeds)
- Health check endpoint
- Middleware for authentication
- CORS support
- Periodic scraping of RSS feeds

## Project Structure

```
RSS/
├── main.go                # Entry point of the application
├── handler_user.go        # Handlers for user-related endpoints
├── handler_feed.go        # Handlers for feed-related endpoints
├── handler_feed_follows.go # Handlers for feed follow-related endpoints
├── internal/
│   ├── auth/             # Authentication logic
│   ├── database/         # Database queries and models
├── sql/                  # SQL schema and queries
│   ├── queries/          # SQL query files
│   ├── schema/           # Database schema files
└── vendor/               # Vendor dependencies
```

## Endpoints

### User Endpoints
- `POST /v1/users`: Create a new user.
- `GET /v1/users`: Retrieve user details (requires authentication).

### Feed Endpoints
- `POST /v1/feeds`: Create a new feed (requires authentication).
- `GET /v1/feeds`: Retrieve a list of feeds.

### Post Endpoints
- `GET /v1/posts`: Retrieve posts for the authenticated user.

### Feed Follow Endpoints
- `POST /v1/feed_follows`: Follow a feed (requires authentication).
- `GET /v1/feed_follows`: Retrieve the list of followed feeds (requires authentication).
- `DELETE /v1/feed_follows/{feedFollowID}`: Unfollow a feed (requires authentication).

### Health Check
- `GET /v1/healthz`: Check the health of the service.

## Environment Variables

The application requires the following environment variables:

- `PORT`: The port on which the server will run.
- `DB_URL`: The connection string for the PostgreSQL database.

