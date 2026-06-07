import { createBrowserRouter, RouterProvider } from 'react-router';
import { getRoom } from './actions/room-action';
import MainLayout from './components/main-layout';
import HomeScreen from './screens/home-screen';
import PlayScreen from './screens/play-screen';

const router = createBrowserRouter([
    {
        Component: MainLayout,
        children: [
            {
                path: '/',
                Component: HomeScreen,
            },
            {
                path: '/play/:roomId',
                Component: PlayScreen,
                loader: async ({ params }) => {
                    const { roomId } = params;
                    if (!roomId) {
                        throw new Response('Room ID not found', { status: 404 });
                    }

                    try {
                        await getRoom(roomId);
                        return { roomId };
                    } catch (error) {
                        throw new Response('Room not found', { status: 404 });
                    }
                },
            },
        ],
    },
]);

const App = () => {
    return <RouterProvider router={router} />;
};

export default App;
