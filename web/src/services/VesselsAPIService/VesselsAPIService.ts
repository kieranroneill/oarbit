// enums
import { APIRouteEnum } from '@/enums/application';

// errors
import { InternalServerError } from '@/errors/internal';

// services
import BaseAPIService from '@/services/BaseAPIService';

// types
import type { RequestOptions } from '@/services/BaseAPIService';
import type { ListResponse } from '@/types/http';
import type { VesselSpecification } from '@/types/vessels';

export default class VesselsAPIService extends BaseAPIService {
  // public
  public static displayName = 'VesselAPIService';

  /**
   * public methods
   */

  public async getVesselSpecifications(options?: RequestOptions): Promise<ListResponse<VesselSpecification>> {
    const __functionName = 'getVesselSpecifications';
    let response: Response;
    let _error: string;

    try {
      response = await fetch(`${APIRouteEnum.BasePath}${APIRouteEnum.VesselsPath}${APIRouteEnum.SpecificationsPath}`, {
        headers: {
          'Content-Type': 'application/json',
        },
        method: 'GET',
        ...(options?.signal && {
          signal: options.signal,
        }),
      });

      if (!response.ok || !response.body) {
        _error = `failed to get response - ${response.status} - ${response.body}`;

        this._logger.error(`${VesselsAPIService.displayName}#${__functionName}: ${_error}`);

        // TODO: handle unauthorized events and when the response is bad
        throw new InternalServerError(_error);
      }

      return (await response.json()) as Promise<ListResponse<VesselSpecification>>;
    } catch (error) {
      throw this._parseError(error as Error, { logPrefix: `${VesselsAPIService.displayName}#${__functionName}` });
    }
  }
}
