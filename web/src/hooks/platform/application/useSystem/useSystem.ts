import { createSystem, defaultConfig, SystemContext } from '@chakra-ui/react';

export default function useSystem(): SystemContext {
  return createSystem(defaultConfig, {
    globalCss: {
      html: {
        scrollBehavior: 'smooth',
      },
      body: {
        fontFamily: 'body',
        fontWeight: 400,
      },
      'code, pre': {
        fontFamily: 'mono',
      },
    },
    theme: {
      semanticTokens: {
        colors: {
          fg: {
            value: {
              _dark: '{colors.whiteAlpha.800}',
              base: '{colors.gray.600}',
            },
          },

          primary: {
            solid: {
              value: {
                base: '{colors.primary.500}',
                _dark: '{colors.primary.300}',
              },
            },
            contrast: {
              value: {
                // black is required for accessible text on primary.500.
                base: '{colors.black}',
                _dark: '{colors.primary.900}',
              },
            },
            fg: {
              value: {
                base: '{colors.primary.700}',
                _dark: '{colors.primary.200}',
              },
            },
            muted: {
              value: {
                base: '{colors.primary.100}',
                _dark: '{colors.primary.800}',
              },
            },
            subtle: {
              value: {
                base: '{colors.primary.50}',
                _dark: '{colors.primary.900}',
              },
            },
            emphasized: {
              value: {
                base: '{colors.primary.200}',
                _dark: '{colors.primary.700}',
              },
            },
            focusRing: {
              value: {
                base: '{colors.primary.500}',
                _dark: '{colors.primary.300}',
              },
            },
          },
        },
      },
      tokens: {
        colors: {
          primary: {
            50: { value: '#FFF7EE' },
            100: { value: '#FFE9D4' },
            200: { value: '#FFD4AD' },
            300: { value: '#F9BA79' },
            400: { value: '#F39B4A' },
            500: { value: '#F3852B' },
            600: { value: '#DC6E17' },
            700: { value: '#BC570F' },
            800: { value: '#8F4210' },
            900: { value: '#713510' },
          },
        },
        fonts: {
          heading: { value: `'Geist Variable', sans-serif` },
          body: { value: `'Geist Variable', sans-serif` },
          mono: { value: `'Geist Mono Variable', monospace` },
        },
      },
    },
  });
}
