import { VStack } from '@chakra-ui/react';
import { type FC } from 'react';
import { useTranslation } from 'react-i18next';

// constants
import { DEFAULT_GAP } from '@/constants';

// types
import type { PageProps } from '@/types/application';

const SetupPage: FC<PageProps> = () => {
  const { t } = useTranslation();

  return (
    <>
      <title>{`${t('titles.setup')} | ${import.meta.env.VITE_TITLE}`}</title>

      <VStack as="section" gap={DEFAULT_GAP + 6}>
        Setup
      </VStack>
    </>
  );
};

SetupPage.displayName = 'SetupPage';

export default SetupPage;
