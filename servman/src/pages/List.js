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
import UserList from '../components/UserList';
import LogoutButton from '../components/LogoutButton';

import styles from '../Styles'

export default function List({ navigation }) {
    const [curr, setCurr] = useState(null);
    const [active, setActive] = useState('user');

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

                updateUsers().then(r => {
                    console.log('UPDATING USERS ', r);
                    setUsers(r)
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

        <View style={{...styles.container, ...styles.center,
                paddingVertical: 20,
            }}>

            <View style={{...styles.line,
                marginTop: 15,
                marginBottom: 5,
            }}></View>
            {active == 'user'?(
            <UserList
                users={users}
                onRefresh={() => {
                    updateUsers().then(r => {
                        console.log('UPDATING USERS ', r);
                        setUsers(r)
                    });
                }}
                onEdit={() => {}}
                onRemove={() => {}}
            />
            ):(<Text onPress={()=>{
                setActive('user');
            }}>Usuários</Text>)}

            <View style={{...styles.line,
                marginTop: 15,
                marginBottom: 5,
            }}></View>
            {active == 'farm'?(
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
            ):(<Text onPress={()=>{
                setActive('farm');
            }}>Fazendas</Text>)}

            <View style={{...styles.line,
                marginTop: 15,
                marginBottom: 5,
            }}></View>
            {active == 'serv'?(
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
            ):(<Text onPress={()=>{
                setActive('serv');
            }}>Serviços</Text>)}

        </View>

    </SafeAreaView>)
}
