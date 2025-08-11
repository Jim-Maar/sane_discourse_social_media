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
                    <span>Not logged in</span>
                )}
            </div>
        </nav>
    );
};