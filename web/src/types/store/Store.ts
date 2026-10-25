// types
import type { Slice as ApplicationSlice } from '@/slices/createApplicationSlice';
import type { Slice as VesselSlice } from '@/slices/createVesselSlice';

type Store = ApplicationSlice | VesselSlice;

export default Store;
