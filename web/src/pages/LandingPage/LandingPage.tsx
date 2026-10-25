import { Button, Heading, Icon, VStack, Text } from '@chakra-ui/react';
import { type FC, useMemo } from 'react';
import { useTranslation } from 'react-i18next';
import { Link } from 'react-router';

// constants
import { DEFAULT_GAP } from '@/constants';

// enums
import { ApplicationRouteEnum } from '@/enums/application';

// icons
import OBLogo from '@/icons/OBLogo';

// utilities
import calculateIconSize from '@/utilities/ui/calculateIconSize';

const LandingPage: FC = () => {
  const { t } = useTranslation();
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
        <Text>{t('captions.landingPageChoice')}</Text>

        <Button asChild={true} colorPalette="primary" size="lg" width="full">
          <Link to={ApplicationRouteEnum.SetupRoute}>{t('labels.setup')}</Link>
        </Button>

        <Text>{t('labels.or')}</Text>

        <Button asChild={true} colorPalette="primary" size="lg" width="full">
          <Link to={ApplicationRouteEnum.LoginRoute}>{t('labels.login')}</Link>
        </Button>
      </VStack>
    </VStack>
  );
};

LandingPage.displayName = 'LandingPage';

export default LandingPage;
