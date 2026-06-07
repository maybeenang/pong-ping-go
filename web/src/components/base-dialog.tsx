import type { ComponentProps, FC, ReactNode } from 'react';
import { Dialog, DialogBody, DialogContent, DialogTrigger } from './ui/dialog';

export type BaseDialogProps = ComponentProps<typeof Dialog> &
    Omit<ComponentProps<typeof DialogContent>, keyof ComponentProps<typeof Dialog>> & {
        children?: ReactNode;
        className?: string;
        renderTrigger?: ReactNode;
    };

const BaseDialog: FC<BaseDialogProps> = ({
    children,
    open,
    defaultOpen,
    onOpenChange,
    className,
    size,
    renderTrigger,
    ...props
}) => {
    return (
        <Dialog open={open} defaultOpen={defaultOpen} onOpenChange={onOpenChange}>
            {renderTrigger ? <DialogTrigger>{renderTrigger}</DialogTrigger> : null}
            <DialogContent {...props} className={className} size={size}>
                <DialogBody>{children as ReactNode}</DialogBody>
            </DialogContent>
        </Dialog>
    );
};

export default BaseDialog;
