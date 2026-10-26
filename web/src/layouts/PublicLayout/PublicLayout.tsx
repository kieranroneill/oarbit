import { Center, Flex, HStack, IconButton, Spacer, Text } from '@chakra-ui/react';
import { useTheme } from 'next-themes';
import { type FC, useCallback, useMemo } from 'react';
import { useTranslation } from 'react-i18next';
import { LuArrowLeft, LuMonitor, LuMoon, LuSun } from 'react-icons/lu';
import { Outlet, useMatch, useNavigate } from 'react-router';

// enums
import { ApplicationRouteEnum } from '@/enums/application';

// types
import type { LayoutProps } from '@/types/application';

const PublicLayout: FC<LayoutProps> = ({ store }) => {
  const { t } = useTranslation();
  const { theme = 'system', setTheme } = useTheme();
  const navigate = useNavigate();
  const setupMatch = useMatch(ApplicationRouteEnum.SetupRoute);
  const loginMatch = useMatch(ApplicationRouteEnum.LoginRoute);
  const showBackButton = Boolean(setupMatch || loginMatch);
  // store
  const { environment, version } = store.getState();
  // memos
  const defaultPadding = useMemo(() => 6, []);
  const nextTheme = useMemo(() => (theme === 'light' ? 'dark' : theme === 'dark' ? 'system' : 'light'), [theme]);
  const themeLabel = useMemo(
    () =>
      t('labels.changeTheme', {
        current: t(`labels.theme_${theme}`),
        next: t(`labels.theme_${nextTheme}`),
      }),
    [theme, nextTheme, t]
  );
  // callbacks
  const handleOnBackClick = useCallback(() => navigate(-1), [navigate]);
  const handleOnThemeChangeClick = useCallback(() => setTheme(nextTheme), [nextTheme, setTheme]);

  return (
    <Flex direction="column" minH="100dvh">
      {/*header*/}
      <HStack as="header" p={defaultPadding}>
        {showBackButton && (
          <IconButton
            aria-label={t('labels.back')}
            onClick={handleOnBackClick}
            title={t('labels.back')}
            variant="ghost"
          >
            <LuArrowLeft />
          </IconButton>
        )}

        <Spacer />

        <IconButton aria-label={themeLabel} onClick={handleOnThemeChangeClick} title={themeLabel} variant="ghost">
          {theme === 'light' ? <LuSun /> : theme === 'dark' ? <LuMoon /> : <LuMonitor />}
        </IconButton>
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
