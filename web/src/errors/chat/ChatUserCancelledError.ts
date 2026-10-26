// enums
import { ErrorCodeEnum } from '@/enums/application';

// errors
import BaseError from '../_base/BaseError';

export default class ChatUserCancelledError extends BaseError {
  public readonly code: ErrorCodeEnum = ErrorCodeEnum.ChatUserCancelledError;
  public readonly type: string = 'ChatUserCancelledError';
}
