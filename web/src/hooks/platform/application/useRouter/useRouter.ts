import { createElement } from 'react';
import { createBrowserRouter, type Router } from 'react-router';

// containers
import NotFoundScreen from '@/containers/NotFoundScreen';
import SplashScreen from '@/containers/SplashScreen';

// enums
import { ApplicationRouteEnum } from '@/enums/application';

// layouts
import PrivateLayout from '@/layouts/PrivateLayout';
import PublicLayout from '@/layouts/PublicLayout';

// loaders
import loginLoader from '@/loaders/loginLoader';
import privateLoader from '@/loaders/privateLoader';
import setupLoader from '@/loaders/setupLoader';

// pages
import DashboardPage from '@/pages/DashboardPage';
import LandingPage from '@/pages/LandingPage';
import LoginPage from '@/pages/LoginPage';
import SetupPage from '@/pages/SetupPage';

// types
import type { Logger } from '@/utilities/logging/createLogger';

export default function useRouter(logger: Logger): Router {
  return createBrowserRouter(
    [
      {
        HydrateFallback: SplashScreen,
        // errorElement: <RouteError />,
        children: [
          {
            element: createElement(PublicLayout),
            children: [
              {
                index: true,
                element: createElement(LandingPage),
              },
              {
                path: ApplicationRouteEnum.SetupRoute,
                loader: setupLoader(logger),
                element: createElement(SetupPage),
              },
              {
                path: ApplicationRouteEnum.LoginRoute,
                loader: loginLoader(logger),
                element: createElement(LoginPage),
              },
            ],
          },
          {
            id: 'private',
            loader: privateLoader(logger),
            element: createElement(PrivateLayout),
            children: [
              {
                path: ApplicationRouteEnum.DashboardRoute,
                element: createElement(DashboardPage),
              },
            ],
          },
          {
            path: '*',
            element: createElement(NotFoundScreen),
          },
        ],
      },
    ],
    {
      basename: import.meta.env.BASE_URL,
    }
  );
}
