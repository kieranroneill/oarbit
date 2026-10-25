// types
import type { User } from '../users';

type AuthenticationStatus =
  | { setupRequired: true; user: null }
  | { setupRequired: false; user: null }
  | {
      setupRequired: false;
      user: User;
    };

export default AuthenticationStatus;
