from uuid import UUID

from fastapi import APIRouter, Depends

from api import controllers, dependencies, enums, schemas
from api.schemas.boat import BoatResponseSchema, BoatSpecificationSchema
from api.schemas.defaults import BaseListResponseSchema

router = APIRouter(
    prefix=f"{enums.api.RouteEnum.BASE}{enums.api.RouteEnum.BOATS}",
    tags=["boat"],
)


@router.get("/{profile_id}", response_model=BoatResponseSchema)
async def boat_by_profile_id(profile_id: UUID, database=Depends(dependencies.database)):
    controller = controllers.BoatController()

    return await controller.boat_by_profile_id(database=database, profile_id=profile_id)


@router.get(f"{enums.api.RouteEnum.SPECIFICATIONS}", response_model=BaseListResponseSchema[BoatSpecificationSchema])
async def boat_specifications(database=Depends(dependencies.database)):
    controller = controllers.BoatController()

    return await controller.boat_specifications(database=database)
