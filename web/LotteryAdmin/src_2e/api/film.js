import { request } from '../utils/http';

export function GetSaveDrawApply(data) {
    return request({ url: 'api/getsavedrawapplyinfo', method: 'post', data: data });
}
