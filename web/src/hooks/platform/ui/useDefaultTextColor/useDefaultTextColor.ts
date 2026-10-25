// types
import type { Theme } from '@/types/ui';

export default function useDefaultTextColor(): Theme {
  return {
    light: 'gray.600',
    dark: 'whiteAlpha.800',
  };
}
