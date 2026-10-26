// types
import type { Sizes } from '@/types/ui';

export default function calculateIconSize(size?: Sizes): number {
  switch (size) {
    case 'xs':
      return 3;
    case 'md':
      return 6;
    case 'lg':
      return 10;
    case 'xl':
      return 16;
    case '2xl':
      return 24;
    case '3xl':
      return 32;
    case '4xl':
      return 48;
    case 'sm':
    default:
      return 4;
  }
}
