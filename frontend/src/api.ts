import axios from 'axios';
import type {
    User,
    Post,
    CreatePostRequest,
    AddPostRequest,
    UserPostsRequest,
    LoginRequest,
    Userpage,
    AddComponentRequest,
    MoveComponentRequest,
    UpdateComponentRequest,
    DeleteComponentRequest
} from './types';

const API_BASE_URL = 'http://localhost:3000';

const api = axios.create({
    baseURL: API_BASE_URL,
    withCredentials: true, // Important for session-based auth
});

// Auth endpoints
export const mockLogin = async (request: LoginRequest): Promise<User> => {
    const response = await api.put('/auth/login', request);
    return response.data;
};

export const getCurrentUser = async (): Promise<User> => {
    const response = await api.put('/auth/me');
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

export const getUserPosts = async (): Promise<Post[]> => {
    const response = await api.get('/user/posts');
    return response.data;
};

export const getFeed = async (): Promise<Post[]> => {
    const response = await api.get('/home');
    return response.data;
};

// Userpage endpoints
export const getUserpage = async (): Promise<Userpage> => {
    const response = await api.get('/userpage');
    return response.data;
};

export const addUserpageComponent = async (request: AddComponentRequest): Promise<Userpage> => {
    const response = await api.put('/userpage/component/add', request);
    return response.data;
};

export const updateUserpageComponent = async (request: UpdateComponentRequest): Promise<Userpage> => {
    const response = await api.put('/userpage/component/update', request);
    return response.data;
};

export const deleteUserpageComponent = async (request: DeleteComponentRequest): Promise<Userpage> => {
    const response = await api.delete('/userpage/component/delete', { data: request });
    return response.data;
};

export const moveUserpageComponent = async (request: MoveComponentRequest): Promise<Userpage> => {
    const response = await api.put('/userpage/component/move', request);
    return response.data;
};