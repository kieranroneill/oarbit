enum ErrorCodeEnum {
  // general
  UnknownError = 1000,

  // authentication
  InvalidPasswordError = 2000,

  // network
  InternalServerError = 5000,
  ParsingError = 5001,
}

export default ErrorCodeEnum;
