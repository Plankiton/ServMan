import React from 'react';
import {Text,
    AsyncStorage,
    Alert,
    Image } from 'react-native';
import { TouchableOpacity } from 'react-native-gesture-handler';
import styles from '../Styles'

import api from '../services/api';
import logo from '../assets/logo.png';

export default function LogoutButton(props) {
    async function logout() {
        try {

            var user = await AsyncStorage.getItem('curr_user');
            console.log(user);

            if (user) {
                user = JSON.parse(user);
            }
            console.log(user.token);

            const r = await api.post('/auth/logout', {
                token: user.token,
            });
            console.log(r.data);

            await AsyncStorage.clear();

            props.navigation.navigate('Login');

        } catch (e) {
            console.log(e);
            Alert.alert(`Não foi possível deslogar!`);
        }
    }

    return (<TouchableOpacity
        style={styles.center}
        onPress={logout}>
        <Image style={styles.logo} source={logo}/>
        <Text style={styles.title}>
            Sair {props.user? ` da conta do ${props.user.name}` :''}
        </Text>
    </TouchableOpacity>);
}
