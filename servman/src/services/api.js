import axios from "axios";
const api = axios.create({
    baseURL: 'http://5f0a13c7e2f4.ngrok.io',
})

async function updateFarms(user = null) {
    var new_farms = [];

    var r = null;
    if (user) {
        console.log(`/user/${user.id}/farm`);
        r = await api.get(`/user/${user.id}/farm`);
    } else if (user && user.roles.indexOf('root')) {
        console.log(`/farm`);
        r = await api.get(`/farm`);
    }

    if (r) {
        new_farms = [...new Set(r.data.data)]
        if (new_farms.length > 0)
            return new_farms;
    }

    return null;
}

async function updateServs(user = null, farm = null) {
    var new_servs = [];

    var r = null;
    if (user) {
        console.log(`/user/${user.id}/serv`);
        r = await api.get(`/user/${user.id}/serv`);
    } else if (farm) {
        console.log(`/user/farm/${farm.id}/serv`);
        r = await api.get(`/user/farm/${farm.id}/serv`);
    } else if (user && user.roles.indexOf('root')) {
        console.log(`/serv`);
        r = await api.get(`/serv`);
    }

    if (r) {
        new_servs = [...new Set(r.data.data)]
        if (new_servs.length > 0)
            return new_servs;
    }

    return null;
}

async function updateUsers(user_id = null) {
    var new_users = [];

    console.log(`/user/`);
    const r = await api.get('/user'+(user_id?
        '/'+user_id
        :''));

    if (r) {
        if (user_id) {
            return r.data.data;
        }

        new_users = [...new Set(r.data.data)]
        if (new_users.length > 0)
            return new_users;
    }

    return null;
}

export {updateUsers, updateServs, updateFarms};
export default api;
