import { useState, useRef, useEffect } from 'react';
import type { HeaderComponentData } from '../../types';

interface HeaderComponentViewProps {
    component: HeaderComponentData;
    isEditMode: boolean;
    onUpdate?: (component: HeaderComponentData) => void;
}

export const HeaderComponentView = ({ 
    component, 
    isEditMode, 
    onUpdate 
}: HeaderComponentViewProps) => {
    const [isEditing, setIsEditing] = useState(false);
    const textRef = useRef<HTMLHeadingElement>(null);

    const handleClick = () => {
        if (isEditMode && !isEditing) {
            setIsEditing(true);
        }
    };

    const handleBlur = () => {
        setIsEditing(false);
        const newContent = textRef.current?.textContent || '';
        if (onUpdate && newContent !== component.content) {
            onUpdate({ ...component, content: newContent });
        }
    };

    const handleKeyDown = (e: React.KeyboardEvent) => {
        if (e.key === 'Enter' && !e.shiftKey) {
            e.preventDefault();
            textRef.current?.blur();
        }
        if (e.key === 'Escape') {
            if (textRef.current) {
                textRef.current.textContent = component.content;
            }
            textRef.current?.blur();
        }
    };

    const sizeMap = {
        1: { tag: 'h1', fontSize: '2.2em' },
        2: { tag: 'h2', fontSize: '1.6em' },
        3: { tag: 'h3', fontSize: '1.3em' },
        4: { tag: 'h4', fontSize: '1.1em' },
    };

    const { tag: Tag, fontSize } = sizeMap[component.size];

    const style: React.CSSProperties = {
        fontSize,
        margin: component.size === 1 ? '1em 0 0.5em 0' : '1.2em 0 0.5em 0',
        outline: 'none',
        cursor: isEditMode ? 'text' : 'default',
    };

    if (isEditMode && !isEditing) {
        style.padding = '0.25rem';
        style.borderRadius = '3px';
        style.transition = 'background-color 0.2s';
    }

    return (
        <Tag
            ref={textRef as any}
            style={style}
            className={isEditMode ? 'editable-content' : ''}
            contentEditable={isEditMode && isEditing}
            suppressContentEditableWarning
            onClick={handleClick}
            onBlur={handleBlur}
            onKeyDown={handleKeyDown}
        >
            {component.content || (isEditMode ? 'Click to edit header...' : '')}
        </Tag>
    );
};
