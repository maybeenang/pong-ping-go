import type { Graphics } from 'pixi.js';
import { useCallback } from 'react';

type BallProps = {
    x: number;
    y: number;
    radius: number;
};

const Ball = ({ x, y, radius }: BallProps) => {
    const draw = useCallback(
        (g: Graphics) => {
            g.clear();
            g.rect(-radius, -radius, radius * 2, radius * 2);
            g.fill('black');
        },
        [radius],
    );

    return <pixiGraphics draw={draw} x={x} y={y} />;
};

export default Ball;
