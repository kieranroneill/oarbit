import { Center } from '@chakra-ui/react';
import { type FC } from 'react';
import { Outlet } from 'react-router';

const PrivateLayout: FC = () => {
  return (
    <Center as="main" minH="100dvh">
      <Outlet />
    </Center>
  );
};

PrivateLayout.displayName = 'PrivateLayout';

export default PrivateLayout;
