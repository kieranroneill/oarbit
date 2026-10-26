// services
import VesselAPIService from '@/services/VesselsAPIService';

// types
import type { ListResponse } from '@/types/http';
import type { ActionOptions } from '@/types/store';
import type { VesselSpecification } from '@/types/vessels';

export default function fetchVesselSpecifications({
  api,
  logger,
}: ActionOptions): () => Promise<ListResponse<VesselSpecification>> {
  return async () => {
    const __functionName = 'fetchVesselSpecifications';
    const apiService = new VesselAPIService({
      logger,
    });
    let results: ListResponse<VesselSpecification>;

    api.setState((state) => ({
      ...state,
      fetchingVesselSpecifications: true,
    }));

    try {
      results = await apiService.specifications();

      logger.debug(`${__functionName}:`, results);

      api.setState((state) => ({
        ...state,
        vesselSpecifications: results.items,
        fetchingVesselSpecifications: false,
      }));

      return results;
    } catch (error) {
      api.setState((state) => ({
        ...state,
        fetchingVesselSpecifications: false,
      }));

      throw error;
    }
  };
}
