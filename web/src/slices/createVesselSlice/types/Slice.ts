// types
import type { ListResponse } from '@/types/http';
import type { VesselSpecification } from '@/types/vessels';

interface Slice {
  // actions
  fetchVesselSpecifications: () => Promise<ListResponse<VesselSpecification>>;
  // state
  vesselSpecifications: VesselSpecification[];
  fetchingVesselSpecifications: boolean;
}

export default Slice;
