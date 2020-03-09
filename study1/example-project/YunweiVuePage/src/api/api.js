import axios from 'axios';

import service from '@/common/request'


let base = '';

export const requestLogin = params => { return axios.post(`${base}/login`, params).then(res => res.data); };

export const getUserList = params => { return axios.get(`${base}/user/list`, { params: params }); };

export const getUserListPage = params => { return axios.get(`${base}/user/listpage`, { params: params }); };

export const removeUser = params => { return axios.get(`${base}/user/remove`, { params: params }); };

export const editUser = params => { return axios.get(`${base}/user/edit`, { params: params }); };

export const addUser = params => { return axios.get(`${base}/user/add`, { params: params }); };


export const login = params => {

    return service({
        url: `/base/login`,
        method: 'post',
        data: params
    })
    };


export const getmenu = params =>{

  return service({
    url: `/menu/getmenu`,
    method: 'post',
    data: params
  })


};


export const getroledata = (params) => {

    alert(params);
    return service({
        url: "/role/getrolelist",
        method: 'post',
        data: params
    })
};

export const roleadd = (params) =>{

    return service({
        url: "/role/roleadd",
        method: 'post',
        data: params

    })
};

export const roledel = (params) =>{

    return service({
        url: "/role/roledel",
        method: 'post',
        data: params

    })
};

export const get_role_menutree = (url) =>{

    return service({
        url: url,
        method: 'post',

    })
};
