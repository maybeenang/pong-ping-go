import { useEffect, useRef } from 'react';

const INPUT_KEYS = new Set(['ArrowUp', 'ArrowDown', 'w', 's']);
const INPUT_INTERFAL = 50; // ms

type SendinputFn = (direction: 'UP' | 'DOWN') => void;

export const useKeyboardInput = (sendInput: SendinputFn, enabled: boolean) => {
    const pressedkeys = useRef<Set<string>>(new Set());

    const handleKeyDown = (event: KeyboardEvent) => {
        if (!INPUT_KEYS.has(event.key)) return;
        event.preventDefault();
        pressedkeys.current.add(event.key);
    };

    const handleKeyUp = (event: KeyboardEvent) => {
        pressedkeys.current.delete(event.key);
    };

    const handleBlur = () => {
        pressedkeys.current.clear();
    };

    useEffect(() => {
        if (!enabled) return;

        window.addEventListener('keydown', handleKeyDown);
        window.addEventListener('keyup', handleKeyUp);
        window.addEventListener('blur', handleBlur);

        const interval = setInterval(() => {
            if (pressedkeys.current.has('ArrowUp') || pressedkeys.current.has('w')) {
                sendInput('UP');
            }
            if (pressedkeys.current.has('ArrowDown') || pressedkeys.current.has('s')) {
                sendInput('DOWN');
            }
        }, INPUT_INTERFAL);

        return () => {
            window.removeEventListener('keydown', handleKeyDown);
            window.removeEventListener('keyup', handleKeyUp);
            window.removeEventListener('blur', handleBlur);
            clearInterval(interval);
            pressedkeys.current.clear();
        };
    }, [sendInput, enabled]);
};
