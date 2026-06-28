// actions
import { fetchBoat } from './actions';

// types
import type { ActionOptions, SliceOptions, StateCreator } from '@/types/store';
import type { Slice } from './types';

const createBoatSlice: (options: SliceOptions) => StateCreator<Slice> =
  ({ logger }) =>
  (setState, getState) => {
    const actionOptions: ActionOptions = {
      api: { getState, setState },
      logger,
    };

    return {
      // actions
      fetchBoat: fetchBoat(actionOptions),
      // state
      boat: null,
      fetchingBoat: false,
    };
  };

export default createBoatSlice;
