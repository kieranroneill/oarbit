import { Center, Flex, Text } from '@chakra-ui/react';
import { type FC } from 'react';
import { Outlet } from 'react-router';

// types
import type { LayoutProps } from '@/types/application';

const PublicLayout: FC<LayoutProps> = ({ store }) => {
  const { environment, version } = store.getState();

  return (
    <Flex direction="column" minH="100dvh">
      <Center as="main" flex="1" p="6">
        <Outlet />
      </Center>

      <Center as="footer" flexShrink="0" p="4">
        <Text color="fg.muted" fontSize="sm">
          {`${import.meta.env.VITE_TITLE}${version && ` v${version}`}${environment && ` (${environment})`}`}
        </Text>
      </Center>
    </Flex>
  );
};

PublicLayout.displayName = 'PublicLayout';

export default PublicLayout;
