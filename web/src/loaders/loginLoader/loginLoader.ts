import { redirect } from 'react-router';

// enums
import { ApplicationRouteEnum } from '@/enums/application';

// services
import AuthenticationAPIService from '@/services/AuthenticationAPIService';

// types
import type { Logger } from '@/utilities/logging/createLogger';

export default function loginLoader(logger: Logger): () => Promise<null> {
  return async () => {
    const __functionName = 'loginLoader';
    const status = await new AuthenticationAPIService({
      logger,
    }).status();

    if (status.setupRequired) {
      logger.debug(`${__functionName} - setup required, redirecting to "${ApplicationRouteEnum.SetupRoute}"`);

      throw redirect(ApplicationRouteEnum.SetupRoute);
    }

    if (status.user) {
      logger.debug(
        `${__functionName} - user "${status.user.id}" logged in, redirecting to "${ApplicationRouteEnum.DashboardRoute}"`
      );

      throw redirect(ApplicationRouteEnum.DashboardRoute);
    }

    return null;
  };
}
