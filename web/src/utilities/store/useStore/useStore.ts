import { create, type StoreApi } from 'zustand';
import { devtools, persist, type PersistOptions } from 'zustand/middleware';

// constants
import { STORE_NAME } from '@/constants';

// enums
import { LogLevelEnum } from '@/utilities/logging/createLogger';

// slices
import createVesselSlice from '@/slices/createVesselSlice';

// types
import type { PersistedState, SliceOptions, Store } from '@/types/store';

// utilities
import createLogger from '@/utilities/logging/createLogger';

export default function useStore(): StoreApi<Store> {
  const sliceOptions: SliceOptions = {
    logger: createLogger((import.meta.env.VITE_LOG_LEVEL as LogLevelEnum) ?? LogLevelEnum.Error),
  };

  return create<Store>()(
    devtools(
      persist(
        (...api) => ({
          ...createVesselSlice(sliceOptions)(...api),
        }),
        {
          name: STORE_NAME,
          // partialize: ({ boat }) => ({
          //   boat,
          // }),
          version: 0,
        } as PersistOptions<Store, PersistedState>
      ),
      {
        name: STORE_NAME,
      }
    )
  );
}
