export interface User {
    id: string;
    username: string;
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

export interface CreatePostRequest {
    url: string;
}

export interface AddPostRequest {
    user_id: string;
    post: Post;
}

export interface UserPostsRequest {
    user_id: string;
}

export interface LoginRequest {
    username: string;
}