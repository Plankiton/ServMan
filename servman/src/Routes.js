import { createAppContainer, createSwitchNavigator } from 'react-navigation';

import Login from './pages/Login';
import List from './pages/List';

import User from './pages/User';
import Farm from './pages/Farm';
import Serv from './pages/Serv';

import SelUser from './pages/SelUser';
import SelFarm from './pages/SelFarm';

const Routes = createAppContainer(
    createSwitchNavigator({
        Login,
        List,

        User,
        Farm,
        Serv,

        SelUser,
        SelFarm,
    })
);

export default Routes;
