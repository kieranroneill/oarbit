import { Center, Flex, HStack, IconButton, Text } from '@chakra-ui/react';
import { type FC, useCallback, useMemo } from 'react';
import { useTranslation } from 'react-i18next';
import { LuArrowLeft } from 'react-icons/lu';
import { Outlet, useMatch, useNavigate } from 'react-router';

// enums
import { ApplicationRouteEnum } from '@/enums/application';

// types
import type { LayoutProps } from '@/types/application';

const PublicLayout: FC<LayoutProps> = ({ store }) => {
  const { t } = useTranslation();
  const navigate = useNavigate();
  const setupMatch = useMatch(ApplicationRouteEnum.SetupRoute);
  const loginMatch = useMatch(ApplicationRouteEnum.LoginRoute);
  const showBackButton = Boolean(setupMatch || loginMatch);
  // store
  const { environment, version } = store.getState();
  // memos
  const defaultPadding = useMemo(() => 6, []);
  // callbacks
  const handleOnBackClick = useCallback(() => navigate(-1), [navigate]);

  return (
    <Flex direction="column" minH="100dvh">
      {/*header*/}
      <HStack as="header" p={defaultPadding}>
        {showBackButton && (
          <IconButton aria-label={t('labels.back')} onClick={handleOnBackClick} variant="ghost">
            <LuArrowLeft />
          </IconButton>
        )}
      </HStack>

      {/*main*/}
      <Center as="main" flex="1" px={defaultPadding}>
        <Outlet />
      </Center>

      {/*footer*/}
      <Center as="footer" flexShrink="0" p={defaultPadding}>
        <Text color="fg.muted" fontSize="sm">
          {`${import.meta.env.VITE_TITLE}${version && ` v${version}`}${environment && ` (${environment})`}`}
        </Text>
      </Center>
    </Flex>
  );
};

PublicLayout.displayName = 'PublicLayout';

export default PublicLayout;
