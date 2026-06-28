'use client';
import { devtools, persist, type PersistOptions } from 'zustand/middleware';
import { createStore, type StoreApi } from 'zustand/vanilla';

// constants
import { STORE_NAME } from '@/constants';

// enums
import { LogLevelEnum } from '@/enums/logging';

// slices
import createBoatSlice from '@/slices/createBoatSlice';

// types
import type { PersistedState, SliceOptions, Store } from '@/types/store';

// utilities
import { createLogger } from '@/utilities/logging';

const store: () => StoreApi<Store> = () => {
  const sliceOptions: SliceOptions = {
    logger: createLogger((process.env.NEXT_PUBLIC_LOG_LEVEL as LogLevelEnum) ?? LogLevelEnum.Error),
  };

  return createStore<Store>()(
    devtools(
      persist(
        (...api) => ({
          ...createBoatSlice(sliceOptions)(...api),
        }),
        {
          name: STORE_NAME,
          partialize: ({ boat }) => ({
            boat,
          }),
          version: 0,
        } as PersistOptions<Store, PersistedState>
      ),
      {
        name: STORE_NAME,
      }
    )
  );
};

export default store;
