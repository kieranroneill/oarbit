// actions
import { fetchVesselSpecifications } from './actions';

// types
import type { Slice } from './types';
import type { ActionOptions, SliceOptions, StateCreator } from '@/types/store';

const createVesselSlice: (options: SliceOptions) => StateCreator<Slice> =
  ({ logger }) =>
  (setState, getState) => {
    const actionOptions: ActionOptions = {
      api: { getState, setState },
      logger,
    };

    return {
      // actions
      fetchVesselSpecifications: fetchVesselSpecifications(actionOptions),
      // state
      vesselSpecifications: [],
      fetchingVesselSpecifications: false,
    };
  };

export default createVesselSlice;
