// enums
import { LogLevelEnum } from '@/utilities/logging/createLogger';

// errors
import { BaseError } from '@/errors/_base';
import { InternalServerError } from '@/errors/internal';

// types
import type { BaseAPIServiceOptions, ParseErrorOptions } from './types';
import type { Logger } from '@/utilities/logging/createLogger';

// utilities
import createLogger from '@/utilities/logging/createLogger';

export default abstract class BaseAPIService {
  // protected
  protected _logger: Logger;

  public constructor({ logger }: BaseAPIServiceOptions = {}) {
    this._logger = logger ?? createLogger((import.meta.env.VITE_LOG_LEVEL as LogLevelEnum) ?? LogLevelEnum.Error);
  }

  protected _parseError(error: Error, { logPrefix }: ParseErrorOptions): BaseError {
    if ((error as BaseError).isOarbitError()) {
      return error as BaseError;
    }

    this._logger.error(`${logPrefix}:`, error);

    return new InternalServerError(error.message);
  }
}
