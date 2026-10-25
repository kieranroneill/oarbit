import { type FC } from 'react';
import { Outlet } from 'react-router';

const PublicLayout: FC = () => {
  return (
    <main>
      <Outlet />
    </main>
  );
};

PublicLayout.displayName = 'PublicLayout';

export default PublicLayout;
