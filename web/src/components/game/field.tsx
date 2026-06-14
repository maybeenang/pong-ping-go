import { GAME } from '@/lib/constants';
import { cn } from '@/lib/utils';
import { Application, extend, useApplication } from '@pixi/react';
import { Graphics } from 'pixi.js';
import { useRef } from 'react';
import { useGameStore } from '../../stores/use-game-store';
import Ball from './ball';
import CourtLine from './court-line';
import Paddle from './paddle';
import SideLine from './side-line';

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
            <SideLine canvasWidth={width} canvasHeight={height} />
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
        <div className="min-h-0 w-full flex-1 p-4">
            <div className={cn('h-full w-full')} ref={ref}>
                <Application resizeTo={ref} backgroundColor={'white'}>
                    <GameContent />
                </Application>
            </div>
        </div>
    );
};

export default Field;
