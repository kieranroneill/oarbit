// enums
import { APIRouteEnum } from '@/enums/application';

// errors
import { InternalServerError } from '@/errors/internal';

// services
import BaseAPIService from '@/services/BaseAPIService';

// types
import type { RequestOptions } from '@/services/BaseAPIService';
import type { AuthenticationStatus } from '@/types/authentication';

export default class AuthenticationAPIService extends BaseAPIService {
  // public
  public static displayName = 'AuthenticationAPIService';

  /**
   * public methods
   */

  public async status(options?: RequestOptions): Promise<AuthenticationStatus> {
    const __functionName = 'status';
    let response: Response;
    let _error: string;

    try {
      response = await fetch(`${APIRouteEnum.BasePath}${APIRouteEnum.AuthPath}${APIRouteEnum.StatusPath}`, {
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

        this._logger.error(`${AuthenticationAPIService.displayName}#${__functionName}: ${_error}`);

        // TODO: handle unauthorized events and when the response is bad
        throw new InternalServerError(_error);
      }

      return (await response.json()) as Promise<AuthenticationStatus>;
    } catch (error) {
      throw this._parseError(error as Error, {
        logPrefix: `${AuthenticationAPIService.displayName}#${__functionName}`,
      });
    }
  }
}
