import type { RoomResponse } from '@/types/game';
import axios from 'axios';

export const getRoom = async (roomId: string) => {
    try {
        const response = await axios.get<RoomResponse>(`/api/rooms/${roomId}`);
        if (response.status === 200) {
            return response.data;
        } else {
            throw new Error('Failed to get room');
        }
    } catch (error) {
        console.error('Error getting room:', error);
        throw error;
    }
};

export const getRooms = async () => {
    try {
        const response = await axios.get<RoomResponse[]>('/api/rooms');
        if (response.status === 200) {
            return response.data;
        } else {
            throw new Error('Failed to get rooms');
        }
    } catch (error) {
        console.error('Error getting rooms:', error);
        throw error;
    }
};
