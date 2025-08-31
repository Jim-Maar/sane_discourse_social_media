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

### Userpage Endpoints

#### Add Component to Userpage
```http
PUT /userpage/component/add
```

**Description:** Adds a new component to a userpage at a specific index.

**Request Body:**
```json
{
  "userpage_id": "ObjectID",
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
  "components": [
    // Array of updated components
  ]
}
```

#### Move Component in Userpage
```http
PUT /userpage/component/move
```

**Description:** Moves a component from one position to another within a userpage.

**Request Body:**
```json
{
  "userpage_id": "ObjectID",
  "prev_index": "number",
  "new_index": "number"
}
```

**Response:**
```json
{
  "id": "ObjectID",
  "components": [
    // Array of updated components
  ]
}
```