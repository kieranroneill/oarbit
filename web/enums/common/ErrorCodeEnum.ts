enum ErrorCodeEnum {
  // general
  UnknownError = 1000,
  MalformedDataError = 1001,
  InternalServerError = 1500,

  // authentication
  InvalidPasswordError = 2000,
}

export default ErrorCodeEnum;
