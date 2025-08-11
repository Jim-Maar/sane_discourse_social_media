import type { Post } from '../types';

interface PostCardProps {
    post: Post;
}

export const PostCard = ({ post }: PostCardProps) => {
    return (
        <div style={{
            border: '1px solid #ddd',
            borderRadius: '8px',
            padding: '1rem',
            marginBottom: '1rem',
            backgroundColor: 'var(--card-bg, #f9f9f9)'
        }}>
            <div style={{ display: 'flex', gap: '1rem' }}>
                {post.thumbnail_url && (
                    <img
                        src={post.thumbnail_url}
                        alt={post.title}
                        style={{
                            width: '120px',
                            height: '80px',
                            objectFit: 'cover',
                            borderRadius: '4px'
                        }}
                    />
                )}
                <div style={{ flex: 1 }}>
                    <h3 style={{ margin: '0 0 0.5rem 0' }}>
                        <a
                            href={post.url}
                            target="_blank"
                            rel="noopener noreferrer"
                            style={{ textDecoration: 'none', color: 'inherit' }}
                        >
                            {post.title}
                        </a>
                    </h3>
                    {post.description && (
                        <p style={{ margin: '0 0 0.5rem 0', color: '#666' }}>
                            {post.description}
                        </p>
                    )}
                    <div style={{ fontSize: '0.9em', color: '#888' }}>
                        {post.site_name && <span>{post.site_name}</span>}
                        {post.author && <span> • by {post.author}</span>}
                        {post.type && <span> • {post.type}</span>}
                    </div>
                </div>
            </div>
        </div>
    );
};