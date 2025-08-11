# Sane Discourse Social Media - API Documentation

## Base URL
```
http://localhost:3000
```

## Data Models

### User
```json
{
  "id": "ObjectID",
  "username": "string"
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

### User Endpoints

#### User Login
```http
PUT /user/login
```

**Request Body:**
```json
{
  "username": "string"
}
```

**Response:**
```json
{
  "id": "ObjectID",
  "username": "string"
}
```

### Post Endpoints

#### Create Posts from URLs
```http
PUT /user/posts/create
```

**Description:** Creates posts by scraping metadata from provided URLs.

**Request Body:**
```json
{
  "urls": ["string"]
}
```

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

#### Add Post Manually
```http
PUT /user/posts/add
```

**Description:** Adds a pre-formed post with user association to the Database

**Request Body:**
```json
{
  "user_id": "ObjectID",
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
POST /user/posts
```

**Description:** Retrieves posts for a specific user.

**Request Body:**
```json
{
  "user_id": "ObjectID"
}
```

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