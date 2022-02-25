import { asyncRoutes, constantRoutes } from '../../router';
import { loadMenu } from '../../api/sys';
import { setToken } from '../../utils/auth';

/**
 * Use meta.role to determine if the current user has permission
 * @param roles
 * @param route
 */
function hasPermission(roles, route) {
    if (route.meta && route.meta.roles) {
        return roles.some(role => route.meta.roles.includes(role));
    } else {
        return true;
    }
}

/**
 * Filter asynchronous routing tables by recursion
 * @param routes asyncRoutes
 * @param roles
 */
export function filterAsyncRoutes(routes, roles) {
    const res = [];

    routes.forEach(route => {
        const tmp = { ...route };
        if (hasPermission(roles, tmp)) {
            if (tmp.children) {
                tmp.children = filterAsyncRoutes(tmp.children, roles);
            }
            res.push(tmp);
        }
    });

    return res;
}

const state = {
    routes: [],
    addRoutes: []
};

const mutations = {
    SET_ROUTES: (state, routes) => {
        state.addRoutes = routes;
        //state.routes = constantRoutes.concat(routes)
        state.routes = routes;
    }
};

const actions = {
    loadMenu({ commit }) {
        return new Promise((resolve, reject) => {
            loadMenu().then(response => {
                const { obj, code, msg } = response;
                if (code == 200) {
                    commit('SET_ROUTES', obj);
                    setToken(obj.token);
                    resolve();
                } else {
                    reject(msg);
                }
            }).catch(error => {
                reject(error);
            });
        });
    },
    generateRoutes({ commit }, roles) {
        return new Promise(resolve => {
            let accessedRoutes;

            accessedRoutes = asyncRoutes || [];
            // if (roles.includes('admin')) {
            //   accessedRoutes = asyncRoutes || []
            // } else {
            //   accessedRoutes = filterAsyncRoutes(asyncRoutes, roles)
            // }
            commit('SET_ROUTES', accessedRoutes);
            resolve(accessedRoutes);
        });
    }
};

export default {
    namespaced: true,
    state,
    mutations,
    actions
};
