import axios from "axios";
const api = axios.create({
    baseURL: 'http://192.168.2.38:8000',
})

async function updateFarms(user = null) {
    var new_farms = [];

    var r = null;
    if (user) {
        console.log(`/user/${user.id}/farm`);
        r = await api.get(`/user/${user.id}/farm`);
    }

    if (user && user.roles.indexOf('root') > -1) {
        console.log('/farm');
        r = await api.get('/farm');
    }

    if (r) {
        new_farms = [...new Set(r.data.data)]
        for (var i in new_farms) {
            try {
                r = await api.get(`/farm/${new_farms[i].id}/addr`);
                new_farms[i].addr = r.data.data;
            } catch (e) {
                new_farms[i].addr = null;
            }

            try {
                r = await api.get(`/user/${new_farms[i].person}`);
                new_farms[i].owner = r.data.data;
            } catch (e) {
                new_farms[i].owner = null;
            }
        }
        if (new_farms.length > 0)
            return new_farms;
    }

    return null;
}

async function updateServs(user = null) {
    var new_servs = [];

    var r = null;
    if (user) {
        console.log(`/user/${user.id}/serv`);
        r = await api.get(`/user/${user.id}/serv`);
    }

    if (user && user.roles.indexOf('root') > -1) {
        console.log(`/serv`);
        r = await api.get(`/serv`);
    }

    if (r) {
        new_servs = [...new Set(r.data.data)]

        console.log(new_servs);
        for (var i in new_servs) {
            try {
                r = await api.get(`/farm/${new_servs[i].farm}`);
                new_servs[i].farm = r.data.data;
            } catch (e) {
                new_servs[i].farm = null;
            }

            try {
                r = await api.get(`/user/${new_servs[i].employee}`);
                new_servs[i].employee = r.data.data;
            } catch (e) {
                new_servs[i].employee = null;
            }
        }

        if (new_servs.length > 0)
            return new_servs;
    }

    return null;
}

async function updateUsers(user = null, force = false) {
    var new_users = [];

    var r = null;
    if (force||(user && user.roles.indexOf('root') > -1)) {
        console.log('/user/');
        r = await api.get('/user');
    }

    if (r) {
        new_users = [...new Set(r.data.data)]
        if (new_users.length > 0)
            return new_users;
    }

    return null;
}

export {updateUsers, updateServs, updateFarms};
export default api;
