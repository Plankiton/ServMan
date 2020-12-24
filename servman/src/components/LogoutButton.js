import React from 'react';
import {Text,
        Image } from 'react-native';
import { TouchableOpacity } from 'react-native-gesture-handler';
import styles from '../Styles'

import logo from '../assets/logo.png';

export default function LogoutButton(props) {
    return (<TouchableOpacity
        style={styles.center}
        onPress={props.action}>
        <Image style={styles.logo} source={logo}/>
        <Text style={styles.title}>
            Sair {props.user? ` da conta do ${props.user.name}` :''}
        </Text>
    </TouchableOpacity>);
}
