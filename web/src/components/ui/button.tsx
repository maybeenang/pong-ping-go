import { Button as BaseButton } from '@base-ui/react';
import { cva, type VariantProps } from 'class-variance-authority';
import type { ComponentProps, FC } from 'react';
import { cn } from '../../utils/helper';

const buttonVariants = cva(['inline-flex items-center justify-center rounded'], {
    variants: {
        variant: {
            primary: 'bg-primary text-white hover:bg-primary/90',
            secondary: 'bg-secondary text-white hover:bg-secondary/90',
        },
        size: {
            sm: 'px-2 py-1 text-sm',
            md: 'px-4 py-2',
            lg: 'px-6 py-3 text-lg',
        },
    },
    defaultVariants: {
        variant: 'primary',
        size: 'md',
    },
});

type ButtonProps = ComponentProps<typeof BaseButton> & VariantProps<typeof buttonVariants>;

const Button: FC<ButtonProps> = ({ variant, size, className, children, ...props }) => {
    return (
        <BaseButton
            data-slot="button"
            data-size={size}
            {...props}
            className={cn(buttonVariants({ variant, size }), className)}
        >
            {children}
        </BaseButton>
    );
};

export default Button;
