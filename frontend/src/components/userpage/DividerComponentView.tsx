import type { DividerComponentData } from '../../types';

interface DividerComponentViewProps {
    component: DividerComponentData;
}

export const DividerComponentView = ({ component }: DividerComponentViewProps) => {
    return <hr style={{ margin: '2em 0' }} />;
};
