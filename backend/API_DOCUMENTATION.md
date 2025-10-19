# Sane Discourse Social Media - API Documentation

## Base URL
```
http://localhost:3000
```

## Authentication

The API uses session-based authentication with OAuth providers (Google). For testing purposes, a mock authentication endpoint is available.

### Mock Authentication (Testing Only)
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

### Get Current User
```http
PUT /auth/me
```

**Description:** Returns the currently authenticated user.

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