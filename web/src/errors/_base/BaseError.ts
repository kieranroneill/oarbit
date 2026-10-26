// enums
import { ErrorCodeEnum } from '@/enums/application';

export default abstract class BaseError extends Error {
  // private
  private readonly __isOarbitError = true;
  // public
  public readonly code: ErrorCodeEnum;
  public message: string;
  public readonly type: string;

  public constructor(message: string) {
    super(message.toLowerCase());
  }

  public isOarbitError(): boolean {
    return this.__isOarbitError;
  }
}
