export interface SuccessResponse<T> {
    status: true;
    message: string;
    data: T;
}

export interface ErrorResponse {
    status: false;
    message: string;
    error: string;
}

export type ApiResponse<T> = SuccessResponse<T> | ErrorResponse;

export type Room = {
    id: string;
    name: string;
    status: 'waiting' | 'playing' | 'finished';
    created_at: string;
};

export type GetRoomResponse = ApiResponse<{
    room: Room;
}>;
export type ListRoomsResponse = ApiResponse<{
    rooms: Room[];
}>;
