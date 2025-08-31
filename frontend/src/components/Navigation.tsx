import { Link } from 'react-router-dom';

interface NavigationProps {
    currentUser: string | null;
    onLogout: () => void;
}

export const Navigation = ({ currentUser, onLogout }: NavigationProps) => {
    return (
        <nav style={{
            padding: '1rem',
            borderBottom: '1px solid #ccc',
            marginBottom: '2rem',
            display: 'flex',
            justifyContent: 'space-between',
            alignItems: 'center'
        }}>
            <div>
                <Link to="/" style={{ marginRight: '1rem' }}>
                    Home
                </Link>
                <Link to="/user">
                    User Page
                </Link>
            </div>
            <div>
                {currentUser ? (
                    <div>
                        <span style={{ marginRight: '1rem' }}>Welcome, {currentUser}!</span>
                        <button onClick={onLogout}>Logout</button>
                    </div>
                ) : (
                    <div>
                        <span style={{ marginRight: '1rem' }}>Not logged in</span>
                        <button
                            onClick={() => window.location.href = 'http://localhost:3000/auth/google'}
                            style={{
                                backgroundColor: '#4285f4',
                                color: 'white',
                                border: 'none',
                                padding: '8px 16px',
                                borderRadius: '4px',
                                cursor: 'pointer'
                            }}
                        >
                            Login with Google
                        </button>
                    </div>
                )}
            </div>
        </nav>
    );
};