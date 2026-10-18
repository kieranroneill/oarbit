'use client';

import { useCallback, useState } from 'react';

// services
import ChatAPIService from '@/services/ChatAPIService';

// types
import type { UseChatSessionsState } from './types';
import type { ChatSessionSummary } from '@/types/chat';

export default function useChatSessions(): UseChatSessionsState {
  // states
  const [sessions, setSessions] = useState<ChatSessionSummary[]>([]);
  const [loading, setLoading] = useState(true);
  // callbacks
  const fetch = useCallback(async () => {
    const apiService = new ChatAPIService();

    setLoading(true);
    setSessions(await apiService.getSessions());
    setLoading(false);
  }, [setLoading, setSessions]);

  return {
    fetch,
    loading,
    sessions,
  };
}
