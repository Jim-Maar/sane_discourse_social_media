import { Link } from 'react-router-dom';

interface NavigationProps {
    currentUser: string | null;
    onLogout: () => void;
}

export const Navigation = ({ currentUser, onLogout }: NavigationProps) => {
    return (
        <nav style={{
            display: 'flex',
            justifyContent: 'space-between',
            alignItems: 'center',
            flexWrap: 'wrap',
            gap: '1rem'
        }}>
            <div style={{ display: 'flex', gap: '1.5rem', alignItems: 'center' }}>
                <Link to="/">HOME</Link>
                <Link to="/user">USER</Link>
            </div>
            <div style={{ 
                display: 'flex', 
                alignItems: 'center', 
                gap: '1rem',
                fontSize: '0.9em',
                fontFamily: '-apple-system, BlinkMacSystemFont, "Segoe UI", "Roboto", sans-serif'
            }}>
                {currentUser ? (
                    <>
                        <span style={{ color: 'var(--text-dim)' }}>
                            {currentUser}
                        </span>
                        <button 
                            onClick={onLogout}
                            style={{
                                padding: '0.4rem 0.8rem',
                                fontSize: '0.85em'
                            }}
                        >
                            Logout
                        </button>
                    </>
                ) : (
                    <button
                        onClick={() => window.location.href = 'http://localhost:3000/auth/google'}
                        style={{
                            padding: '0.4rem 0.8rem',
                            fontSize: '0.85em',
                            backgroundColor: '#4285f4',
                            color: 'white',
                            borderColor: '#4285f4'
                        }}
                    >
                        Sign in with Google
                    </button>
                )}
            </div>
        </nav>
    );
};
