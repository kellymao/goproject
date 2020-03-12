
export const prase_tree = (checked_list,roleid) => {

  return checked_list.slice(0).filter(item => {

    if(item.children){
      item.children  = prase_tree(item.children.slice(0),roleid);

    }

    item.roleid = roleid;
    return item.checked || (typeof(item.children) !== "undefined" && item.children.length !==0 )
    //return item.checked || JSON.stringify(item.children) !== '[]' ;

  });
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

  const cpy_list = checked_list.slice(0);
  cpy_list.map(item => {

    let ck = item.checked;
    delete item['checked']


    if (ck === "1") {
      if (item.children){
        item.selectd = true;
      }else{
        item.checked = true;
      }

    }

     if (item.children) {
       item.children = menutree_convert_checked(item.children);


    }
  });

  return cpy_list;
};



