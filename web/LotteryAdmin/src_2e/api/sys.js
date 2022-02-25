import { request } from '../utils/http';

export function loadMenu(data) {
    return request({ url: 'sys/getmenulist', method: 'post', data: data });
}