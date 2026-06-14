import type { Graphics } from 'pixi.js';
import { useCallback } from 'react';

interface PaddleProps {
    x: number;
    y: number;
    width: number;
    height: number;
    player: 'left' | 'right';
}

const Paddle = ({ x = 0, y = 0, player = 'left', width, height }: PaddleProps) => {
    const draw = useCallback(
        (g: Graphics) => {
            g.clear();

            g.rect(0, -height / 2, width, height);

            g.fill({
                color: 'black',
            });
        },
        [width, height, player],
    );

    return <pixiGraphics draw={draw} x={x} y={y} />;
};

export default Paddle;
