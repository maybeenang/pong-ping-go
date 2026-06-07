import { useCallback, useEffect } from 'react';
import useWebSocketModule, { ReadyState } from 'react-use-websocket';
import type { GameState } from '../types/game';
import { getWebSocketUrl } from '../utils/constants';
import { useGameStore, type ConnectionStatus } from './use-game-store';

const { default: useWebSocket = useWebSocketModule } = useWebSocketModule as unknown as {
    default: typeof useWebSocketModule;
};

const STATUS_MAP: Record<ReadyState, ConnectionStatus> = {
    [ReadyState.CONNECTING]: 'connecting',
    [ReadyState.OPEN]: 'connected',
    [ReadyState.CLOSING]: 'closing',
    [ReadyState.CLOSED]: 'disconnected',
    [ReadyState.UNINSTANTIATED]: 'uninstantiated',
};

export const useGameSocket = (roomId: string) => {
    const setGameState = useGameStore((state) => state.setGameState);
    const setConnectionStatus = useGameStore((state) => state.setConnectionStatus);

    const { sendJsonMessage, readyState } = useWebSocket(getWebSocketUrl(roomId), {
        onMessage: (event: MessageEvent) => {
            try {
                const data: GameState = JSON.parse(event.data);
                setGameState(data);
            } catch (error) {
                console.error('Failed to parse game state:', error);
            }
        },
        filter: () => false,
        shouldReconnect: () => true,
        reconnectAttempts: 10,
        reconnectInterval: (attemptNumber) => {
            const delay = Math.min(1000 * 2 ** attemptNumber, 30000);
            return delay;
        },
        onOpen: () => {
            console.log('WebSocket connection opened');
        },
        onClose: () => console.log('WebSocket connection closed'),
        onError: (event) => {
            console.error('WebSocket error:', event);
        },
    });

    useEffect(() => {
        setConnectionStatus(STATUS_MAP[readyState]);
    }, [readyState, setConnectionStatus]);

    const sendInput = useCallback(
        (direction: 'UP' | 'DOWN') => {
            if (readyState === ReadyState.OPEN) {
                sendJsonMessage({ direction });
            }
        },
        [sendJsonMessage, readyState],
    );

    return { sendInput, isConnected: readyState === ReadyState.OPEN };
};
