import { useState } from 'react';
import { useNavigate } from 'react-router';
import { createRoom } from '../actions/room-action';
import BaseDialog from '../components/base-dialog';
import MainMenuButton from '../components/main-menu-button';
import Button from '../components/ui/button';

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
            label: 'Quick Play',
            className: 'row-span-4',
            onClick: () => {
                navigate('/play');
            },
        },
        {
            label: 'Create Room',
            className: 'row-span-3',
            onClick: () => {
                setOpenCreate(true);
            },
        },
        { label: 'Join Room', className: 'row-span-3', onClick: () => {} },
        { label: 'Room List', className: 'row-span-2', onClick: () => {} },
    ];

    return (
        <div className="h-full p-8">
            <div className="grid h-full grid-cols-2 grid-rows-6 gap-2 p-32">
                {buttons.map((button) => (
                    <MainMenuButton
                        key={button.label}
                        onClick={button.onClick}
                        className={button.className}
                    >
                        {button.label}
                    </MainMenuButton>
                ))}
            </div>
            <BaseDialog size="sm" open={openCreate} onOpenChange={(open) => setOpenCreate(open)}>
                <div className="flex h-full flex-col items-center justify-center gap-4">
                    <h2 className="text-2xl font-bold">Create Room</h2>
                    <form action={handleSubmitCreateRoom}>
                        <input
                            type="text"
                            name="roomName"
                            placeholder="Room Name"
                            className="w-full rounded border border-gray-300 px-4 py-2 focus:ring-2 focus:ring-blue-500 focus:outline-none"
                        />
                        <Button type="submit" className="mt-4 w-full">
                            Create
                        </Button>
                    </form>
                </div>
            </BaseDialog>
        </div>
    );
};

export default HomeScreen;
