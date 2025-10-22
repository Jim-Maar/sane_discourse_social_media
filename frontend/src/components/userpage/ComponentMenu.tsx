import { useEffect, useRef } from 'react';

// Export the type
type ComponentType = 'header' | 'paragraph' | 'post' | 'divider';

interface ComponentMenuProps {
    position: { x: number; y: number };
    onSelect: (type: ComponentType) => void;
    onClose: () => void;
}

const ComponentMenu = ({ position, onSelect, onClose }: ComponentMenuProps) => {
    const menuRef = useRef<HTMLDivElement>(null);

    useEffect(() => {
        const handleClickOutside = (event: MouseEvent) => {
            if (menuRef.current && !menuRef.current.contains(event.target as Node)) {
                onClose();
            }
        };

        const handleEscape = (event: KeyboardEvent) => {
            if (event.key === 'Escape') {
                onClose();
            }
        };

        document.addEventListener('mousedown', handleClickOutside);
        document.addEventListener('keydown', handleEscape);

        return () => {
            document.removeEventListener('mousedown', handleClickOutside);
            document.removeEventListener('keydown', handleEscape);
        };
    }, [onClose]);

    return (
        <div
            ref={menuRef}
            className="component-menu"
            style={{
                left: `${position.x}px`,
                top: `${position.y}px`,
            }}
        >
            <button
                className="component-menu-item"
                onClick={() => onSelect('header')}
            >
                📝 Header
            </button>
            <button
                className="component-menu-item"
                onClick={() => onSelect('paragraph')}
            >
                📄 Paragraph
            </button>
            <button
                className="component-menu-item"
                onClick={() => onSelect('post')}
            >
                🔗 Post
            </button>
            <button
                className="component-menu-item"
                onClick={() => onSelect('divider')}
            >
                ➖ Divider
            </button>
        </div>
    );
};

// Named exports at the bottom
export { ComponentMenu, type ComponentType };
