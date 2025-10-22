# Sane Discourse Social Media - API Documentation

## Base URL
```
http://localhost:3000
```

## Authentication

The API uses session-based authentication with OAuth providers (Google).

### Google OAuth Authentication

#### Initiate Google Login
```http
GET /auth/google
```

**Description:** Redirects to Google OAuth login page. After successful authentication, redirects back to the callback URL.

**Response:** HTTP 302 redirect to Google OAuth

#### OAuth Callback (Automatic)
```http
GET /auth/google/callback
```

**Description:** Handles the OAuth callback from Google. Creates or logs in the user, sets session cookie, and redirects to `http://localhost:5173/home`.

**Response:** HTTP 302 redirect to frontend home page with session cookie set

### Get Current User
```http
PUT /auth/me
```

**Description:** Returns the currently authenticated user based on session cookie. Requires authentication.

**Response:**
```json
{
  "id": "ObjectID",
  "username": "string",
  "email": "string"
}
```

**Error Response (401):** User not authenticated

---

### Mock Authentication (Testing/Development Only)

**Note:** This endpoint is intended for automated testing only and should not be used in production.

```http
PUT /auth/login
```

**Request Body:**
```json
{
  "name": "string",
  "email": "string"
}
```

**Response:**
```json
{
  "id": "ObjectID",
  "username": "string",
  "email": "string"
}
```

## Data Models

### User
```json
{
  "id": "ObjectID",
  "username": "string",
  "email": "string"
}
```

### Post
```json
{
  "id": "ObjectID",
  "title": "string",
  "description": "string",
  "thumbnail_url": "string",
  "site_name": "string",
  "url": "string",
  "type": "string",
  "author": "string"
}
```

### Reaction
```json
{
  "id": "ObjectID",
  "reaction_type": "string",
  "user_id": "ObjectID",
  "post_id": "ObjectID"
}
```

### Userpage
```json
{
  "id": "ObjectID",
  "components": [
    // Array of different component types (see Component Types below)
  ]
}
```

### Component Types

Components are polymorphic - each type has a different structure:

#### PostComponent
```json
{
  "post_id": "ObjectID",
  "size": 1|2|3  // 1=large, 2=medium, 3=small
}
```

#### HeaderComponent
```json
{
  "content": "string",
  "size": 1|2|3|4  // 1=large, 2=medium, 3=small, 4=very small
}
```

#### ParagraphComponent
```json
{
  "content": "string"
}
```

#### DividerComponent
```json
{
  "type": "regular"
}
```

**Example Userpage with Mixed Components:**
```json
{
  "id": "507f1f77bcf86cd799439012",
  "components": [
    {
      "content": "Welcome to My Page",
      "size": 1
    },
    {
      "content": "This is a description paragraph."
    },
    {
      "post_id": "507f1f77bcf86cd799439013",
      "size": 2
    },
    {
      "type": "regular"
    }
  ]
}
```

### Reaction Types
- `agree`
- `strong_agree`
- `disagree`
- `strong_disagree`
- `important`
- `strong_important`
- `unimportant`
- `strong_unimportant`
- `upvote`
- `strong_upvote`
- `downvote`
- `strong_downvote`

## API Endpoints

### Post Endpoints

#### Create Post from URL
```http
PUT /user/posts/create
```

**Description:** Creates a post by scraping metadata from a provided URL.

**Request Body:**
```json
{
  "url": "string"
}
```

**Response:**
```json
{
  "id": "ObjectID",
  "title": "string",
  "description": "string",
  "thumbnail_url": "string",
  "site_name": "string",
  "url": "string",
  "type": "string",
  "author": "string"
}
```

#### Add Post Manually
```http
PUT /user/posts/add
```

**Description:** Adds a pre-formed post with user association to the Database. Requires authentication.

**Request Body:**
```json
{
  "post": {
    "title": "string",
    "description": "string",
    "thumbnail_url": "string",
    "site_name": "string",
    "url": "string",
    "type": "string",
    "author": "string"
  }
}
```

**Response:**
```json
{
  "id": "ObjectID",
  "title": "string",
  "description": "string",
  "thumbnail_url": "string",
  "site_name": "string",
  "url": "string",
  "type": "string",
  "author": "string"
}
```

#### Get User Posts
```http
GET /user/posts
```

**Description:** Retrieves posts for the currently authenticated user. Requires authentication.

**Response:**
```json
[
  {
    "id": "ObjectID",
    "title": "string",
    "description": "string",
    "thumbnail_url": "string",
    "site_name": "string",
    "url": "string",
    "type": "string",
    "author": "string"
  }
]
```

#### Get Feed (Home)
```http
GET /home
```

**Description:** Retrieves the main feed of posts.

**Response:**
```json
[
  {
    "id": "ObjectID",
    "title": "string",
    "description": "string",
    "thumbnail_url": "string",
    "site_name": "string",
    "url": "string",
    "type": "string",
    "author": "string"
  }
]
```

### Userpage Endpoints

#### Get Userpage
```http
GET /userpage
```

**Description:** Retrieves the authenticated user's userpage. If no userpage exists, creates and returns an empty one. Requires authentication.

**Response:**
```json
{
  "id": "ObjectID",
  "user_id": "ObjectID",
  "components": [
    // Array of components
  ]
}
```

#### Add Component to Userpage
```http
PUT /userpage/component/add
```

**Description:** Adds a new component to the authenticated user's userpage at a specific index. Requires authentication.

**Request Body:**
```json
{
  "index": "number",
  "component": {
    // Component object (PostComponent, HeaderComponent, ParagraphComponent, or DividerComponent)
    // See Component Types section for structure
  }
}
```

**Response:**
```json
{
  "id": "ObjectID",
  "user_id": "ObjectID",
  "components": [
    // Array of updated components
  ]
}
```

#### Update Component in Userpage
```http
PUT /userpage/component/update
```

**Description:** Updates an existing component at a specific index in the authenticated user's userpage. Requires authentication.

**Request Body:**
```json
{
  "index": "number",
  "component": {
    // Updated component object
  }
}
```

**Response:**
```json
{
  "id": "ObjectID",
  "user_id": "ObjectID",
  "components": [
    // Array of updated components
  ]
}
```

#### Delete Component from Userpage
```http
DELETE /userpage/component/delete
```

**Description:** Deletes a component at a specific index from the authenticated user's userpage. Requires authentication.

**Request Body:**
```json
{
  "index": "number"
}
```

**Response:**
```json
{
  "id": "ObjectID",
  "user_id": "ObjectID",
  "components": [
    // Array of updated components
  ]
}
```

#### Move Component in Userpage
```http
PUT /userpage/component/move
```

**Description:** Moves a component from one position to another within the authenticated user's userpage. Requires authentication.

**Request Body:**
```json
{
  "prev_index": "number",
  "new_index": "number"
}
```

**Response:**
```json
{
  "id": "ObjectID",
  "user_id": "ObjectID",
  "components": [
    // Array of updated components
  ]
}
```