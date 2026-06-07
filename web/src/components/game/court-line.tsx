import type { Graphics } from 'pixi.js';
import { useCallback } from 'react';

type CourtLineProps = {
    canvasWidth: number;
    canvasHeight: number;
};

const CourtLine = ({ canvasWidth, canvasHeight }: CourtLineProps) => {
    const draw = useCallback(
        (g: Graphics) => {
            g.clear();

            const centerX = canvasWidth / 2;
            const dashHeight = 10;
            const gapHeight = 8;

            for (let y = 0; y < canvasHeight; y += dashHeight + gapHeight) {
                g.rect(centerX - 1, y, 2, dashHeight);
            }

            g.fill({
                color: 0xffffff,
                alpha: 0.6,
            });
        },
        [canvasWidth, canvasHeight],
    );

    return <pixiGraphics draw={draw} />;
};

export default CourtLine;
