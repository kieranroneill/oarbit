import { type FC } from 'react';
import { Outlet } from 'react-router';

const PrivateLayout: FC = () => {
  return (
    <main>
      <Outlet />
    </main>
  );
};

PrivateLayout.displayName = 'PrivateLayout';

export default PrivateLayout;
