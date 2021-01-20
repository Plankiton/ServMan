import React from 'react';
import {Text,
    AsyncStorage,
    Alert,
    Image } from 'react-native';
import { TouchableOpacity } from 'react-native-gesture-handler';
import styles from '../Styles'

import api from '../services/api';
import logo from '../assets/logo.png';
import trans from '../Translate';

export default function LogoutButton(props) {
    async function back() {
        props.navigation.navigate(props.back);
    }

    return (<TouchableOpacity
        style={styles.center}
        onPress={back}>
        <Image style={styles.logo} source={logo}/>
        <Text style={styles.title}>
            Voltar {props.back? `para ${trans[props.back]?trans[props.back]:props.back}` :''}
        </Text>
    </TouchableOpacity>);
}
