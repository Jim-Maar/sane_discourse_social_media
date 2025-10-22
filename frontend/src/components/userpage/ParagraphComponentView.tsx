import { useState, useRef, useEffect } from 'react';
import type { ParagraphComponentData } from '../../types';

interface ParagraphComponentViewProps {
    component: ParagraphComponentData;
    isEditMode: boolean;
    onUpdate?: (component: ParagraphComponentData) => void;
}

export const ParagraphComponentView = ({ 
    component, 
    isEditMode, 
    onUpdate 
}: ParagraphComponentViewProps) => {
    const [isEditing, setIsEditing] = useState(false);
    const textRef = useRef<HTMLParagraphElement>(null);

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
        if (e.key === 'Escape') {
            if (textRef.current) {
                textRef.current.textContent = component.content;
            }
            textRef.current?.blur();
        }
    };

    const style: React.CSSProperties = {
        margin: '0 0 1em 0',
        outline: 'none',
        cursor: isEditMode ? 'text' : 'default',
        textAlign: isEditMode ? 'left' : 'justify',
        hyphens: isEditMode ? 'none' : 'auto',
    };

    if (isEditMode && !isEditing) {
        style.padding = '0.25rem';
        style.borderRadius = '3px';
        style.transition = 'background-color 0.2s';
    }

    return (
        <p
            ref={textRef}
            style={style}
            className={isEditMode ? 'editable-content' : ''}
            contentEditable={isEditMode && isEditing}
            suppressContentEditableWarning
            onClick={handleClick}
            onBlur={handleBlur}
            onKeyDown={handleKeyDown}
        >
            {component.content || (isEditMode ? 'Click to edit paragraph...' : '')}
        </p>
    );
};
