import { type StoreApi } from 'zustand';

// types
import type { Store } from '@/types/store';
import type { Logger } from '@/utilities/logging/createLogger';

interface PageProps {
  logger: Logger;
  store: StoreApi<Store>;
}

export default PageProps;
