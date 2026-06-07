import { Dialog as BaseDialog } from '@base-ui/react';
import { cva, type VariantProps } from 'class-variance-authority';
import type { ComponentProps, FC } from 'react';
import { cn } from '../../utils/helper';

export const Dialog: FC<ComponentProps<typeof BaseDialog.Root>> = ({ ...props }) => {
    return <BaseDialog.Root data-slot="dialog" {...props} />;
};

export const DialogTrigger: FC<ComponentProps<typeof BaseDialog.Trigger>> = ({ ...props }) => {
    return <BaseDialog.Trigger data-slot="trigger" {...props} />;
};

const dialogContentVariants = cva(
    [
        'transition-all',
        'fixed top-1/2 left-1/2 w-full translate-x-[-50%] translate-y-[-50%] rounded-lg bg-white p-6 shadow-lg',
        'focus:outline-none',
        'data-[ending-style]:opacity-0 data-[starting-style]:opacity-0',
    ],
    {
        variants: {
            size: {
                sm: 'w-[90%] max-w-sm',
                md: 'w-[90%] max-w-md',
                lg: 'w-[90%] max-w-lg',
            },
        },
        defaultVariants: {
            size: 'md',
        },
    },
);

export type DialogContentProps = ComponentProps<typeof BaseDialog.Popup> &
    VariantProps<typeof dialogContentVariants>;

export const DialogContent: FC<DialogContentProps> = ({ size, children, className, ...props }) => {
    return (
        <BaseDialog.Portal>
            <BaseDialog.Backdrop
                data-slot="backdrop"
                className={cn(
                    'fixed inset-0 min-h-dvh bg-black/60 backdrop-blur-sm transition-[color,opacity]',
                    'data-[ending-style]:opacity-0 data-[starting-style]:opacity-0',
                )}
            />
            <BaseDialog.Popup
                data-slot="popup"
                {...props}
                className={cn(dialogContentVariants({ size }), className)}
            >
                {children}
            </BaseDialog.Popup>
        </BaseDialog.Portal>
    );
};

export const DialogHeader: FC<ComponentProps<'header'>> = ({ children, className, ...props }) => {
    return (
        <header
            data-slot="header"
            {...props}
            className={cn('mb-4 text-xl font-semibold', className)}
        >
            {children}
        </header>
    );
};

export const DialogTitle: FC<ComponentProps<typeof BaseDialog.Title>> = ({
    children,
    className,
    ...props
}) => {
    return (
        <BaseDialog.Title
            data-slot="title"
            {...props}
            className={cn('text-lg font-medium', className)}
        >
            {children}
        </BaseDialog.Title>
    );
};

export const DialogDescription: FC<ComponentProps<typeof BaseDialog.Description>> = ({
    children,
    className,
    ...props
}) => {
    return (
        <BaseDialog.Description
            data-slot="description"
            {...props}
            className={cn('text-gray-600', className)}
        >
            {children}
        </BaseDialog.Description>
    );
};

export const DialogBody: FC<ComponentProps<'div'>> = ({ children, className, ...props }) => {
    return (
        <div data-slot="body" {...props} className={cn('mb-4 text-gray-700', className)}>
            {children}
        </div>
    );
};

export const DialogFooter: FC<ComponentProps<'footer'>> = ({ children, className, ...props }) => {
    return (
        <footer
            data-slot="footer"
            className={cn('mt-4 flex justify-end space-x-2', className)}
            {...props}
        >
            {children}
        </footer>
    );
};

export const DialogClose: FC<ComponentProps<typeof BaseDialog.Close>> = ({
    children,
    className,
    ...props
}) => {
    return (
        <BaseDialog.Close
            data-slot="close"
            {...props}
            className={cn(
                'rounded-md bg-gray-200 px-4 py-2 text-sm font-medium text-gray-700 hover:bg-gray-300',
                className,
            )}
        >
            {children || 'Close'}
        </BaseDialog.Close>
    );
};
