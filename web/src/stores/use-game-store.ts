import { create } from 'zustand';
import type { GameState } from '../types/game';

export type ConnectionStatus =
    | 'connecting'
    | 'connected'
    | 'disconnected'
    | 'closing'
    | 'uninstantiated';

interface GameStore {
    roomId: string | null;
    setRoomId: (roomId: string) => void;

    gameState: GameState;
    setGameState: (state: GameState) => void;

    conectionStatus: ConnectionStatus;
    setConnectionStatus: (status: ConnectionStatus) => void;
}

export const useGameStore = create<GameStore>((set) => ({
    roomId: null,
    setRoomId: (roomId: string) => set({ roomId }),

    gameState: {
        ball_x: 50,
        ball_y: 50,
        paddle_1: 50,
        paddle_2: 50,
    },
    setGameState: (state: GameState) => set({ gameState: state }),

    conectionStatus: 'uninstantiated',
    setConnectionStatus: (status: ConnectionStatus) => set({ conectionStatus: status }),
}));
