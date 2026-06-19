import axios from 'axios';
import type { RoomResponse } from '../types/game';

export const createRoom = async (roomName: string) => {
    try {
        const response = await axios.post<RoomResponse>('/api/rooms', { name: roomName });
        return response.data;
    } catch (error) {
        console.error('Error creating room:', error);
        throw error;
    }
};
