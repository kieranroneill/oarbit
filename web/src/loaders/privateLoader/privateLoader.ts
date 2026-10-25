import { redirect } from 'react-router';

// enums
import { ApplicationRouteEnum } from '@/enums/application';

// services
import AuthenticationAPIService from '@/services/AuthenticationAPIService';

// types
import type { Logger } from '@/utilities/logging/createLogger';

export default function privateLoader(logger: Logger): () => Promise<void> {
  return async () => {
    const __functionName = 'privateLoader';
    const status = await new AuthenticationAPIService({
      logger,
    }).status();

    if (status.setupRequired) {
      logger.debug(`${__functionName} - setup required, redirecting to "${ApplicationRouteEnum.SetupRoute}"`);

      throw redirect(ApplicationRouteEnum.SetupRoute);
    }

    if (!status.user) {
      logger.debug(`${__functionName} - user not logged in, redirecting to "${ApplicationRouteEnum.LoginRoute}"`);

      throw redirect(ApplicationRouteEnum.LoginRoute);
    }
  };
}
