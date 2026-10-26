import { VStack } from '@chakra-ui/react';
import { type FC } from 'react';
import { useTranslation } from 'react-i18next';

// constants
import { DEFAULT_GAP } from '@/constants';

// types
import type { PageProps } from '@/types/application';

const LoginPage: FC<PageProps> = () => {
  const { t } = useTranslation();

  return (
    <>
      <title>{`${t('titles.login')} | ${import.meta.env.VITE_TITLE}`}</title>

      <VStack as="section" gap={DEFAULT_GAP + 6}>
        Login
      </VStack>
    </>
  );
};

LoginPage.displayName = 'LoginPage';

export default LoginPage;
