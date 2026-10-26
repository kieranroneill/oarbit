interface ListResponse<Item> {
  items: Item[];
  limit: number;
  offset: number;
  total: number;
}

export default ListResponse;
