// enums
import { ErrorCodeEnum } from '@/enums/application';

// errors
import BaseError from '../_base/BaseError';

export default class InternalServerError extends BaseError {
  public readonly code: ErrorCodeEnum = ErrorCodeEnum.InternalServerError;
  public readonly type: string = 'InternalServerError';
}
