// enums
import { APIRouteEnum } from '@/enums/application';

// errors
import { InternalServerError } from '@/errors/internal';

// services
import BaseAPIService from '@/services/BaseAPIService';

// types
import type { RequestOptions } from '@/services/BaseAPIService';
import type { Version } from '@/types/application';

export default class VersionsAPIService extends BaseAPIService {
  // public
  public static displayName = 'VersionsAPIService';

  /**
   * public methods
   */

  public async versions(options?: RequestOptions): Promise<Version> {
    const __functionName = 'versions';
    let response: Response;
    let _error: string;

    try {
      response = await fetch(`${APIRouteEnum.BasePath}${APIRouteEnum.VersionsPath}`, {
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

        this._logger.error(`${VersionsAPIService.displayName}#${__functionName}: ${_error}`);

        throw new InternalServerError(_error);
      }

      return (await response.json()) as Promise<Version>;
    } catch (error) {
      throw this._parseError(error as Error, { logPrefix: `${VersionsAPIService.displayName}#${__functionName}` });
    }
  }
}
