import axios from 'axios';
import type { RoomResponse } from '../types/game';

export const createRoom = async (roomName: string) => {
    try {
        const response = await axios.post<RoomResponse>('/api/create-room', { name: roomName });
        if (response.status === 200) {
            return response.data.room_id;
        } else {
            throw new Error('Failed to create room');
        }
    } catch (error) {
        console.error('Error creating room:', error);
        throw error;
    }
};
