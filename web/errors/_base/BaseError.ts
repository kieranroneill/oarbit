// enums
import { ErrorCodeEnum } from '@/enums/common';

export default abstract class BaseError extends Error {
  // private
  private readonly __isOarbitError = true;
  // public
  public readonly code: ErrorCodeEnum;
  public readonly displayName: string;
  public message: string;

  public constructor(message: string) {
    super(message.toLowerCase());
  }

  public isOarbitError(): boolean {
    return this.__isOarbitError;
  }
}
