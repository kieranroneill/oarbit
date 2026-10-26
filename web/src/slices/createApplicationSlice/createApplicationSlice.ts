// types
import type { Slice } from './types';
import type { SliceOptions, StateCreator } from '@/types/store';

const createApplicationSlice: (options: SliceOptions) => StateCreator<Slice> = () => () => {
  return {
    // state
    environment: null,
    version: null,
  };
};

export default createApplicationSlice;
