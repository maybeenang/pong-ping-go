export interface GameState {
    ball_x: number;
    ball_y: number;
    paddle_1: number;
    paddle_2: number;
    score_1?: number;
    score_2?: number;
}

export interface PlayerInput {
    direction: 'UP' | 'DOWN';
}

export interface RoomResponse {
    room_id: string;
}
