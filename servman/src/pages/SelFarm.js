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

import FarmSelList from '../components/FarmSelList';
import api, {updateFarms} from '../services/api'

export default function SelFarm({navigation}) {
    const [farms,  setFarms] = useState([]);
    const [farm,  setFarm] = useState(null);

    const [token, setToke] = useState('');

    function handleBackButtonClick() {
        console.log("backing...");
        try {
            navigation.navigate(navigation.getParam('back'));
        } catch {
            navigation.navigate('List');
        }
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

                updateFarms(curr).then(r => {
                    console.log('UPDATING USERS ', r);
                    setFarms(r);
                });
            }
        });
    },[]);

    return (<SafeAreaView style={styles.container}>

        <Image style={{marginTop: 30}} source={logo}/>

        <FarmSelList
            farms={farms}
            curr={farm}
            onRefresh={() => {
                updateFarms(farm).then(r => {
                    console.log('UPDATING USERS ', r);
                    setFarms(r)
                });
            }}
            onSelect={(farm) => {
                setFarm(farm);

                const user = navigation.getParam('user');
                navigation.navigate(navigation.getParam('dest'),
                    {farm, user, back:'List'});
            }}
        />
    </SafeAreaView>);
}
