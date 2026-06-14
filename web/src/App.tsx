import { createBrowserRouter, RouterProvider } from 'react-router';
import MainLayout from './components/main-layout';
import { getRoom, getRooms } from './loaders/room-loader';
import GameScreen from './screens/game-screen';
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
                path: '/play',
                Component: PlayScreen,
                loader: async () => {
                    try {
                        const res = await getRooms();
                        return res;
                    } catch (error) {
                        throw new Response('Failed to load rooms', { status: 500 });
                    }
                },
            },
            {
                path: '/game/:roomId',
                Component: GameScreen,
                loader: async ({ params }) => {
                    const { roomId } = params;
                    if (!roomId) {
                        throw new Response('Room ID not found', { status: 404 });
                    }

                    try {
                        const res = await getRoom(roomId);
                        return res;
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
