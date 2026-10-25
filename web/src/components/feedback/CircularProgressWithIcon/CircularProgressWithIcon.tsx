import { AbsoluteCenter, Icon, ProgressCircle } from '@chakra-ui/react';
import { type FC, useMemo } from 'react';

// hooks
import useSubTextColor from '@/hooks/ui/useSubTextColor';

// types
import type { Props } from './types';
import type { Theme } from '@/types/ui';

// utilities
import calculateIconSize from '@/utilities/ui/calculateIconSize';

const CircularProgressWithIcon: FC<Props> = ({ icon, iconColor, progress, progressColor }) => {
  // hooks
  const subTextColor = useSubTextColor();
  // memos
  const iconSize = useMemo(() => calculateIconSize('lg'), []);
  const trackColor = useMemo(
    () =>
      ({
        light: 'gray.300',
        dark: 'whiteAlpha.400',
      }) satisfies Theme,
    []
  );
  const value = useMemo(
    () => (progress ? (progress[1] > 0 ? (progress[0] / progress[1]) * 100 : 0) : null),
    [progress]
  );

  return (
    <ProgressCircle.Root value={value} aria-label="Progress">
      <ProgressCircle.Circle
        css={{
          '--size': '100px',
          '--thickness': '4px',
        }}
      >
        <ProgressCircle.Track
          stroke={{
            _dark: trackColor.dark,
            base: trackColor.light,
          }}
        />
        <ProgressCircle.Range stroke={progressColor || 'primary.solid'} />
      </ProgressCircle.Circle>

      <AbsoluteCenter>
        <Icon
          as={icon}
          color={
            iconColor || {
              _dark: subTextColor.dark,
              base: subTextColor.light,
            }
          }
          boxSize={iconSize}
        />
      </AbsoluteCenter>
    </ProgressCircle.Root>
  );
};

export default CircularProgressWithIcon;
