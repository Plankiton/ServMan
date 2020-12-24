import { createAppContainer, createSwitchNavigator } from 'react-navigation';

import Login from './pages/Login';
import Register from './pages/Register';
import List from './pages/List';

const Routes = createAppContainer(
    createSwitchNavigator({
        Login,
        List,
        Register,
    })
);

export default Routes;
