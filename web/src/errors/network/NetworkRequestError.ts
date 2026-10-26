// enums
import { ErrorCodeEnum } from '@/enums/application';

// errors
import BaseError from '../_base/BaseError';

export default class NetworkRequestError extends BaseError {
  public readonly code: ErrorCodeEnum = ErrorCodeEnum.NetworkRequestError;
  public readonly displayName: string = 'NetworkRequestError';
}
