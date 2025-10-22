import type { Post } from '../types';

interface PostCardProps {
    post: Post;
    size?: 2 | 3; // 2=medium (default), 3=small
}

export const PostCard = ({ post, size = 2 }: PostCardProps) => {
    if (size === 3) {
        // Small size: title only
        return (
            <div className="post-card" style={{ padding: '0.75rem' }}>
                <h3 style={{ 
                    margin: 0, 
                    fontSize: '1em',
                    fontWeight: 'normal'
                }}>
                    <a
                        href={post.url}
                        target="_blank"
                        rel="noopener noreferrer"
                    >
                        {post.title}
                    </a>
                </h3>
            </div>
        );
    }

    // Medium size: thumbnail, title, description
    return (
        <div className="post-card">
            <div style={{ display: 'flex', gap: '1rem' }}>
                {post.thumbnail_url && (
                    <img
                        src={post.thumbnail_url}
                        alt={post.title}
                        style={{
                            width: '140px',
                            height: '100px',
                            objectFit: 'cover',
                            borderRadius: '3px',
                            flexShrink: 0
                        }}
                    />
                )}
                <div style={{ flex: 1, minWidth: 0 }}>
                    <h3 style={{ 
                        margin: '0 0 0.5rem 0',
                        fontSize: '1.2em'
                    }}>
                        <a
                            href={post.url}
                            target="_blank"
                            rel="noopener noreferrer"
                        >
                            {post.title}
                        </a>
                    </h3>
                    {post.description && (
                        <p style={{ 
                            margin: '0 0 0.5rem 0', 
                            color: 'var(--text-dim)',
                            fontSize: '0.95em'
                        }}>
                            {post.description}
                        </p>
                    )}
                    <div style={{ 
                        fontSize: '0.85em', 
                        color: 'var(--text-dim)',
                        fontFamily: '-apple-system, BlinkMacSystemFont, "Segoe UI", "Roboto", sans-serif'
                    }}>
                        {post.site_name && <span>{post.site_name}</span>}
                        {post.author && <span> • {post.author}</span>}
                        {post.type && <span> • {post.type}</span>}
                    </div>
                </div>
            </div>
        </div>
    );
};
