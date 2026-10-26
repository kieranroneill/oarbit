import { Center, Text, VStack } from '@chakra-ui/react';
import { type FC, useEffect, useState } from 'react';
import { useTranslation } from 'react-i18next';

// components
import CircularProgressWithIcon from '@/components/feedback/CircularProgressWithIcon';

// constants
import { DEFAULT_GAP } from '@/constants';

// icons
import OBLogo from '@/icons/OBLogo';

const SplashScreen: FC = () => {
  const { t } = useTranslation();
  const expressions = t('captions.loadingExpressions', { returnObjects: true }) as string[];
  // states
  const [expressionIndex, setExpressionIndex] = useState(() => Math.floor(Math.random() * expressions.length));

  useEffect(() => {
    if (expressions.length < 2) return;

    const timer = window.setInterval(() => {
      setExpressionIndex((previous) => {
        const offset = 1 + Math.floor(Math.random() * (expressions.length - 1));
        return (previous + offset) % expressions.length;
      });
    }, 4000);

    return () => window.clearInterval(timer);
  }, [expressions.length]);

  return (
    <Center as="main" minH="100dvh">
      <VStack gap={DEFAULT_GAP}>
        <CircularProgressWithIcon icon={OBLogo} />

        <Text textAlign="center">{expressions[expressionIndex % expressions.length]}</Text>
      </VStack>
    </Center>
  );
};

SplashScreen.displayName = 'SplashScreen';

export default SplashScreen;
