// enums
import { LogLevelEnum } from '@/enums/logging';

// types
import type { BaseAPIServiceOptions } from './types';
import type { Logger } from '@/types/logging';

// utilities
import { createLogger } from '@/utilities/logging';

export default abstract class BaseAPIService {
  // protected
  protected _logger: Logger;

  public constructor({ logger }: BaseAPIServiceOptions = {}) {
    this._logger = logger ?? createLogger((process.env.NEXT_PUBLIC_LOG_LEVEL as LogLevelEnum) ?? LogLevelEnum.Error);
  }
}
