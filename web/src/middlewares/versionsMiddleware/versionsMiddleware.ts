import { type StoreApi } from 'zustand';

// services
import VersionsAPIService from '@/services/VersionsAPIService';

// types
import type { Store } from '@/types/store';
import type { Logger } from '@/utilities/logging/createLogger';

export default function versionsMiddleware(logger: Logger, { setState }: StoreApi<Store>): () => Promise<void> {
  return async () => {
    const apiService = new VersionsAPIService({
      logger,
    });

    try {
      const { environment, version } = await apiService.versions();

      setState((state) => ({
        ...state,
        environment,
        version,
      }));
    } catch (error) {
      return;
    }
  };
}
