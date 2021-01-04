import React, { useState, useEffect } from 'react';
import { View,
    AsyncStorage,
    SafeAreaView,
    BackHandler,
    Alert,
    Image,
    Text,
    TextInput,
    ScrollView,
    TouchableOpacity,
    StyleSheet} from 'react-native';
import logo from '../assets/logo.png';
import { Platform } from '@unimodules/core';
import styles from '../Styles';

import UserSelList from '../components/UserSelList';
import api, {updateUsers} from '../services/api'
export default function User({navigation}) {
    const [users,  setUsers] = useState([]);
    const [user,  setUser] = useState(null);

    const [token, setToke] = useState('');

    function handleBackButtonClick() {
        console.log("backing...");
        navigation.navigate(navigation.getParam('back'));
        return true;
    }

    useEffect(() => {
        BackHandler.addEventListener('hardwareBackPress', handleBackButtonClick);
        return () => {
            BackHandler.removeEventListener('hardwareBackPress', handleBackButtonClick);
        };
    }, []);


    useEffect(()=>{
        AsyncStorage.getItem('curr_user').then(curr=>{
            if (curr) {
                curr = JSON.parse(curr);
                setToke(curr.token);

                setUser(curr);
                updateUsers(curr, true).then(r => {
                    console.log('UPDATING USERS ', r);
                    setUsers(r);
                });
            }
        });
    },[]);

    async function handleSubmit() {
        navigation.navigate('Farm', {user,
            back:'FarmPrepar'});
    }

    return (<SafeAreaView style={styles.container}>

        <Image style={{marginTop: 30}} source={logo}/>

        <UserSelList
            users={users}
            curr={user}
            onRefresh={() => {
                updateUsers(user).then(r => {
                    console.log('UPDATING USERS ', r);
                    setUsers(r)
                });
            }}
            onSelect={setUser}
        />
        <View style={styles.form}>
            <TouchableOpacity onPress={handleSubmit} style={styles.button}>
                <Text style={styles.buttonText}>Próximo</Text>
            </TouchableOpacity>
        </View>
    </SafeAreaView>);
}
