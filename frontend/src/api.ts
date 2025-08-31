import axios from 'axios';
import type {
    User,
    Post,
    CreatePostRequest,
    AddPostRequest,
    UserPostsRequest,
    LoginRequest
} from './types';

const API_BASE_URL = 'http://localhost:3000';

const api = axios.create({
    baseURL: API_BASE_URL,
});

// User endpoints
export const userLogin = async (request: LoginRequest): Promise<User> => {
    const response = await api.put('/user/login', request);
    return response.data;
};

// Post endpoints
export const createPostFromUrl = async (request: CreatePostRequest): Promise<Post> => {
    const response = await api.put('/user/posts/create', request);
    return response.data;
};

export const addPost = async (request: AddPostRequest): Promise<Post> => {
    const response = await api.put('/user/posts/add', request);
    return response.data;
};

export const getUserPosts = async (request: UserPostsRequest): Promise<Post[]> => {
    const response = await api.post('/user/posts', request);
    return response.data;
};

export const getFeed = async (): Promise<Post[]> => {
    const response = await api.get('/home');
    return response.data;
};