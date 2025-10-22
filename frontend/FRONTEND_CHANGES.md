# Frontend Rebuild - Gwern.net Style

## Overview
The frontend has been completely rebuilt with a gwern.net-inspired aesthetic and full userpage editing functionality.

## Key Features

### 1. Gwern.net Aesthetic
- **Dark Theme**: Deep dark background (#1a1a1a) with light gray text (#d0d0d0)
- **Typography**: Serif fonts (Georgia) for body text, sans-serif for UI elements
- **Minimalist Design**: Clean borders, subtle shadows, focused on content
- **Link Style**: Blue links (#88f) with hover effects
- **Justified Text**: Body text is justified with automatic hyphenation

### 2. Userpage Editing System

#### Edit Mode
- Click "Edit Page" button to enter edit mode
- Click "Done Editing" to exit edit mode

#### Adding Components
- In edit mode, hover over components to see + buttons on the left side (top and bottom)
- Click + to open a menu with component options:
  - **Header**: Editable header with 4 size options (h1-h4)
  - **Paragraph**: Editable paragraph text
  - **Post**: Link to a post with size options (medium/small)
  - **Divider**: Horizontal rule separator

#### Editing Components
- **Headers & Paragraphs**: Click on the text to edit inline
  - Press Enter (headers) or Escape to finish editing
  - Changes save automatically on blur
- **Posts**: 
  - When adding, first enter a URL
  - System fetches metadata automatically
  - Review and edit the post details
  - Click "Add Post" to save
  - In edit mode, use size buttons to toggle between:
    - **Medium**: Shows thumbnail, title, and description
    - **Small**: Shows title only (still clickable)

#### Drag and Drop
- In edit mode, drag components by their drag handle (⋮⋮)
- Drop to reorder components
- Visual feedback shows where component will be placed

### 3. Component Types

#### PostComponent
```typescript
{
  post: {
    post_id: string,
    size: 2 | 3  // 2=medium, 3=small
  }
}
```

#### HeaderComponent
```typescript
{
  header: {
    content: string,
    size: 1 | 2 | 3 | 4  // 1=h1, 2=h2, 3=h3, 4=h4
  }
}
```

#### ParagraphComponent
```typescript
{
  paragraph: {
    content: string
  }
}
```

#### DividerComponent
```typescript
{
  divider: {
    style: 'regular'
  }
}
```

## Technical Details

### Backend Integration
- Uses session-based authentication with Google OAuth (withCredentials: true)
- Google login flow:
  1. User clicks "Sign in with Google"
  2. Redirects to `GET /auth/google`
  3. Google OAuth flow
  4. Callback to `GET /auth/google/callback`
  5. Backend sets session cookie and redirects to `/home`
  6. Frontend checks auth status with `PUT /auth/me`
- Userpage endpoints:
  - `PUT /userpage/component/add` - Add component at index
  - `PUT /userpage/component/move` - Move component to new index
- Note: Backend doesn't have GET /userpage endpoint yet, so we start with empty state and build from mutation responses

### State Management
- React Query for server state
- Local state for edit mode and UI interactions
- Cache updates from mutation responses

### File Structure
```
frontend/src/
├── api.ts                          # API client
├── types.ts                        # TypeScript types
├── App.tsx                         # Main app component
├── index.css                       # Global gwern.net styles
├── App.css                         # App-specific styles
├── components/
│   ├── Navigation.tsx              # Top navigation bar
│   ├── PostCard.tsx                # Post display component
│   ├── EditablePostCard.tsx        # (legacy, can be removed)
│   └── userpage/
│       ├── ComponentMenu.tsx       # + button dropdown menu
│       ├── HeaderComponentView.tsx # Header renderer
│       ├── ParagraphComponentView.tsx # Paragraph renderer
│       ├── PostComponentView.tsx   # Post renderer + creator
│       └── DividerComponentView.tsx # Divider renderer
└── pages/
    ├── HomePage.tsx                # Feed page
    └── UserPage.tsx                # Userpage with editing
```

## Styling Approach
- CSS variables for theming
- Inline styles for component-specific layout
- CSS classes for reusable patterns
- Responsive design with mobile breakpoints

## Backend Endpoints Added
✅ `GET /userpage` - Retrieve user's userpage (creates empty one if doesn't exist)
✅ `PUT /userpage/component/add` - Add component at index
✅ `PUT /userpage/component/update` - Update component at index
✅ `DELETE /userpage/component/delete` - Delete component at index
✅ `PUT /userpage/component/move` - Move component from one index to another

## Fully Implemented Features
✅ Get userpage on load
✅ Add components (header, paragraph, post, divider)
✅ Edit component text inline (persists to backend)
✅ Delete components (with confirmation)
✅ Drag and drop to reorder
✅ Post size selection (medium/small)
✅ Auto-create userpage if doesn't exist

## Known Limitations
1. Size 1 (large) for posts not implemented yet
2. No public userpage viewing (only owner can see their page)

## Future Enhancements
- Add large post size (size 1) with full-width display
- Add public userpage viewing at /user/:username
- Add more component types (images, code blocks, quotes, etc.)
- Add userpage visibility settings (public/private)
- Add component styling options (colors, alignment, etc.)
- Add userpage themes
- Add component duplication

