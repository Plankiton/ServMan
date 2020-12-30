import { createAppContainer, createSwitchNavigator } from 'react-navigation';

import Login from './pages/Login';
import List from './pages/List';

import User from './pages/User';
// import Farm from './pages/Farm';
// import Serv from './pages/Serv';

const Routes = createAppContainer(
    createSwitchNavigator({
        Login,
        List,
        User,
    })
);

export default Routes;
