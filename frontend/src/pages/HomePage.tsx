import { useQuery } from '@tanstack/react-query';
import { getFeed } from '../api';
import { PostCard } from '../components/PostCard';

export const HomePage = () => {
    const { data: posts, isLoading, error } = useQuery({
        queryKey: ['feed'],
        queryFn: getFeed,
    });

    if (isLoading) {
        return <div style={{ textAlign: 'center', padding: '2rem' }}>Loading feed...</div>;
    }

    if (error) {
        return (
            <div style={{ textAlign: 'center', padding: '2rem', color: 'red' }}>
                Error loading feed: {error instanceof Error ? error.message : 'Unknown error'}
            </div>
        );
    }

    return (
        <div style={{ maxWidth: '800px', margin: '0 auto', padding: '1rem' }}>
            <h1>Home Feed</h1>
            {posts && posts.length > 0 ? (
                <div>
                    {posts.map((post) => (
                        <PostCard key={post.id || post.url} post={post} />
                    ))}
                </div>
            ) : (
                <div style={{ textAlign: 'center', padding: '2rem', color: '#666' }}>
                    No posts available
                </div>
            )}
        </div>
    );
};