export const GAME = {
    WIDTH: 800,
    HEIGHT: 600,
    PADDLE_WIDTH: 10,
    PADDLE_HEIGHT: 100,
    BALL_RADIUS: 10,
    LEFT_PADDLE_X: 0,
    RIGHT_PADDLE_X: 97,
} as const;

export const getWebSocketUrl = (roomId: string) => {
    const protocol = window.location.protocol === 'https:' ? 'wss' : 'ws';
    const host = window.location.host;
    return `${protocol}://${host}/ws?room=${roomId}`;
};
