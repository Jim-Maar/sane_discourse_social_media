import { useState } from 'react';
import type { Post } from '../types';

interface EditablePostCardProps {
    post: Post;
    onAccept: (editedPost: Post) => void;
    onReject: () => void;
}

export const EditablePostCard = ({ post, onAccept, onReject }: EditablePostCardProps) => {
    const [editedPost, setEditedPost] = useState<Post>(post);

    const handleFieldChange = (field: keyof Post, value: string) => {
        setEditedPost(prev => ({
            ...prev,
            [field]: value
        }));
    };

    const isFieldEmpty = (value: string) => !value || value.trim() === '';

    return (
        <div style={{
            border: '2px solid #007bff',
            borderRadius: '8px',
            padding: '1rem',
            marginBottom: '1rem',
            backgroundColor: 'var(--card-bg, #f9f9f9)'
        }}>
            <div style={{ display: 'flex', gap: '1rem', marginBottom: '1rem' }}>
                {editedPost.thumbnail_url && (
                    <img
                        src={editedPost.thumbnail_url}
                        alt={editedPost.title}
                        style={{
                            width: '120px',
                            height: '80px',
                            objectFit: 'cover',
                            borderRadius: '4px'
                        }}
                    />
                )}
                <div style={{ flex: 1 }}>
                    <div style={{ marginBottom: '0.5rem' }}>
                        <label style={{ display: 'block', fontWeight: 'bold', marginBottom: '0.25rem' }}>
                            Title (readonly):
                        </label>
                        <input
                            type="text"
                            value={editedPost.title}
                            disabled
                            style={{
                                width: '100%',
                                padding: '0.5rem',
                                border: '1px solid #ddd',
                                borderRadius: '4px',
                                backgroundColor: '#f5f5f5',
                                color: '#666'
                            }}
                        />
                    </div>

                    <div style={{ marginBottom: '0.5rem' }}>
                        <label style={{ display: 'block', fontWeight: 'bold', marginBottom: '0.25rem' }}>
                            Description:
                        </label>
                        <textarea
                            value={editedPost.description}
                            onChange={(e) => handleFieldChange('description', e.target.value)}
                            placeholder={isFieldEmpty(post.description) ? "Add description..." : ""}
                            style={{
                                width: '100%',
                                padding: '0.5rem',
                                border: '1px solid #ddd',
                                borderRadius: '4px',
                                minHeight: '60px',
                                backgroundColor: isFieldEmpty(post.description) ? '#fff7e6' : 'white',
                                color: '#333'
                            }}
                        />
                    </div>

                    <div style={{ display: 'grid', gridTemplateColumns: '1fr 1fr', gap: '0.5rem' }}>
                        <div>
                            <label style={{ display: 'block', fontWeight: 'bold', marginBottom: '0.25rem' }}>
                                Site Name:
                            </label>
                            <input
                                type="text"
                                value={editedPost.site_name}
                                onChange={(e) => handleFieldChange('site_name', e.target.value)}
                                placeholder={isFieldEmpty(post.site_name) ? "Add site name..." : ""}
                                style={{
                                    width: '100%',
                                    padding: '0.5rem',
                                    border: '1px solid #ddd',
                                    borderRadius: '4px',
                                    backgroundColor: isFieldEmpty(post.site_name) ? '#fff7e6' : 'white',
                                    color: '#333'
                                }}
                            />
                        </div>

                        <div>
                            <label style={{ display: 'block', fontWeight: 'bold', marginBottom: '0.25rem' }}>
                                Author:
                            </label>
                            <input
                                type="text"
                                value={editedPost.author}
                                onChange={(e) => handleFieldChange('author', e.target.value)}
                                placeholder={isFieldEmpty(post.author) ? "Add author..." : ""}
                                style={{
                                    width: '100%',
                                    padding: '0.5rem',
                                    border: '1px solid #ddd',
                                    borderRadius: '4px',
                                    backgroundColor: isFieldEmpty(post.author) ? '#fff7e6' : 'white',
                                    color: '#333'
                                }}
                            />
                        </div>

                        <div>
                            <label style={{ display: 'block', fontWeight: 'bold', marginBottom: '0.25rem' }}>
                                Type:
                            </label>
                            <input
                                type="text"
                                value={editedPost.type}
                                onChange={(e) => handleFieldChange('type', e.target.value)}
                                placeholder={isFieldEmpty(post.type) ? "Add type..." : ""}
                                style={{
                                    width: '100%',
                                    padding: '0.5rem',
                                    border: '1px solid #ddd',
                                    borderRadius: '4px',
                                    backgroundColor: isFieldEmpty(post.type) ? '#fff7e6' : 'white',
                                    color: '#333'
                                }}
                            />
                        </div>

                        <div>
                            <label style={{ display: 'block', fontWeight: 'bold', marginBottom: '0.25rem' }}>
                                Thumbnail URL:
                            </label>
                            <input
                                type="url"
                                value={editedPost.thumbnail_url}
                                onChange={(e) => handleFieldChange('thumbnail_url', e.target.value)}
                                placeholder={isFieldEmpty(post.thumbnail_url) ? "Add thumbnail URL..." : ""}
                                style={{
                                    width: '100%',
                                    padding: '0.5rem',
                                    border: '1px solid #ddd',
                                    borderRadius: '4px',
                                    backgroundColor: isFieldEmpty(post.thumbnail_url) ? '#fff7e6' : 'white',
                                    color: '#333'
                                }}
                            />
                        </div>
                    </div>

                    <div style={{ marginTop: '0.5rem' }}>
                        <label style={{ display: 'block', fontWeight: 'bold', marginBottom: '0.25rem' }}>
                            URL (readonly):
                        </label>
                        <input
                            type="url"
                            value={editedPost.url}
                            disabled
                            style={{
                                width: '100%',
                                padding: '0.5rem',
                                border: '1px solid #ddd',
                                borderRadius: '4px',
                                backgroundColor: '#f5f5f5',
                                color: '#666'
                            }}
                        />
                    </div>
                </div>
            </div>

            <div style={{ display: 'flex', gap: '0.5rem', justifyContent: 'flex-end' }}>
                <button
                    onClick={onReject}
                    style={{
                        padding: '0.5rem 1rem',
                        backgroundColor: '#dc3545',
                        color: 'white',
                        border: 'none',
                        borderRadius: '4px',
                        cursor: 'pointer'
                    }}
                >
                    Reject
                </button>
                <button
                    onClick={() => onAccept(editedPost)}
                    style={{
                        padding: '0.5rem 1rem',
                        backgroundColor: '#28a745',
                        color: 'white',
                        border: 'none',
                        borderRadius: '4px',
                        cursor: 'pointer'
                    }}
                >
                    Accept & Add to Database
                </button>
            </div>
        </div>
    );
};