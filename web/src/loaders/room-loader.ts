import type { GetRoomResponse, ListRoomsResponse } from '@/types/api';
import axios from 'axios';

export const getRoom = async (roomId: string) => {
    try {
        const response = await axios.get<GetRoomResponse>(`/api/rooms/${roomId}`);
        if (response.data.status) {
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
        const response = await axios.get<ListRoomsResponse>('/api/rooms');
        if (response.data.status) {
            return response.data;
        } else {
            throw new Error('Failed to get rooms');
        }
    } catch (error) {
        console.error('Error getting rooms:', error);
        throw error;
    }
};
