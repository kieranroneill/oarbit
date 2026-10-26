import { redirect } from 'react-router';

// enums
import { ApplicationRouteEnum } from '@/enums/application';

// services
import AuthenticationAPIService from '@/services/AuthenticationAPIService';

// types
import type { Logger } from '@/utilities/logging/createLogger';

export default function setupLoader(logger: Logger): () => Promise<null> {
  return async () => {
    const __functionName = 'privateLoader';
    const status = await new AuthenticationAPIService({
      logger,
    }).status();

    if (!status.setupRequired) {
      if (!status.user) {
        logger.debug(
          `${__functionName} - setup not required and user not logged in, redirecting to "${ApplicationRouteEnum.LoginRoute}"`
        );

        throw redirect(ApplicationRouteEnum.LoginRoute);
      }

      logger.debug(
        `${__functionName} - setup not required, and user "${status.user.id}" logged in, redirecting to "${ApplicationRouteEnum.DashboardRoute}"`
      );

      throw redirect(ApplicationRouteEnum.DashboardRoute);
    }

    return null;
  };
}
