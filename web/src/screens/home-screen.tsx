import { Button } from '@/components/ui/button';
import { useState } from 'react';
import { useNavigate } from 'react-router';
import { createRoom } from '../actions/room-action';

const HomeScreen = () => {
    const navigate = useNavigate();

    const [openCreate, setOpenCreate] = useState(false);

    const handleSubmitCreateRoom = async (formdata: FormData) => {
        const roomName = formdata.get('roomName') as string;
        if (!roomName) {
            alert('Please enter a room name');
            return;
        }
        try {
            const roomId = await createRoom(roomName);
            navigate(`/play/${roomId}`);
        } catch (error) {
            alert('Failed to create room. Please try again.');
        }
    };

    const buttons = [
        {
            label: 'Play',
            onClick: () => {
                navigate('/play');
            },
            variant: 'default' as 'outline' | 'default',
        },
        {
            label: 'Profile',
        },
        {
            label: 'Settings',
        },
    ];

    return (
        <div className="h-full w-full flex-col items-center justify-center gap-8 p-4 md:flex">
            <h1 className="text-center text-8xl font-bold">PING PONG</h1>
            <div className="mx-auto flex w-full max-w-sm flex-col gap-4">
                {buttons.map((button) => (
                    <Button
                        key={button.label}
                        onClick={button.onClick}
                        size={'xl'}
                        variant={button.variant || 'outline'}
                    >
                        {button.label}
                    </Button>
                ))}
            </div>
        </div>
    );
};

export default HomeScreen;
