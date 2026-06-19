import { createRoom } from '@/actions/room-action';
import { Button } from '@/components/ui/button';
import {
    Dialog,
    DialogContent,
    DialogDescription,
    DialogFooter,
    DialogHeader,
    DialogTitle,
    DialogTrigger,
} from '@/components/ui/dialog';
import { useActionState, useState } from 'react';
import { Field, FieldGroup } from '../ui/field';
import { Input } from '../ui/input';
import { Label } from '../ui/label';

const CreateRoomDialog = () => {
    const [open, setOpen] = useState(false);

    const [error, submitAction, isLoading] = useActionState(
        async (prev: unknown, formData: FormData) => {
            try {
                const roomName = formData.get('name') as string;
                await createRoom(roomName);
                setOpen(false);
            } catch (error) {
                return {
                    error: 'Failed to create room. Please try again.',
                };
            }
        },
        null,
    );

    return (
        <Dialog
            open={open}
            onOpenChange={(isOpen) => {
                setOpen(isOpen);
            }}
        >
            <DialogTrigger render={<Button>Create New Room</Button>} />
            <DialogContent className="sm:max-w-sm">
                <form action={submitAction} className="space-y-4">
                    <DialogHeader>
                        <DialogTitle>Create New Room</DialogTitle>
                        {error && <DialogDescription>{error.error}</DialogDescription>}
                    </DialogHeader>
                    <FieldGroup>
                        <Field>
                            <Label htmlFor="name">Room Name</Label>
                            <Input id="name" name="name" defaultValue="Pedro Duarte" />
                        </Field>
                    </FieldGroup>
                    <DialogFooter className="flex sm:flex-col" showCloseButton={true}>
                        <Button type="submit" disabled={isLoading}>
                            Create
                        </Button>
                    </DialogFooter>
                </form>
            </DialogContent>
        </Dialog>
    );
};

export default CreateRoomDialog;
