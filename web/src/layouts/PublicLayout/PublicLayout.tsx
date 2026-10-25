import { Center } from '@chakra-ui/react';
import { type FC } from 'react';
import { Outlet } from 'react-router';

const PublicLayout: FC = () => {
  return (
    <Center as="main" minH="100dvh">
      <Outlet />
    </Center>
  );
};

PublicLayout.displayName = 'PublicLayout';

export default PublicLayout;
