# Event Schedule API

A RESTful API for managing events and user registrations built with Go. This is a final project from the Udemy course "Go - The Complete Guide".

## Description

The Event Schedule API allows users to:
- Create, read, update, and delete events
- Register and unregister for events
- Authenticate using JWT tokens
- Secure event operations (only event owners can modify/delete their events)

## Technology Stack

- **Go 1.24.0** - Programming language
- **Gin** - HTTP web framework
- **SQLite** - Lightweight database
- **JWT** - Authentication using golang-jwt/jwt
- **Swagger/OpenAPI** - API documentation
- **bcrypt** - Password hashing

## Project Structure

```
.
├── main.go              # Application entry point
├── db/                  # Database initialization
├── models/              # Data models (Event, User)
├── routes/              # API route handlers
├── middlewares/         # Authentication middleware
├── utils/               # Utility functions (JWT, hashing)
└── docs/                # Swagger documentation
```

## Prerequisites

- Go 1.24.0 or higher
- Git (optional, for cloning)

## Step-by-Step: Running the Project

### 1. Clone or Navigate to the Project Directory

```bash
cd /path/to/events-schedule-api
```

### 2. Install Dependencies

```bash
go mod download
```

### 3. Build the Application (Optional)

```bash
go build -o events-api
```

### 4. Run the Application

**Option A: Run directly with Go**
```bash
go run main.go
```

**Option B: Run the compiled binary**
```bash
./events-api
```

### 5. Verify the Server is Running

The server will start on `http://localhost:9090`

You should see output similar to:
```
[GIN-debug] Listening and serving HTTP on :9090
```

## How to Access Swagger Documentation

Once the server is running, you can access the interactive API documentation:

### Swagger UI URL

Open your browser and navigate to:
```
http://localhost:9090/docs/
```

### What You Can Do in Swagger

1. **View All Endpoints** - See all available API endpoints organized by tags (auth, events, registrations)

2. **Test API Endpoints** - Click on any endpoint to expand it and see:
   - Required parameters
   - Request body schema
   - Response examples
   - Try it out directly from the browser

3. **Authenticate** - For protected endpoints:
   - First, use the `/signup` endpoint to create a user
   - Then use `/login` to get a JWT token
   - Click the "Authorize" button at the top
   - Enter: `Bearer <your-jwt-token>`
   - Now you can test protected endpoints

## API Endpoints Overview

### Authentication
- `POST /signup` - Register a new user
- `POST /login` - Login and receive JWT token

### Events (Public)
- `GET /events` - Get all events
- `GET /events/:id` - Get a specific event

### Events (Protected - Requires Authentication)
- `POST /events` - Create a new event
- `PUT /events/:id` - Update an event (owner only)
- `DELETE /events/:id` - Delete an event (owner only)

### Registration (Protected - Requires Authentication)
- `POST /events/:id/register` - Register for an event
- `DELETE /events/:id/register` - Cancel registration

## Example API Usage

### 1. Create a User
```bash
curl -X POST http://localhost:9090/signup \
  -H "Content-Type: application/json" \
  -d '{"email":"user@example.com","password":"password123"}'
```

### 2. Login
```bash
curl -X POST http://localhost:9090/login \
  -H "Content-Type: application/json" \
  -d '{"email":"user@example.com","password":"password123"}'
```

Response will include a JWT token:
```json
{"message":"Login successful","token":"eyJhbGciOiJIUzI1NiIs..."}
```

### 3. Create an Event (with authentication)
```bash
curl -X POST http://localhost:9090/events \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer <your-token>" \
  -d '{
    "name":"Tech Conference 2026",
    "description":"Annual tech conference",
    "location":"San Francisco",
    "dateTime":"2026-06-15T09:00:00Z"
  }'
```

### 4. Get All Events
```bash
curl http://localhost:9090/events
```

### 5. Register for an Event
```bash
curl -X POST http://localhost:9090/events/1/register \
  -H "Authorization: Bearer <your-token>"
```

## Database

The application uses SQLite with the database file `api.db` created automatically in the project root on first run.

## Development

### Regenerate Swagger Documentation

If you modify the API annotations in the code:

```bash
# Install swag CLI (if not already installed)
go install github.com/swaggo/swag/cmd/swag@latest

# Generate docs
swag init
```

### Run Tests

```bash
go test ./...
```

## Security Features

- Password hashing using bcrypt
- JWT-based authentication
- Protected routes require valid JWT tokens
- Event ownership validation (users can only modify their own events)

## Troubleshooting

**Port already in use:**
- Change the port in `main.go` line 38: `server.Run(":9090")`

**Database locked:**
- Stop all running instances of the application
- Delete `api.db` and restart (will recreate the database)

**Cannot access Swagger:**
- Ensure the server is running
- Check the URL is exactly `http://localhost:9090/docs/`
- Check browser console for errors

## License

MIT
