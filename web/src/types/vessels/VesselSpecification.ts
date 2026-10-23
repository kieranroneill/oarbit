// types
import type VesselDimensions from './VesselDimensions';
import type Engine from './Engine';
import type FuelTanks from './FuelTanks';
import type Rigging from './Rigging';
import type SailArea from './SailArea';

interface VesselSpecification {
  builder: string | null;
  construction: string | null;
  designer: string | null;
  dimensions: VesselDimensions;
  engine: Engine | null;
  firstBuilt: number | null;
  hullType: string | null;
  lastBuilt: number | null;
  make: string;
  model: string;
  rigging: Rigging | null;
  sailArea: SailArea | null;
  tanks: FuelTanks | null;
}

export default VesselSpecification;
