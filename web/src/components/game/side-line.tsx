import type { Graphics } from 'pixi.js';
import { useCallback } from 'react';

type Props = {
    canvasWidth: number;
    canvasHeight: number;
};

const SideLine = ({ canvasWidth, canvasHeight }: Props) => {
    const draw = useCallback(
        (g: Graphics) => {
            g.clear();

            g.rect(0, 0, 2, canvasHeight);
            g.rect(canvasWidth - 2, 0, 2, canvasHeight);
            g.rect(0, 0, canvasWidth, 2);
            g.rect(0, canvasHeight - 2, canvasWidth, 2);

            g.fill({
                color: 0x000000,
                alpha: 0.6,
            });
        },
        [canvasWidth, canvasHeight],
    );

    return <pixiGraphics draw={draw} />;
};

export default SideLine;
