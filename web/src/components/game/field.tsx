import { Application, extend, useApplication } from '@pixi/react';
import { Graphics } from 'pixi.js';
import { useRef } from 'react';
import { useGameStore } from '../../stores/use-game-store';
import { GAME } from '../../utils/constants';
import { cn } from '../../utils/helper';
import Ball from './ball';
import CourtLine from './court-line';
import Paddle from './paddle';

extend({
    Graphics,
});

const GameContent = () => {
    const { app } = useApplication();

    const gameState = useGameStore((state) => state.gameState);

    const width = app.screen.width;
    const height = app.screen.height;

    if (width === 0 || height === 0) return null;

    const sx = (value: number) => (value / 100) * width;
    const sy = (value: number) => (value / 100) * height;

    const ballRadius = GAME.BALL_RADIUS * Math.min(width / GAME.WIDTH, height / GAME.HEIGHT);

    return (
        <>
            <CourtLine canvasWidth={width} canvasHeight={height} />
            <Ball x={sx(gameState.ball_x)} y={sy(gameState.ball_y)} radius={ballRadius} />
            <Paddle
                x={sx(GAME.LEFT_PADDLE_X)}
                y={sy(gameState.paddle_1)}
                width={GAME.PADDLE_WIDTH}
                height={GAME.PADDLE_HEIGHT}
                player="left"
            />

            <Paddle
                x={sx(GAME.RIGHT_PADDLE_X)}
                y={sy(gameState.paddle_2)}
                width={GAME.PADDLE_WIDTH}
                height={GAME.PADDLE_HEIGHT}
                player="right"
            />
        </>
    );
};

const Field = () => {
    const ref = useRef<HTMLDivElement>(null);

    return (
        <div className={cn('h-full w-full bg-sky-200')} ref={ref}>
            <Application resizeTo={ref} backgroundColor={0x87ceeb}>
                <GameContent />
            </Application>
        </div>
    );
};

export default Field;
