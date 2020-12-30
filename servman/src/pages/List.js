import React, { useState, useEffect } from 'react';
import { SafeAreaView,
    Text,
    View,
    AsyncStorage,
    StyleSheet,
    BackHandler,
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

    function handleBackButtonClick() {
        console.log("backing...");
        return true;
    }

    useEffect(() => {
        BackHandler.addEventListener('hardwareBackPress', handleBackButtonClick);
        return () => {
            BackHandler.removeEventListener('hardwareBackPress', handleBackButtonClick);
        };
    }, []);

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

        var user = await AsyncStorage.getItem('curr_user');

        if (user) {
            user = JSON.parse(user);
            setCurr(user);

            try {
                const r = await api.post(`/user/${user.id}`)
                updated_user = {
                    ...r.data.data,
                    token: token,
                }

                await AsyncStorage.setItem('curr_user',
                    JSON.stringify(updated_user));

                updateScreen(
                    user.roles.indexOf('root')>-1?
                    'user':'serv',
                    updated_user);
            } catch (e) {
                updateScreen(
                    user.roles.indexOf('root')>-1?
                    'user':'serv',
                    user);
            }
        }

    }

    function onRemove(obj, type = null) {
        var trans = {user:'usuário',serv:'serviço',farm:'fazenda'};
        Alert.alert(
            'Aviso',
            `Quer mesmo apagar ${obj.name?'"'+obj.name+'"':'ess'+ (type=='farm'?'a':'e') +' '+trans[type]}?`,
            [
                {
                    text: 'Cancelar',
                    onPress: () => null,
                    style: 'cancel',
                },
                {text: 'Apagar', onPress: () => {
                    console.log("\n\n\n DELETING ", obj, "\n\n\n");
                    api.delete(`/${type}/${obj.id}`).then(r=> {;
                        updateScreen(type);
                    }).catch(()=>{
                        Alert.alert('Não foi possível apagar '+ obj.name?'"'+obj.name+'"':'ess'+ (type=='farm'?'a':'e') +' '+trans[type]);
                    });
                }},
            ]
        );
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
                onCreate={() => {
                    navigation.navigate('User', {back:'List'});
                }}
                onEdit={(user) => {
                    navigation.navigate('User', {user,
                        back:'List'});
                }}
                onRemove={(user) => onRemove(user, 'user')}
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
                onRemove={(farm) => onRemove(farm, 'farm')}
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
                onRemove={(serv) => onRemove(serv, 'serv')}
            />
            ):(<Text style={styles.title} onPress={()=>{
                updateScreen('serv');
            }}>Serviços</Text>)}

        </View>

    </SafeAreaView>)
}
