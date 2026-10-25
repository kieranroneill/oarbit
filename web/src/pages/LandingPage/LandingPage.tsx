import { Button, Heading, Icon, VStack, Text } from '@chakra-ui/react';
import { type FC, useMemo } from 'react';
import { useTranslation } from 'react-i18next';
import { Link, useLoaderData } from 'react-router';

// constants
import { DEFAULT_GAP } from '@/constants';

// enums
import { ApplicationRouteEnum } from '@/enums/application';

// icons
import OBLogo from '@/icons/OBLogo';

// loaders
import type landingLoader from '@/loaders/landingLoader';

// types
import type { PageProps } from '@/types/application';

// utilities
import calculateIconSize from '@/utilities/ui/calculateIconSize';

const LandingPage: FC<PageProps> = () => {
  const { t } = useTranslation();
  const { setupRequired } = useLoaderData<ReturnType<typeof landingLoader>>();
  // memos
  const iconSize = useMemo(() => calculateIconSize('3xl'), []);

  return (
    <VStack as="section" gap={DEFAULT_GAP + 6}>
      {/*welcome message*/}
      <VStack gap={DEFAULT_GAP - 2}>
        <Icon as={OBLogo} boxSize={iconSize} />

        <Heading size="3xl">{t('headings.welcome')}</Heading>

        <Text>{t('captions.welcomeDescription')}</Text>
      </VStack>

      {/*setup or login ctas*/}
      <VStack gap={DEFAULT_GAP - 2} width="full">
        {!setupRequired && <Text>{t('captions.landingPageChoice')}</Text>}

        <Button asChild={true} colorPalette="primary" size="lg" width="full">
          <Link to={ApplicationRouteEnum.SetupRoute}>{t('labels.setup')}</Link>
        </Button>

        {!setupRequired && (
          <>
            <Text>{t('labels.or')}</Text>

            <Button asChild={true} colorPalette="primary" size="lg" width="full">
              <Link to={ApplicationRouteEnum.LoginRoute}>{t('labels.login')}</Link>
            </Button>
          </>
        )}
      </VStack>
    </VStack>
  );
};

LandingPage.displayName = 'LandingPage';

export default LandingPage;
