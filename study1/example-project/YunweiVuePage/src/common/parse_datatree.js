
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



// export const menutree_convert_checked  = (checked_list) => {
//
//
//
//   for (let i =0; i<checked_list.length; i ++){
//     console.log(checked_list[i].checked);
//     if(checked_list[i].checked !== "1") {
//
//       checked_list[i].checked = false;
//
//     }else {
//       alert("ho");
//       checked_list[i].checked = true;
//     }
//
//     // 判断checked_list[i].children 值为真
//     if (checked_list[i].children){
//       checked_list[i].children = menutree_convert_checked(checked_list[i].children)
//     }
//
//   }
//
//   alert(JSON.stringify(checked_list));
//   console.log(checked_list);
//
//   return checked_list;
//
//
// };


export const menutree_convert_checked = (checked_list) => {
  checked_list.map(item => {

    if (item.checked === "1") {
      alert(item.title);
      alert(item.checked);
      item.checked = true;
    } else {
      delete item['checked']
    }

    if (item.children) {
      menutree_convert_checked(item.children)
    }
  })
}