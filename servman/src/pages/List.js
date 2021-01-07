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

import Moment from 'moment';
import api, { updateServs, updateFarms, updateUsers } from '../services/api';
import ServList from '../components/ServList';
import FarmList from '../components/FarmList';
import UserList from '../components/UserList';
import LogoutButton from '../components/LogoutButton';
import Footer from '../components/Footer';
import trans from '../Translate';

import styles from '../Styles';

export default function List({ navigation }) {
    const [curr, setCurr] = useState(null);
    const [active, setActive] = useState(false);

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

            if (!active) {
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

    }

    function onRemove(obj, type = null) {
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
    return (<SafeAreaView style={{...styles.container, ...styles.root}}>
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
                        curr={curr}
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
                        onDetail={(user) => {
                            var items = [{title: user.name}]
                            var subitems = [];
                            for (var i in user) {
                                if (['farm',
                                    'employee',
                                    'created_at',
                                    'updated_at',
                                    'roles',
                                    'price',
                                    'person',
                                    'address',
                                    'id'].indexOf(i)>=0)continue;

                                if (user[i] && typeof user[i] == 'object') {
                                    subitems.push({...user[i], type: i})
                                } else {
                                    items.push({
                                        key: i,
                                        value: user[i]
                                    });
                                }

                            }

                            items.push({
                                key: 'created_at',
                                value: user.created_at,
                            });
                            items.push({
                                key: 'updated_at',
                                value: user.updated_at,
                            });

                            for (var i in subitems) {
                                console.log('!', i, subitems[i])
                                if (subitems[i].length <= 1) continue;

                                items.push({title: trans[subitems[i].type]});
                                for (var c in subitems[i]) {
                                    if (['user',
                                        'created_at',
                                        'updated_at',
                                        'type',
                                        'person',
                                        'address',
                                        'id'].indexOf(c)>=0)continue;

                                    items.push({
                                        parent: i,
                                        key: c,
                                        value: subitems[i][c],
                                    });
                                }
                            }

                            navigation.navigate('Detail', {items, back:'List'});
                        }}
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
                    onCreate={() => {
                        navigation.navigate('SelUser', {
                            title: 'dono da fazenda',
                            back:'List', dest:'Farm'});
                    }}
                    onEdit={(farm) => {
                        navigation.navigate('Farm', {farm,
                            back:'List'});
                    }}
                    onRemove={(farm) => onRemove(farm, 'farm')}
                    onDetail={(farm) => {
                        var items = [{title: farm.name}]
                        var subitems = [];
                        for (var i in farm) {
                            if (['farm', 'created_at', 'type', 'updated_at', 'person', 'address', 'id'].indexOf(i)>=0)continue;

                            if (farm[i] && typeof farm[i] == 'object') {
                                subitems.push({...farm[i], type: i})
                            } else {
                                items.push({
                                    key: i,
                                    value: farm[i]
                                });
                            }
                        }

                        items.push({
                            key: 'created_at',
                            value: farm.created_at,
                        });
                        items.push({
                            key: 'updated_at',
                            value: farm.updated_at,
                        });


                        for (var i in subitems) {
                            items.push({title: trans[subitems[i].type]});
                            for (var c in subitems[i]) {
                                if (['farm', 'created_at', 'type', 'updated_at', 'person', 'address', 'id'].indexOf(c)>=0)continue;

                                items.push({
                                    parent: i,
                                    key: c,
                                    value: subitems[i][c],
                                });
                            }
                        }
                        navigation.navigate('Detail', {items, back:'List'});
                    }}
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
                    onCreate={() => {
                        navigation.navigate('SelUser',
                            {
                                title: 'funcionário',
                                back: 'List',
                                dest: 'SelFarm',
                                next: 'Serv'});
                    }}
                    onEdit={(serv) => {
                        navigation.navigate('Serv', {serv,
                            back:'List'});
                    }}
                    onRemove={(serv) => onRemove(serv, 'serv')}

                    onDetail={(serv) => {
                        var items = [{title: serv.description}]
                        var subitems = [];
                        for (var i in serv) {
                            if ([
                                'person',
                                'started_at',
                                'finished_at',
                                'price',
                                'address',
                                'id'].indexOf(i)>=0)continue;

                            if (serv[i] && typeof serv[i] == 'object') {
                                subitems.push({...serv[i], type: i})
                            } else {
                                items.push({
                                    key: i,
                                    value: serv[i]
                                });
                            }

                        }

                        Moment.locale('pt-BR');
                        var begin = Moment(serv.started_at);
                        var end = Moment(serv.finished_at);
                        var diff = Math.abs(
                            end - begin
                        );
                        var hours = diff/1000/60/60; // converting milisec to hours

                        if (serv.price) {
                            items.push({
                                key: 'price',
                                value: `${(Number(serv.price)*hours).toFixed(2).replace('.',',')} (${serv.price.toFixed(2).replace('.',',')+'/hora'})`,
                            });
                        }

                        if (diff > 0) {
                            items.push({
                                key: 'started_at',
                                value: serv.started_at,
                            });
                            items.push({
                                key: 'finished_at',
                                value: serv.finished_at,
                            });
                        }

                        for (var i in subitems) {
                            items.push({title: trans[subitems[i].type]});
                            for (var c in subitems[i]) {
                                if (['serv',
                                    'person',
                                    'created_at',
                                    'updated_at',
                                    'address',
                                    'type',
                                    'id'].indexOf(c)>=0)continue;

                                items.push({
                                    parent: i,
                                    key: c,
                                    value: subitems[i][c],
                                });
                            }
                        }

                        navigation.navigate('Detail', {items, back:'List'});
                    }}
                />
            ):(<Text style={styles.title} onPress={()=>{
                updateScreen('serv');
            }}>Serviços</Text>)}

        </View>
        <Footer/>
    </SafeAreaView>)
}
