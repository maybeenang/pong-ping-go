import { Outlet } from 'react-router';
import { cn } from '../utils/helper';

const MainLayout = () => {
    return (
        <main className={cn('w-screen', 'h-screen', 'bg-zinc-50')}>
            <Outlet />
        </main>
    );
};

export default MainLayout;
