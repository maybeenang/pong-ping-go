import { cn } from '@/lib/utils';
import { Outlet } from 'react-router';

const MainLayout = () => {
    return (
        <main className={cn('w-screen', 'h-screen', 'bg-background', 'p-4')}>
            <div className="h-full w-full border border-primary">
                <Outlet />
            </div>
        </main>
    );
};

export default MainLayout;
