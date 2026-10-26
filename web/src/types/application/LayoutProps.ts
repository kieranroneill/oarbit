import { type StoreApi } from 'zustand';

// types
import type { Store } from '@/types/store';
import type { Logger } from '@/utilities/logging/createLogger';

interface LayoutProps {
  logger: Logger;
  store: StoreApi<Store>;
}

export default LayoutProps;
