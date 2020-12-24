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

    const [users, setUsers] = useState([]);
    const [servs, setServs] = useState([]);
    const [farms, setFarms] = useState([]);

    async function updateScreen(screen, user = null) {
        setActive(screen);

        if (screen == 'serv') {
            setFarms([]);
            setUsers([]);
            updateServs(curr?curr:user).then(r => {
                console.log('UPDATING SERVICES ', r);
                setServs(r)
            });
        }

        else if (screen == 'farm') {
            setServs([]);
            setUsers([]);
            updateFarms(curr?curr:user).then(r => {
                console.log('UPDATING FARMS ', r);
                setFarms(r)
            });
        }

        else if (screen == 'user') {
            setFarms([]);
            setServs([]);
            updateUsers(curr?curr:user).then(r => {
                console.log('UPDATING USERS ', r);
                setUsers(r);
            });
        }
    }

    async function updateCurrUser() {
        AsyncStorage.getItem('curr_user').then(user=> {
            if (user) {
                user = JSON.parse(user);
                setCurr(user);

                updateScreen(user.roles.indexOf('root')>-1?
                    'user':'serv', user);
            }
        });
    }

    useEffect(()=>{
        updateCurrUser();
    },[]);
    return (<SafeAreaView style={styles.container}>
        <LogoutButton
            navigation={navigation}
            user={curr}/>

        <View style={{...styles.container, ...styles.center,
                paddingVertical: 20,
            }}>

            {curr && curr.roles.indexOf('root')>-1?(<><View style={{...styles.line,
                marginTop: 15,
                marginBottom: 5,
            }}></View>
            {active == 'user'?(
            <UserList
                users={users}
                onRefresh={() => {
                    updateUsers(curr).then(r => {
                        console.log('UPDATING USERS ', r);
                        setUsers(r)
                    });
                }}
                onEdit={() => {}}
                onRemove={() => {}}
            />
            ):(<Text style={styles.title} onPress={()=>{
                updateScreen('user');
            }}>Usuários</Text>)}</>):null}

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
            ):(<Text style={styles.title} onPress={()=>{
                updateScreen('farm');
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
            ):(<Text style={styles.title} onPress={()=>{
                updateScreen('serv');
            }}>Serviços</Text>)}

        </View>

    </SafeAreaView>)
}
