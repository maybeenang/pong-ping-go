import CreateRoomDialog from '@/components/dialogs/create-room-dialog';
import { Button } from '@/components/ui/button';
import { Card, CardAction, CardContent, CardHeader, CardTitle } from '@/components/ui/card';
import { InputGroup, InputGroupAddon, InputGroupInput } from '@/components/ui/input-group';
import { SearchIcon } from 'lucide-react';
import { useNavigate } from 'react-router';

const CardRoom = () => {
    const navigate = useNavigate();
    return (
        <Card size="default">
            <CardHeader>
                <CardTitle>Room Name</CardTitle>
                <CardAction>Code: abcde</CardAction>
            </CardHeader>
            <CardContent className="flex flex-row items-center justify-between">
                <p>Players: 1/2</p>
                <Button
                    className=""
                    size={'sm'}
                    onClick={() => {
                        navigate('/game/abcde');
                    }}
                >
                    Join
                </Button>
            </CardContent>
        </Card>
    );
};

const PlayScreen = () => {
    return (
        <section className="grid grid-cols-4 gap-4 p-4">
            <div className="col-span-1 space-y-4">
                <Card size="sm">
                    <CardHeader>
                        <CardTitle className="font-bold">Filter</CardTitle>
                    </CardHeader>
                    <CardContent>
                        <p>Card Content</p>
                    </CardContent>
                </Card>

                <Card size="sm">
                    <CardHeader>
                        <CardTitle className="font-bold">Sort By</CardTitle>
                    </CardHeader>
                    <CardContent>
                        <p>Card Content</p>
                    </CardContent>
                </Card>
            </div>
            <div className="col-span-3 grid grid-cols-3 gap-4">
                <div className="col-span-3 flex flex-row items-center justify-between gap-4">
                    <div className="flex w-full flex-row items-center gap-4">
                        <Button variant="outline" className="">
                            Main Menu
                        </Button>
                        <InputGroup className="max-w-xs">
                            <InputGroupInput placeholder="Search..." />
                            <InputGroupAddon align={'inline-end'}>
                                <SearchIcon />
                            </InputGroupAddon>
                        </InputGroup>
                    </div>
                    <CreateRoomDialog />
                </div>
                {Array.from({ length: 10 }).map((_, index) => (
                    <CardRoom key={index} />
                ))}
            </div>
        </section>
    );
};

export default PlayScreen;
