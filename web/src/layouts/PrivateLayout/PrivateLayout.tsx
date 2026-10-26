import { Center } from '@chakra-ui/react';
import { type FC } from 'react';
import { Outlet } from 'react-router';

// types
import type { LayoutProps } from '@/types/application';

const PrivateLayout: FC<LayoutProps> = () => {
  return (
    <Center as="main" minH="100dvh">
      <Outlet />
    </Center>
  );
};

PrivateLayout.displayName = 'PrivateLayout';

export default PrivateLayout;
