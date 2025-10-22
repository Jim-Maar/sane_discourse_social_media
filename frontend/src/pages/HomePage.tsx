import { useQuery } from '@tanstack/react-query';
import { getFeed } from '../api';
import { PostCard } from '../components/PostCard';

export const HomePage = () => {
    const { data: posts, isLoading, error } = useQuery({
        queryKey: ['feed'],
        queryFn: getFeed,
    });

    if (isLoading) {
        return (
            <div className="container">
                <div className="loading">Loading feed...</div>
            </div>
        );
    }

    if (error) {
        return (
            <div className="container">
                <div className="error">
                    Error loading feed: {error instanceof Error ? error.message : 'Unknown error'}
                </div>
            </div>
        );
    }

    return (
        <div className="container">
            <h1>Home Feed</h1>
            {posts && posts.length > 0 ? (
                <div>
                    {posts.map((post) => (
                        <PostCard key={post.id || post.url} post={post} />
                    ))}
                </div>
            ) : (
                <div style={{ 
                    textAlign: 'center', 
                    padding: '3rem 1rem', 
                    color: 'var(--text-dim)' 
                }}>
                    No posts available yet.
                </div>
            )}
        </div>
    );
};
