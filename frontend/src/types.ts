export interface User {
    id: string;
    username: string;
    email?: string;
}

export interface Post {
    id?: string;
    title: string;
    description: string;
    thumbnail_url: string;
    site_name: string;
    url: string;
    type: string;
    author: string;
}

export interface Reaction {
    id: string;
    reaction_type: string;
    user_id: string;
    post_id: string;
}

export type ReactionType =
    | 'agree'
    | 'strong_agree'
    | 'disagree'
    | 'strong_disagree'
    | 'important'
    | 'strong_important'
    | 'unimportant'
    | 'strong_unimportant'
    | 'upvote'
    | 'strong_upvote'
    | 'downvote'
    | 'strong_downvote';

// Userpage Component Types (matching backend structure)
export interface PostComponentData {
    post_id: string;
    size: 1 | 2 | 3; // 1=large (not implemented), 2=medium (thumbnail+title+desc), 3=small (title only)
}

export interface HeaderComponentData {
    content: string;
    size: 1 | 2 | 3 | 4; // 1=large, 2=medium, 3=small, 4=very small
}

export interface ParagraphComponentData {
    content: string;
}

export interface DividerComponentData {
    style: 'regular';
}

// Backend wraps components in this structure
export interface UserpageComponent {
    header?: HeaderComponentData;
    post?: PostComponentData;
    paragraph?: ParagraphComponentData;
    divider?: DividerComponentData;
}

export interface Userpage {
    id: string;
    user_id: string;
    components: UserpageComponent[];
}

// API Request/Response Types
export interface CreatePostRequest {
    url: string;
}

export interface AddPostRequest {
    post: Post;
}

export interface UserPostsRequest {
    user_id: string;
}

export interface LoginRequest {
    name: string;
    email: string;
}

export interface AddComponentRequest {
    index: number;
    component: UserpageComponent;
}

export interface MoveComponentRequest {
    prev_index: number;
    new_index: number;
}

export interface UpdateComponentRequest {
    index: number;
    component: UserpageComponent;
}

export interface DeleteComponentRequest {
    index: number;
}