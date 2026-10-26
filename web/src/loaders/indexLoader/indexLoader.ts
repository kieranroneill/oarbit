import { redirect } from 'react-router';

// enums
import { ApplicationRouteEnum } from '@/enums/application';

// services
import AuthenticationAPIService from '@/services/AuthenticationAPIService';

// types
import type { AuthenticationStatus } from '@/types/authentication';
import type { Logger } from '@/utilities/logging/createLogger';

export default function indexLoader(logger: Logger): () => Promise<AuthenticationStatus> {
  return async () => {
    const __functionName = 'indexLoader';
    const status = await new AuthenticationAPIService({
      logger,
    }).status();

    if (!status.setupRequired && status.user) {
      logger.debug(
        `${__functionName} - user "${status.user.id}" logged in, redirecting to "${ApplicationRouteEnum.DashboardRoute}"`
      );

      throw redirect(ApplicationRouteEnum.DashboardRoute);
    }

    return status;
  };
}
