import axios from 'axios';
import qs from 'qs';
import router from '../../../router';

export function request(config = {}) {
    let service = axios.create({
        baseURL: process.env.VUE_APP_API
    });
    //拦截器
    service.interceptors.request.use(function(config) {
        if (config.method === 'post') {
            config.data = qs.stringify(config.data);
        } else if (config.method === 'file') {
            config.method = 'post';
            if (config.data && config.data instanceof Object) {
                let params = new FormData();
                for (let key in config.data) {
                    params.append(key, config.data[key]);
                }
                config.data = params;
            }
        }
        return config;
    }, function(error) {
        return Promise.reject(error);
    });
    //响应之后
    service.interceptors.response.use(function(response) {
        console.log(response.data.code == 401);
        if (response.data.code == 401) {
            router.push('/login');
        }
        return response;
    }, function(error) {
        return Promise.reject(error);
    });
    let token = localStorage.getItem('token');
    config.headers = {
        'token': token
    };
    console.log('AAAAAAAAAAAAA');

    return service.request(config).then(res => res.data);
}