import { useEffect } from 'react';
import { useLoaderData, useNavigate } from 'react-router';
import Field from '../components/game/field';
import { useGameStore } from '../stores/use-game-store';
import { useGameSocket } from '../stores/use-gamesocket';
import { useKeyboardInput } from '../stores/use-keyboard-input';

const PlayerInfo = ({ player }: { player: 'player 1' | 'player 2' }) => {
    return (
        <div className="flex flex-col items-center gap-2">
            <p className="">{player}</p>
            <p className="text-6xl font-bold">04</p>
        </div>
    );
};

const GameScreen = () => {
    const { room_id: roomId } = useLoaderData<{
        room_id: string;
    }>();

    const navigate = useNavigate();
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
        <div className="flex h-full flex-col items-center overflow-hidden">
            <div className="flex w-full items-end justify-between gap-4 px-16 py-8">
                <PlayerInfo player="player 1" />
                <div>
                    <p className="text-center text-2xl font-bold">VS</p>
                    <span>{connectionStatus}</span>
                </div>
                <PlayerInfo player="player 1" />
            </div>
            <Field />
        </div>
    );
};

export default GameScreen;
