import { request } from '../utils/http';

export function login(data) {
    return request({ url: 'loginreg/dologin', method: 'post', data: data });
}

export function logout(data) {
    return request({ url: 'loginreg/dologout', method: 'post', data: data });
}