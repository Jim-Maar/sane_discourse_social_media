import { useState } from 'react';
import { useQuery, useMutation, useQueryClient } from '@tanstack/react-query';
import { getUserpage, addUserpageComponent, updateUserpageComponent, deleteUserpageComponent, moveUserpageComponent } from '../api';
import { HeaderComponentView } from '../components/userpage/HeaderComponentView';
import { ParagraphComponentView } from '../components/userpage/ParagraphComponentView';
import { PostComponentView, PostCreator } from '../components/userpage/PostComponentView';
import { DividerComponentView } from '../components/userpage/DividerComponentView';
import { ComponentMenu, type ComponentType } from '../components/userpage/ComponentMenu';
import type { User, UserpageComponent } from '../types';

interface UserPageProps {
    currentUser: User | null;
    setCurrentUser: (user: User | null) => void;
}

export const Userpage = ({ currentUser, setCurrentUser }: UserPageProps) => {
    const [isEditMode, setIsEditMode] = useState(false);
    const [menuState, setMenuState] = useState<{
        isOpen: boolean;
        position: { x: number; y: number };
        insertIndex: number;
    } | null>(null);
    const [creatingPostAtIndex, setCreatingPostAtIndex] = useState<number | null>(null);
    const [draggedIndex, setDraggedIndex] = useState<number | null>(null);
    const [dragOverIndex, setDragOverIndex] = useState<number | null>(null);

    const queryClient = useQueryClient();

    // Userpage query
    const { data: userpage, isLoading: userpageLoading } = useQuery({
        queryKey: ['userpage'],
        queryFn: getUserpage,
        enabled: !!currentUser,
    });

    // Add component mutation
    const addComponentMutation = useMutation({
        mutationFn: addUserpageComponent,
        onSuccess: (updatedUserpage) => {
            // Update the cache with the returned userpage
            queryClient.setQueryData(['userpage'], updatedUserpage);
        },
        onError: (error) => {
            console.error('Failed to add component:', error);
            alert('Failed to add component. Please try again.');
        }
    });

    // Update component mutation
    const updateComponentMutation = useMutation({
        mutationFn: updateUserpageComponent,
        onSuccess: (updatedUserpage) => {
            queryClient.setQueryData(['userpage'], updatedUserpage);
        },
        onError: (error) => {
            console.error('Failed to update component:', error);
            alert('Failed to update component. Please try again.');
        }
    });

    // Delete component mutation
    const deleteComponentMutation = useMutation({
        mutationFn: deleteUserpageComponent,
        onSuccess: (updatedUserpage) => {
            queryClient.setQueryData(['userpage'], updatedUserpage);
        },
        onError: (error) => {
            console.error('Failed to delete component:', error);
            alert('Failed to delete component. Please try again.');
        }
    });

    // Move component mutation
    const moveComponentMutation = useMutation({
        mutationFn: moveUserpageComponent,
        onSuccess: (updatedUserpage) => {
            // Update the cache with the returned userpage
            queryClient.setQueryData(['userpage'], updatedUserpage);
        },
        onError: (error) => {
            console.error('Failed to move component:', error);
            alert('Failed to move component. Please try again.');
        }
    });

    const handleGoogleLogin = () => {
        window.location.href = 'http://localhost:3000/auth/google';
    };

    const handlePlusClick = (e: React.MouseEvent, index: number) => {
        e.stopPropagation();
        const rect = (e.target as HTMLElement).getBoundingClientRect();
        setMenuState({
            isOpen: true,
            position: { x: rect.right + 5, y: rect.top },
            insertIndex: index,
        });
    };

    const handleComponentSelect = (type: ComponentType) => {
        if (!menuState) return;

        if (type === 'post') {
            setCreatingPostAtIndex(menuState.insertIndex);
            setMenuState(null);
            return;
        }

        let component: UserpageComponent;
        
        switch (type) {
            case 'header':
                component = { header: { content: 'New Header', size: 2 } };
                break;
            case 'paragraph':
                component = { paragraph: { content: 'New paragraph text.' } };
                break;
            case 'divider':
                component = { divider: { style: 'regular' } };
                break;
            default:
                return;
        }

        addComponentMutation.mutate({
            index: menuState.insertIndex,
            component,
        });

        setMenuState(null);
    };

    const handlePostCreated = (postId: string) => {
        if (creatingPostAtIndex === null) return;

        const component: UserpageComponent = {
            post: {
                post_id: postId,
                size: 2,
            }
        };

        addComponentMutation.mutate({
            index: creatingPostAtIndex,
            component,
        });

        setCreatingPostAtIndex(null);
    };

    const handleUpdateComponent = (index: number, component: UserpageComponent) => {
        updateComponentMutation.mutate({
            index,
            component,
        });
    };

    const handleDeleteComponent = (index: number) => {
        if (confirm('Are you sure you want to delete this component?')) {
            deleteComponentMutation.mutate({ index });
        }
    };

    // Drag and drop handlers
    const handleDragStart = (e: React.DragEvent, index: number) => {
        setDraggedIndex(index);
        e.dataTransfer.effectAllowed = 'move';
    };

    const handleDragOver = (e: React.DragEvent, index: number) => {
        e.preventDefault();
        e.dataTransfer.dropEffect = 'move';
        setDragOverIndex(index);
    };

    const handleDragLeave = () => {
        setDragOverIndex(null);
    };

    const handleDrop = (e: React.DragEvent, dropIndex: number) => {
        e.preventDefault();
        
        if (draggedIndex === null || draggedIndex === dropIndex) {
            setDraggedIndex(null);
            setDragOverIndex(null);
            return;
        }

        moveComponentMutation.mutate({
            prev_index: draggedIndex,
            new_index: dropIndex,
        });

        setDraggedIndex(null);
        setDragOverIndex(null);
    };

    const handleDragEnd = () => {
        setDraggedIndex(null);
        setDragOverIndex(null);
    };

    // Render component based on type
    const renderComponent = (component: UserpageComponent, index: number) => {
        if (component.post) {
            return (
                <PostComponentView
                    component={component.post}
                    isEditMode={isEditMode}
                    onUpdate={(updated) => handleUpdateComponent(index, { post: updated })}
                />
            );
        } else if (component.header) {
            return (
                <HeaderComponentView
                    component={component.header}
                    isEditMode={isEditMode}
                    onUpdate={(updated) => handleUpdateComponent(index, { header: updated })}
                />
            );
        } else if (component.paragraph) {
            return (
                <ParagraphComponentView
                    component={component.paragraph}
                    isEditMode={isEditMode}
                    onUpdate={(updated) => handleUpdateComponent(index, { paragraph: updated })}
                />
            );
        } else if (component.divider) {
            return <DividerComponentView component={component.divider} />;
        }
        return null;
    };

    if (!currentUser) {
        return (
            <div className="container">
                <div style={{ 
                    maxWidth: '400px', 
                    margin: '4rem auto',
                    textAlign: 'center'
                }}>
                    <h1>Welcome</h1>
                    <p style={{ 
                        textAlign: 'center', 
                        hyphens: 'none',
                        marginBottom: '2rem',
                        color: 'var(--text-dim)'
                    }}>
                        Please sign in to create and manage your userpage.
                    </p>

                    <button
                        onClick={handleGoogleLogin}
                        style={{
                            width: '100%',
                            padding: '0.75rem 1.5rem',
                            backgroundColor: '#4285f4',
                            color: 'white',
                            border: 'none',
                            borderRadius: '4px',
                            fontSize: '1em',
                            cursor: 'pointer',
                            display: 'flex',
                            alignItems: 'center',
                            justifyContent: 'center',
                            gap: '0.75rem',
                            transition: 'background-color 0.2s'
                        }}
                        onMouseOver={(e) => e.currentTarget.style.backgroundColor = '#357ae8'}
                        onMouseOut={(e) => e.currentTarget.style.backgroundColor = '#4285f4'}
                    >
                        <svg width="18" height="18" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 48 48">
                            <path fill="#EA4335" d="M24 9.5c3.54 0 6.71 1.22 9.21 3.6l6.85-6.85C35.9 2.38 30.47 0 24 0 14.62 0 6.51 5.38 2.56 13.22l7.98 6.19C12.43 13.72 17.74 9.5 24 9.5z"/>
                            <path fill="#4285F4" d="M46.98 24.55c0-1.57-.15-3.09-.38-4.55H24v9.02h12.94c-.58 2.96-2.26 5.48-4.78 7.18l7.73 6c4.51-4.18 7.09-10.36 7.09-17.65z"/>
                            <path fill="#FBBC05" d="M10.53 28.59c-.48-1.45-.76-2.99-.76-4.59s.27-3.14.76-4.59l-7.98-6.19C.92 16.46 0 20.12 0 24c0 3.88.92 7.54 2.56 10.78l7.97-6.19z"/>
                            <path fill="#34A853" d="M24 48c6.48 0 11.93-2.13 15.89-5.81l-7.73-6c-2.15 1.45-4.92 2.3-8.16 2.3-6.26 0-11.57-4.22-13.47-9.91l-7.98 6.19C6.51 42.62 14.62 48 24 48z"/>
                            <path fill="none" d="M0 0h48v48H0z"/>
                        </svg>
                        Sign in with Google
                    </button>
                </div>
            </div>
        );
    }

    if (userpageLoading) {
        return (
            <div className="container">
                <div className="loading">Loading your page...</div>
            </div>
        );
    }

    return (
        <div className={`container ${isEditMode ? 'edit-mode' : ''}`}>
            <div style={{ 
                display: 'flex', 
                justifyContent: 'flex-end',
                alignItems: 'center',
                marginBottom: '1.5rem',
                paddingTop: '0.5rem'
            }}>
                <button
                    onClick={() => setIsEditMode(!isEditMode)}
                    style={{
                        backgroundColor: isEditMode ? 'var(--link-color)' : 'var(--button-bg)',
                        color: isEditMode ? 'var(--bg-color)' : 'var(--text-color)',
                        borderColor: isEditMode ? 'var(--link-color)' : 'var(--border-color)',
                    }}
                >
                    {isEditMode ? '✓ Done Editing' : '✎ Edit Page'}
                </button>
            </div>

            {/* Show post creator if active */}
            {creatingPostAtIndex !== null && (
                <PostCreator
                    onPostCreated={handlePostCreated}
                    onCancel={() => setCreatingPostAtIndex(null)}
                />
            )}

            {/* Render components */}
            {userpage?.components && userpage.components.length > 0 ? (
                <div>
                    {/* Plus button before first component */}
                    {isEditMode && (
                        <div style={{ 
                            position: 'relative', 
                            height: '1.5rem',
                            marginBottom: '0.5rem'
                        }}>
                            <button
                                className="plus-button top"
                                style={{ 
                                    position: 'static',
                                    opacity: 1,
                                    marginLeft: '0'
                                }}
                                onClick={(e) => handlePlusClick(e, 0)}
                            >
                                +
                            </button>
                        </div>
                    )}

                    {userpage.components.map((component, index) => (
                        <div
                            key={index}
                            className={`component-wrapper ${draggedIndex === index ? 'dragging' : ''} ${dragOverIndex === index ? 'drag-over' : ''}`}
                            draggable={isEditMode}
                            onDragStart={(e) => handleDragStart(e, index)}
                            onDragOver={(e) => handleDragOver(e, index)}
                            onDragLeave={handleDragLeave}
                            onDrop={(e) => handleDrop(e, index)}
                            onDragEnd={handleDragEnd}
                        >
                            {isEditMode && (
                                <>
                                    <span className="drag-handle" style={{ 
                                        position: 'absolute',
                                        left: '0.5rem',
                                        top: '0.5rem',
                                        fontSize: '1.2em'
                                    }}>
                                        ⋮⋮
                                    </span>
                                    <button
                                        className="plus-button bottom"
                                        onClick={(e) => handlePlusClick(e, index + 1)}
                                    >
                                        +
                                    </button>
                                    <button
                                        onClick={() => handleDeleteComponent(index)}
                                        style={{
                                            position: 'absolute',
                                            right: '0.5rem',
                                            top: '0.5rem',
                                            padding: '0.3rem 0.6rem',
                                            fontSize: '0.85em',
                                            backgroundColor: '#d44',
                                            color: 'white',
                                            border: '1px solid #c33',
                                            borderRadius: '3px',
                                            cursor: 'pointer',
                                            opacity: 0,
                                            transition: 'opacity 0.2s'
                                        }}
                                        className="delete-button"
                                    >
                                        🗑️
                                    </button>
                                </>
                            )}
                            {renderComponent(component, index)}
                        </div>
                    ))}
                </div>
            ) : (
                <div style={{ 
                    textAlign: 'center', 
                    padding: '3rem 1rem',
                    color: 'var(--text-dim)'
                }}>
                    {isEditMode ? (
                        <div>
                            <p>Click the + button to add your first component.</p>
                            <button
                                className="plus-button"
                                style={{ 
                                    position: 'static',
                                    opacity: 1,
                                    marginTop: '1rem'
                                }}
                                onClick={(e) => handlePlusClick(e, 0)}
                            >
                                +
                            </button>
                        </div>
                    ) : (
                        <p>Click "Edit Page" to start building your page.</p>
                    )}
                </div>
            )}

            {/* Component menu */}
            {menuState?.isOpen && (
                <ComponentMenu
                    position={menuState.position}
                    onSelect={handleComponentSelect}
                    onClose={() => setMenuState(null)}
                />
            )}
        </div>
    );
};
