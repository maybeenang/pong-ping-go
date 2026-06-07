import { useEffect } from 'react';
import { useParams } from 'react-router';
import Field from '../components/game/field';
import { useGameStore } from '../stores/use-game-store';
import { useGameSocket } from '../stores/use-gamesocket';
import { useKeyboardInput } from '../stores/use-keyboard-input';
import { cn } from '../utils/helper';

const PlayScreen = () => {
    const { roomId } = useParams();
    const setRoomId = useGameStore((state) => state.setRoomId);
    const connectionStatus = useGameStore((state) => state.conectionStatus);

    useEffect(() => {
        if (roomId) {
            setRoomId(roomId);
        }
    }, [roomId, setRoomId]);

    const { sendInput, isConnected } = useGameSocket(roomId || '');
    useKeyboardInput(sendInput, isConnected);

    return (
        <div className="flex h-full flex-col items-center gap-4">
            <h1 className="text-4xl font-bold">Room: {roomId}</h1>
            <span>
                Connection status:{' '}
                <span
                    className={cn('font-semibold', isConnected ? 'text-green-500' : 'text-red-500')}
                >
                    {connectionStatus}
                </span>
            </span>
            <Field />
        </div>
    );
};

export default PlayScreen;
