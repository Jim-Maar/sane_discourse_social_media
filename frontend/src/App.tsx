import { useState, useEffect } from 'react';
import { BrowserRouter as Router, Routes, Route } from 'react-router-dom';
import { QueryClient, QueryClientProvider } from '@tanstack/react-query';
import { Navigation } from './components/Navigation';
import { HomePage } from './pages/HomePage';
import { Userpage } from './pages/UserPage';
import type { User } from './types';
import './App.css';

const queryClient = new QueryClient({
  defaultOptions: {
    queries: {
      retry: 1,
      refetchOnWindowFocus: false,
    },
  },
});

function App() {
  const [currentUser, setCurrentUser] = useState<User | null>(null);
  const [isCheckingAuth, setIsCheckingAuth] = useState(true);

  // Check if user is authenticated on app start
  useEffect(() => {
    const checkAuth = async () => {
      try {
        const response = await fetch('http://localhost:3000/auth/me', {
          method: 'PUT',
          credentials: 'include',
          headers: {
            'Content-Type': 'application/json',
          },
        });
        
        if (response.ok) {
          const user = await response.json();
          setCurrentUser(user);
        }
      } catch (error) {
        console.log('Not authenticated');
      } finally {
        setIsCheckingAuth(false);
      }
    };

    checkAuth();
  }, []);

  const handleLogout = () => {
    setCurrentUser(null);
    // Optionally redirect to logout endpoint
    window.location.href = 'http://localhost:3000/auth/logout';
  };

  if (isCheckingAuth) {
    return (
      <div style={{ 
        display: 'flex', 
        alignItems: 'center', 
        justifyContent: 'center', 
        minHeight: '100vh',
        color: 'var(--text-color)'
      }}>
        Loading...
      </div>
    );
  }

  return (
    <QueryClientProvider client={queryClient}>
      <Router>
        <div style={{ minHeight: '100vh' }}>
          <Navigation
            currentUser={currentUser?.username || null}
            onLogout={handleLogout}
          />

          <Routes>
            <Route path="/" element={<HomePage />} />
            <Route path="/home" element={<HomePage />} />
            <Route
              path="/user"
              element={
                <Userpage
                  currentUser={currentUser}
                  setCurrentUser={setCurrentUser}
                />
              }
            />
          </Routes>
        </div>
      </Router>
    </QueryClientProvider>
  );
}

export default App;
