import React from 'react';
import {Text,
        StyleSheet,
        Image } from 'react-native';
import { TouchableOpacity } from 'react-native-gesture-handler';

import logo from '../assets/logo.png';

function LogoutButton(props) {
    return (<TouchableOpacity
        style={styles.center}
        onPress={props.action}>
        <Image style={styles.logo} source={logo}/>
        <Text style={styles.title}>
            Sair {props.user? ` da conta do ${props.user.name}` :''}
        </Text>
    </TouchableOpacity>);
}

const styles = StyleSheet.create({
    logo: {
        height: 40,
        resizeMode: 'contain',
        marginTop: 20
    },
    center: {
        alignItems:'center',
        justifyContent:'center',
    },
    title: {
        color: '#23B185',
        fontWeight: 'bold',
        fontSize: 16,
        marginBottom: 30
    },
});

export default LogoutButton;
