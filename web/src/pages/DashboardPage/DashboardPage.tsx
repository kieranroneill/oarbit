import { VStack } from '@chakra-ui/react';
import { type FC } from 'react';
import { useTranslation } from 'react-i18next';

// constants
import { DEFAULT_GAP } from '@/constants';

// types
import type { PageProps } from '@/types/application';

const DashboardPage: FC<PageProps> = () => {
  const { t } = useTranslation();

  return (
    <>
      <title>{`${t('titles.dashboard')} | ${import.meta.env.VITE_TITLE}`}</title>

      <VStack as="section" gap={DEFAULT_GAP + 6}>
        Setup
      </VStack>
    </>
  );
};

DashboardPage.displayName = 'DashboardPage';

export default DashboardPage;
