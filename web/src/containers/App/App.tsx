import { type FC } from 'react';
import { I18nextProvider } from 'react-i18next';
import { RouterProvider } from 'react-router/dom';

// hooks
import useRouter from '@/hooks/platform/application/useRouter';

// types
import type { Props } from './types';

const App: FC<Props> = ({ i18n, logger }) => {
  const router = useRouter(logger);

  return (
    <I18nextProvider i18n={i18n}>
      <RouterProvider router={router} />
    </I18nextProvider>
  );
};

App.displayName = 'App';

export default App;
