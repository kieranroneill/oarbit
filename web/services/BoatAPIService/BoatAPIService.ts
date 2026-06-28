// enums
import { RouteEnum } from '@/enums/api';

// errors
import { BaseError, InternalServerError } from '@/errors';

// services
import BaseAPIService from '@/services/BaseAPIService';

// types
import type { RequestOptions } from '@/services/BaseAPIService';
import type { BoatContext } from '@/types/boat';

export default class BoatAPIService extends BaseAPIService {
  // public
  public static displayName = 'BoatAPIService';

  /**
   * public methods
   */

  public async getBoatByID(profileID: string, options?: RequestOptions): Promise<BoatContext> {
    const __functionName = 'getBoatByID';
    let response: Response;
    let _error: string;

    try {
      response = await fetch(`${RouteEnum.Base}${RouteEnum.Boat}/${profileID}`, {
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

        this._logger.debug(`${BoatAPIService.displayName}#${__functionName}: ${_error}`);

        throw new InternalServerError(_error);
      }

      return (await response.json()) as Promise<BoatContext>;
    } catch (error) {
      if ((error as BaseError).isOarbitError()) {
        throw error;
      }

      // TODO: handle unauthorized events and when the response is bad
      console.log(error);

      throw error;
    }
  }
}
