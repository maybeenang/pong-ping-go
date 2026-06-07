import { cn } from '../utils/helper';

const MainMenuButton: React.FC<{
    onClick: () => void;
    className?: string;
    children: React.ReactNode;
}> = ({ onClick, className, children, ...props }) => {
    return (
        <div
            className={cn(
                'flex cursor-pointer items-center justify-center rounded-xl border-2 border-gray-300 p-4 text-lg font-semibold text-gray-500 transition-colors duration-200 hover:bg-gray-100',
                className,
            )}
            onClick={onClick}
            {...props}
        >
            {children}
        </div>
    );
};

export default MainMenuButton;
