import React, { useState, useEffect } from 'react';
import { SafeAreaView,
    Text,
    View,
    AsyncStorage,
    StyleSheet,
    ScrollView,
    Image,
    Alert } from 'react-native';
import {Button} from 'react-native-paper';
import { TouchableOpacity } from 'react-native-gesture-handler';

import api, { updateServs, updateFarms, updateUsers } from '../services/api';
import ServList from '../components/ServList';
import FarmList from '../components/FarmList';
import LogoutButton from '../components/LogoutButton';

export default function List({ navigation }) {
    const [curr, setCurr] = useState(null);

    const [users, setUsers] = useState(null);
    const [servs, setServs] = useState(null);
    const [farms, setFarms] = useState(null);

    function logout() {
        AsyncStorage.clear();
        navigation.navigate('Login');
    }

    async function updateCurrUser() {
        AsyncStorage.getItem('curr_user').then(user=> {
            if (user) {
                user = JSON.parse(user);
                    setCurr(user)

                console.log('current_user [state]:', user);
                updateServs(user).then(r => {
                    console.log('UPDATING SERVICES ', r);
                    setServs(r)
                });

                updateFarms(user).then(r => {
                    console.log('UPDATING FARMS ', r);
                    setFarms(r)
                });
            }
        });
    }

    useEffect(()=>{
        updateCurrUser();
    },[]);
    return (<SafeAreaView style={styles.container}>
        <LogoutButton
            user={curr}
            action={logout}/>

        <View style={{...styles.container, ...styles.center}}>
            <FarmList
                farms={farms}
                onRefresh={() => {
                    updateFarms(curr).then(r => {
                        console.log('UPDATING FARMS ', r);
                        setFarms(r)
                    });
                }}
                onEdit={() => {}}
                onRemove={() => {}}
            />

            <ServList
                servs={servs}
                onRefresh={() => {
                    updateServs(curr).then(r => {
                        console.log('UPDATING SERVICES ', r);
                        setServs(r)
                    });
                }}
                onEdit={() => {}}
                onRemove={() => {}}
            />
        </View>

    </SafeAreaView>)
}

const styles = StyleSheet.create({
    container: {
        flex: 1,
    },
    logo: {
        height: 40,
        resizeMode: 'contain',
        marginTop: 20
    },
    row: {
        flex: 1,
        flexDirection: 'row',
        alignItems: 'stretch',
        justifyContent: 'center',
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
    button: {
        height: 32,
        backgroundColor: '#23B185',
        justifyContent: 'center',
        alignItems:'center',
        borderRadius:2,
        marginTop: 15,
        padding: 10,
    },
    buttonText:{
        color: '#FFF',
        fontWeight:'bold',
        fontSize:15,
    },
    box: {
        padding: 10,
        minWidth: 300,
        borderRadius: 2,
        borderColor: '#23B185',
        borderWidth: 1,
        marginTop: 5,
        marginBottom: 5,
    },
});
