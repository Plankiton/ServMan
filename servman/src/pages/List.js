import React, { useState, useEffect } from 'react';
import { SafeAreaView,
    Text,
    View,
    ScrollView,
    AsyncStorage,
    StyleSheet,
    Image,
    Alert } from 'react-native';
import {Button} from 'react-native-paper';
import { TouchableOpacity } from 'react-native-gesture-handler';

import logo from '../assets/logo.png';


import {ServList} from '../components/Lists';
import api from '../services/api';

export default function List({ navigation }) {
    const [curr, setCurr] = useState(null);

    const [servs, setServs] = useState(null);
    const [users, setUsers] = useState(null);
    const [farms, setFarms] = useState(null);

    function logout() {
        AsyncStorage.clear();
        navigation.navigate('Login');
    }

    async function updateServs(user = null, farm = null) {
        var new_servs = servs;

        var r = null;
        if (user) {
            console.log(`/user/${user.id}/serv`);
            r = await api.get(`/user/${user.id}/serv`);
        } else if (farm) {
            console.log(`/user/farm/${farm.id}/serv`);
            r = await api.get(`/user/farm/${farm.id}/serv`);
        } else if (user && user.roles.indexOf('root')) {
            console.log(`/serv`);
            r = await api.get(`/serv`);
        }

        if (r) {
            new_servs = [...new Set(r.data.data)]
            console.log('UPDATING SERVICES ', new_servs)
            if (new_servs.lenght > 0) {
                setServs(new_servs);
            } else {
                setServs(null);
            }
        }

        return new_servs;
    }

    async function updateCurrUser() {
        AsyncStorage.getItem('curr_user').then(user=> {
            setCurr(JSON.parse(user))

            console.log('current_user [state]:', curr);
            updateServs(user = curr);
        });
    }

    useEffect(()=>{
        updateCurrUser();
    },[]);
    return (<SafeAreaView style={styles.container}>
        <TouchableOpacity
            style={styles.center}
            onPress={()=>logout()}>
            <Image style={styles.logo} source={logo}/>
            <Text style={styles.title}>
                Sair {curr? ` da conta do ${curr.name}` :''}
            </Text>
        </TouchableOpacity>
        <View style={{
            alignItems: 'center',
            ...styles.box}}>
            <Button
                onPress={() => {
                    updateServs(user = curr);
                }}
                icon={({ size, color }) => (
                    <Image
                        source={require("../assets/refresh.png")}
                        style={{
                            width: size,
                            height: size,
                            tintColor: '#23B185',
                            padding: 10,
                        }}
                    />)}>
                <Text style={{
                    ...styles.title,
                    fontSize: 20
                }}>Serviços</Text>
            </Button>
        </View>
        <ServList servs={servs}/>
    </SafeAreaView>)
}
const styles = StyleSheet.create({
    container: {
        flex: 1,
        alignItems: 'center',
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
        padding: 15,
        minWidth: 300,
        borderRadius: 2,
        borderColor: '#23B185',
        borderWidth: 1,
        marginTop: 15,
    },
});
