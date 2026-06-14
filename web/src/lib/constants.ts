export const GAME = {
    WIDTH: 700,
    HEIGHT: 600,
    BALL_RADIUS: 10,
    PADDLE_WIDTH: 10,
    PADDLE_HEIGHT: 100,
    LEFT_PADDLE_X: 2,
    RIGHT_PADDLE_X: 97,
};

export const getWebsocketUrl = (roomId: string) => {
    const protocol = window.location.protocol === 'https:' ? 'wss' : 'ws';
    const host = window.location.host;
    return `${protocol}://${host}/ws?room=${roomId}`;
};
