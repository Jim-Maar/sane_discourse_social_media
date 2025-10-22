import { useState } from 'react';
import { useQuery, useMutation, useQueryClient } from '@tanstack/react-query';
import { getUserPosts, createPostFromUrl, addPost } from '../../api';
import { PostCard } from '../PostCard';
import type { PostComponentData, Post } from '../../types';

interface PostComponentViewProps {
    component: PostComponentData;
    isEditMode: boolean;
    onUpdate?: (component: PostComponentData) => void;
}

export const PostComponentView = ({ 
    component, 
    isEditMode, 
    onUpdate 
}: PostComponentViewProps) => {
    const [selectedSize, setSelectedSize] = useState(component.size);
    const queryClient = useQueryClient();

    // Fetch user posts to find the post by ID
    const { data: userPosts } = useQuery({
        queryKey: ['userPosts'],
        queryFn: getUserPosts,
    });

    const post = userPosts?.find(p => p.id === component.post_id);

    const handleSizeChange = (newSize: 2 | 3) => {
        setSelectedSize(newSize);
        if (onUpdate) {
            onUpdate({ ...component, size: newSize });
        }
    };

    if (!post) {
        return (
            <div style={{ 
                padding: '1rem', 
                border: '1px dashed var(--border-color)',
                borderRadius: '4px',
                color: 'var(--text-dim)',
                textAlign: 'center'
            }}>
                Post not found (ID: {component.post_id})
            </div>
        );
    }

    return (
        <div>
            <PostCard post={post} size={selectedSize} />
            {isEditMode && (
                <div className="size-selector">
                    <span style={{ 
                        fontSize: '0.85em', 
                        color: 'var(--text-dim)',
                        marginRight: '0.5rem'
                    }}>
                        Size:
                    </span>
                    <button
                        className={`size-button ${selectedSize === 2 ? 'active' : ''}`}
                        onClick={() => handleSizeChange(2)}
                    >
                        Medium
                    </button>
                    <button
                        className={`size-button ${selectedSize === 3 ? 'active' : ''}`}
                        onClick={() => handleSizeChange(3)}
                    >
                        Small
                    </button>
                </div>
            )}
        </div>
    );
};

// Component for creating a new post in edit mode
interface PostCreatorProps {
    onPostCreated: (postId: string) => void;
    onCancel: () => void;
}

export const PostCreator = ({ onPostCreated, onCancel }: PostCreatorProps) => {
    const [url, setUrl] = useState('');
    const [pendingPost, setPendingPost] = useState<Post | null>(null);
    const queryClient = useQueryClient();

    const createMutation = useMutation({
        mutationFn: createPostFromUrl,
        onSuccess: (post) => {
            setPendingPost(post);
        },
        onError: (error) => {
            console.error('Failed to create post:', error);
            alert('Failed to create post. Please check the URL and try again.');
        }
    });

    const addMutation = useMutation({
        mutationFn: addPost,
        onSuccess: (addedPost) => {
            queryClient.invalidateQueries({ queryKey: ['userPosts'] });
            queryClient.invalidateQueries({ queryKey: ['feed'] });
            if (addedPost.id) {
                onPostCreated(addedPost.id);
            }
        },
        onError: (error) => {
            console.error('Failed to add post:', error);
            alert('Failed to add post. Please try again.');
        }
    });

    const handleCreatePost = (e: React.FormEvent) => {
        e.preventDefault();
        if (url.trim()) {
            createMutation.mutate({ url: url.trim() });
        }
    };

    const handleAddPost = () => {
        if (pendingPost) {
            addMutation.mutate({ post: pendingPost });
        }
    };

    const handleEditField = (field: keyof Post, value: string) => {
        if (pendingPost) {
            setPendingPost({ ...pendingPost, [field]: value });
        }
    };

    if (pendingPost) {
        return (
            <div style={{
                border: '2px solid var(--link-color)',
                borderRadius: '4px',
                padding: '1rem',
                backgroundColor: 'var(--card-bg)',
                marginBottom: '1rem'
            }}>
                <h4 style={{ marginTop: 0 }}>Review Post</h4>
                
                <div style={{ marginBottom: '1rem' }}>
                    <PostCard post={pendingPost} size={2} />
                </div>

                <div style={{ display: 'grid', gap: '0.75rem', marginBottom: '1rem' }}>
                    <div>
                        <label className="form-label">Title</label>
                        <input
                            className="form-input"
                            type="text"
                            value={pendingPost.title}
                            disabled
                        />
                    </div>
                    
                    <div>
                        <label className="form-label">Description</label>
                        <textarea
                            className="form-input"
                            value={pendingPost.description}
                            onChange={(e) => handleEditField('description', e.target.value)}
                            rows={3}
                        />
                    </div>

                    <div style={{ display: 'grid', gridTemplateColumns: '1fr 1fr', gap: '0.75rem' }}>
                        <div>
                            <label className="form-label">Site Name</label>
                            <input
                                className="form-input"
                                type="text"
                                value={pendingPost.site_name}
                                onChange={(e) => handleEditField('site_name', e.target.value)}
                            />
                        </div>
                        
                        <div>
                            <label className="form-label">Author</label>
                            <input
                                className="form-input"
                                type="text"
                                value={pendingPost.author}
                                onChange={(e) => handleEditField('author', e.target.value)}
                            />
                        </div>
                    </div>

                    <div>
                        <label className="form-label">Type</label>
                        <input
                            className="form-input"
                            type="text"
                            value={pendingPost.type}
                            onChange={(e) => handleEditField('type', e.target.value)}
                        />
                    </div>

                    <div>
                        <label className="form-label">Thumbnail URL</label>
                        <input
                            className="form-input"
                            type="url"
                            value={pendingPost.thumbnail_url}
                            onChange={(e) => handleEditField('thumbnail_url', e.target.value)}
                        />
                    </div>
                </div>

                <div style={{ display: 'flex', gap: '0.5rem', justifyContent: 'flex-end' }}>
                    <button onClick={onCancel}>
                        Cancel
                    </button>
                    <button 
                        onClick={handleAddPost}
                        disabled={addMutation.isPending}
                        style={{
                            backgroundColor: 'var(--link-color)',
                            color: 'var(--bg-color)',
                            borderColor: 'var(--link-color)'
                        }}
                    >
                        {addMutation.isPending ? 'Adding...' : 'Add Post'}
                    </button>
                </div>
            </div>
        );
    }

    return (
        <div style={{
            border: '1px solid var(--border-color)',
            borderRadius: '4px',
            padding: '1rem',
            backgroundColor: 'var(--card-bg)',
            marginBottom: '1rem'
        }}>
            <h4 style={{ marginTop: 0 }}>Create Post from URL</h4>
            <form onSubmit={handleCreatePost}>
                <div className="form-group">
                    <label className="form-label">URL</label>
                    <input
                        className="form-input"
                        type="url"
                        value={url}
                        onChange={(e) => setUrl(e.target.value)}
                        placeholder="https://example.com/article"
                        required
                    />
                </div>
                <div style={{ display: 'flex', gap: '0.5rem', justifyContent: 'flex-end' }}>
                    <button type="button" onClick={onCancel}>
                        Cancel
                    </button>
                    <button 
                        type="submit"
                        disabled={createMutation.isPending || !url.trim()}
                        style={{
                            backgroundColor: 'var(--link-color)',
                            color: 'var(--bg-color)',
                            borderColor: 'var(--link-color)'
                        }}
                    >
                        {createMutation.isPending ? 'Processing...' : 'Fetch Post'}
                    </button>
                </div>
            </form>
        </div>
    );
};
