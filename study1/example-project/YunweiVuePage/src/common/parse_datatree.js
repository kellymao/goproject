
export const prase_tree = (checked_list) => {

  console.log(checked_list);

  if (checked_list.length === 0) {
    return []
  }

  var rel = [];
  for (var i =0; i<checked_list.length; i++){
    rel.push(checked_list[i].id)
  }

  return rel
};
