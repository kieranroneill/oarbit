import { ChakraProvider } from '@chakra-ui/react';
import { ThemeProvider } from 'next-themes';
import { type FC } from 'react';
import { I18nextProvider } from 'react-i18next';
import { RouterProvider } from 'react-router/dom';

// utilities
import createRouter from '@/utilities/application/createRouter';
import createStore from '@/utilities/application/createStore';
import createSystem from '@/utilities/application/createSystem';

// types
import type { Props } from './types';

const App: FC<Props> = ({ i18n, logger }) => {
  // hooks
  const store = createStore();
  const router = createRouter({
    logger,
    store,
  });
  const system = createSystem();

  return (
    <ChakraProvider value={system}>
      <ThemeProvider attribute="class" defaultTheme="system" enableSystem={true}>
        <I18nextProvider i18n={i18n}>
          <RouterProvider router={router} />
        </I18nextProvider>
      </ThemeProvider>
    </ChakraProvider>
  );
};

App.displayName = 'App';

export default App;
