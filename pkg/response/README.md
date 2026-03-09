# Standard API Response

All API endpoints use a standardized response format for consistency.

## Response Structure

```go
{
  "status": int,      // HTTP status code
  "data": interface{}, // Response data (omitted in errors)
  "message": string    // Human-readable message
}
```

## Usage

### Success Response

```go
import "first-go-project/pkg/response"

// With data
response.Success(c, http.StatusOK, userData, "User retrieved successfully")

// With simple data
response.Success(c, http.StatusCreated, gin.H{
    "id": 123,
}, "Resource created")
```

**Output:**
```json
{
  "status": 200,
  "data": { /* your data */ },
  "message": "User retrieved successfully"
}
```

### Error Response

```go
import "first-go-project/pkg/response"

// Bad request
response.Error(c, http.StatusBadRequest, "Invalid input")

// Not found
response.Error(c, http.StatusNotFound, "Resource not found")

// Internal error
response.Error(c, http.StatusInternalServerError, "Internal server error")
```

**Output:**
```json
{
  "status": 400,
  "message": "Invalid input"
}
```

## Common Status Codes

- `200 OK` - Success
- `201 Created` - Resource created
- `400 Bad Request` - Invalid input
- `401 Unauthorized` - Authentication required
- `403 Forbidden` - Access denied
- `404 Not Found` - Resource not found
- `422 Unprocessable Entity` - Validation failed
- `500 Internal Server Error` - Server error

## Benefits

1. **Consistency** - All endpoints return the same structure
2. **Client-friendly** - Easy to parse and handle
3. **Predictable** - Clients know what to expect
4. **Maintainable** - Changes in one place affect all endpoints
