// types
import type ActionAPI from './ActionAPI';
import type { Logger } from '@/utilities/logging/createLogger';

interface ActionOptions {
  api: ActionAPI;
  logger: Logger;
}

export default ActionOptions;
