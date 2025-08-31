import { useState } from 'react';
import { useMutation, useQuery, useQueryClient } from '@tanstack/react-query';
import { userLogin, getUserPosts, createPostFromUrl, addPost } from '../api';
import { PostCard } from '../components/PostCard';
import { EditablePostCard } from '../components/EditablePostCard';
import type { User, Post } from '../types';

interface UserpageProps {
  currentUser: User | null;
  setCurrentUser: (user: User | null) => void;
}

export const Userpage = ({ currentUser, setCurrentUser }: UserpageProps) => {
  const [username, setUsername] = useState('');
  const [urlText, setUrlText] = useState('');
  const [pendingPosts, setPendingPosts] = useState<Post[]>([]);

  const queryClient = useQueryClient();

  // Login mutation
  const loginMutation = useMutation({
    mutationFn: userLogin,
    onSuccess: (user) => {
      setCurrentUser(user);
      localStorage.setItem('currentUser', JSON.stringify(user));
    },
    onError: (error) => {
      console.error('Login failed:', error);
      alert('Login failed. Please try again.');
    }
  });

  // Create post from URL mutation
  const createPostMutation = useMutation({
    mutationFn: createPostFromUrl,
    onSuccess: (post) => {
      setPendingPosts(prev => [...prev, post]);
      setUrlText('');
    },
    onError: (error) => {
      console.error('Failed to create post:', error);
      alert('Failed to create post. Please check the URL and try again.');
    }
  });

  // Add post mutation
  const addPostMutation = useMutation({
    mutationFn: addPost,
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['userPosts'] });
      queryClient.invalidateQueries({ queryKey: ['feed'] });
    },
    onError: (error) => {
      console.error('Failed to add post:', error);
      alert('Failed to add post. Please try again.');
    }
  });

  // User posts query
  const { data: userPosts, isLoading: userPostsLoading } = useQuery({
    queryKey: ['userPosts', currentUser?.id],
    queryFn: () => currentUser ? getUserPosts({ user_id: currentUser.id }) : Promise.resolve([]),
    enabled: !!currentUser,
  });

  const handleLogin = (e: React.FormEvent) => {
    e.preventDefault();
    if (username.trim()) {
      loginMutation.mutate({ username: username.trim() });
    }
  };

  const handleCreatePost = (e: React.FormEvent) => {
    e.preventDefault();
    if (!urlText.trim()) return;

    createPostMutation.mutate({ url: urlText.trim() });
  };

  const handleAcceptPost = (post: Post) => {
    if (!currentUser) return;

    addPostMutation.mutate({
      user_id: currentUser.id,
      post
    });

    // Remove from pending posts
    setPendingPosts(prev => prev.filter(p => p.url !== post.url));
  };

  const handleRejectPost = (post: Post) => {
    setPendingPosts(prev => prev.filter(p => p.url !== post.url));
  };

  if (!currentUser) {
    return (
      <div style={{ maxWidth: '400px', margin: '2rem auto', padding: '1rem' }}>
        <h1>User Login</h1>
        <p>Please enter your username to continue:</p>

        <form onSubmit={handleLogin} style={{ marginTop: '1rem' }}>
          <div style={{ marginBottom: '1rem' }}>
            <label style={{ display: 'block', marginBottom: '0.5rem' }}>
              Username:
            </label>
            <input
              type="text"
              value={username}
              onChange={(e) => setUsername(e.target.value)}
              required
              style={{
                width: '100%',
                padding: '0.5rem',
                border: '1px solid #ddd',
                borderRadius: '4px'
              }}
            />
          </div>
          <button
            type="submit"
            disabled={loginMutation.isPending}
            style={{
              width: '100%',
              padding: '0.75rem',
              backgroundColor: '#007bff',
              color: 'white',
              border: 'none',
              borderRadius: '4px',
              cursor: 'pointer'
            }}
          >
            {loginMutation.isPending ? 'Logging in...' : 'Login'}
          </button>
        </form>
      </div>
    );
  }

  return (
    <div style={{ maxWidth: '800px', margin: '0 auto', padding: '1rem' }}>
      <h1>User Page</h1>

      {/* URL Input Section */}
      <div style={{
        border: '1px solid #ddd',
        borderRadius: '8px',
        padding: '1rem',
        marginBottom: '2rem',
        backgroundColor: 'var(--card-bg, #f9f9f9)'
      }}>
        <h2>Add New Post</h2>
        <form onSubmit={handleCreatePost}>
          <div style={{ marginBottom: '1rem' }}>
            <label style={{ display: 'block', marginBottom: '0.5rem' }}>
              Paste URL:
            </label>
            <input
              type="url"
              value={urlText}
              onChange={(e) => setUrlText(e.target.value)}
              placeholder="https://example.com/article"
              style={{
                width: '100%',
                padding: '0.5rem',
                border: '1px solid #ddd',
                borderRadius: '4px'
              }}
            />
          </div>
          <button
            type="submit"
            disabled={createPostMutation.isPending || !urlText.trim()}
            style={{
              padding: '0.75rem 1.5rem',
              backgroundColor: '#28a745',
              color: 'white',
              border: 'none',
              borderRadius: '4px',
              cursor: 'pointer'
            }}
          >
            {createPostMutation.isPending ? 'Processing URL...' : 'Create Post'}
          </button>
        </form>
      </div>

      {/* Pending Posts Section */}
      {pendingPosts.length > 0 && (
        <div style={{ marginBottom: '2rem' }}>
          <h2>Review & Accept Posts</h2>
          <p style={{ color: '#666', marginBottom: '1rem' }}>
            Review the post below. You can edit empty fields and the description.
            Accept the post to add it to your collection.
          </p>
          {pendingPosts.map((post, index) => (
            <EditablePostCard
              key={post.url || index}
              post={post}
              onAccept={handleAcceptPost}
              onReject={() => handleRejectPost(post)}
            />
          ))}
        </div>
      )}

      {/* User Posts Section */}
      <div>
        <h2>Your Posts</h2>
        {userPostsLoading ? (
          <div>Loading your posts...</div>
        ) : userPosts && userPosts.length > 0 ? (
          <div>
            {userPosts.map((post) => (
              <PostCard key={post.id || post.url} post={post} />
            ))}
          </div>
        ) : (
          <div style={{ textAlign: 'center', padding: '2rem', color: '#666' }}>
            You haven't added any posts yet
          </div>
        )}
      </div>
    </div>
  );
};
