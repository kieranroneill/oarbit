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
import landingLoader from '@/loaders/landingLoader';
import loginLoader from '@/loaders/loginLoader';
import privateLoader from '@/loaders/privateLoader';
import setupLoader from '@/loaders/setupLoader';

// middlewares
import versionsMiddleware from '@/middlewares/versionsMiddleware';

// pages
import DashboardPage from '@/pages/DashboardPage';
import LandingPage from '@/pages/LandingPage';
import LoginPage from '@/pages/LoginPage';
import SetupPage from '@/pages/SetupPage';

// types
import type { Options } from './types';
import type { LayoutProps, PageProps } from '@/types/application';

export default function createRouter({ logger, store }: Options): Router {
  const baseLayoutProps: LayoutProps = { logger, store };
  const basePageProps: PageProps = { logger, store };

  return createBrowserRouter(
    [
      {
        HydrateFallback: SplashScreen,
        children: [
          {
            element: createElement(PublicLayout, baseLayoutProps),
            children: [
              {
                index: true,
                loader: landingLoader(logger),
                element: createElement(LandingPage, basePageProps),
              },
              {
                path: ApplicationRouteEnum.SetupRoute,
                loader: setupLoader(logger),
                element: createElement(SetupPage, basePageProps),
              },
              {
                path: ApplicationRouteEnum.LoginRoute,
                loader: loginLoader(logger),
                element: createElement(LoginPage, basePageProps),
              },
            ],
          },
          {
            id: 'private',
            loader: privateLoader(logger),
            element: createElement(PrivateLayout, baseLayoutProps),
            children: [
              {
                path: ApplicationRouteEnum.DashboardRoute,
                element: createElement(DashboardPage, basePageProps),
              },
            ],
          },
          {
            path: '*',
            element: createElement(NotFoundScreen),
          },
        ],
        // errorElement: <RouteError />,
        middleware: [versionsMiddleware(logger, store)],
      },
    ],
    {
      basename: import.meta.env.BASE_URL,
    }
  );
}
