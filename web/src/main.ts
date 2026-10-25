import I18next, { type i18n as I18n } from 'i18next';
import { createElement } from 'react';
import { initReactI18next } from 'react-i18next';
import { createRoot } from 'react-dom/client';

// components
import App from '@/containers/App';

// enums
import { LogLevelEnum } from '@/utilities/logging/createLogger';

// fonts
import '@fontsource-variable/geist';
import '@fontsource-variable/geist-mono';

// translations
import { en } from '@/translations';

// utilities
import createLogger from '@/utilities/logging/createLogger';

export async function onDOMContentLoaded(): Promise<void> {
  const logger = createLogger((import.meta.env.VITE_LOG_LEVEL as LogLevelEnum) ?? LogLevelEnum.Error);
  const rootElement = document.getElementById('root');
  let i18n: I18n;

  if (!rootElement) {
    logger.error(`failed to find "root" element to render react app`);

    return;
  }

  i18n = I18next.use(initReactI18next);

  await i18n.init({
    fallbackLng: 'en',
    debug: true,
    interpolation: {
      escapeValue: false,
    },
    resources: {
      en: {
        translation: en,
      },
    },
  });

  createRoot(rootElement).render(
    createElement(App, {
      i18n,
      logger,
    })
  );
}

window.addEventListener('DOMContentLoaded', onDOMContentLoaded);
