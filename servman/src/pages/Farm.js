import React, { useState, useEffect } from 'react';
import { View,
    AsyncStorage,
    KeyboardAvoidingView,
    BackHandler,
    Alert,
    Image,
    Text,
    TextInput,
    TouchableOpacity,
    StyleSheet} from 'react-native';
import logo from '../assets/logo.png';
import { Platform } from '@unimodules/core';
import styles from '../Styles';

import api from '../services/api'
export default function User({navigation}) {
    const [name, setName] = useState('');

    const [token, setToke] = useState('');

    const [farm, setFarm] = useState(null);

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
        const lfarm = navigation.getParam('farm');
        AsyncStorage.getItem('curr_user').then(curr=>{
            if (curr) {
                curr = JSON.parse(curr);
                setToke(curr.token);
            }
        });

        if(lfarm) {
            setName(lfarm.name);
            setFarm(lfarm);
        }

    },[]);

    async function handleSubmit() {
        var url = '/user';
        if (user) {
            url += `/${user.id}`;
            setPass(null);
        }

        console.log(url,' ',token);
        try {
            const response = await api.post(url, {token,data: {
                password: pass,
                document: doc,
                phone,
                name,
            }})

            navigation.navigate('List');
        } catch {
            Alert.alert(`Não foi possível ${user?'editar':'criar'} ${name}`);
        }

    }

    return (
        <KeyboardAvoidingView
         enabled={Platform.OS== 'ios'} 
         behavior="padding" style={styles.container}>
            <View style={styles.centerBox}>
                <Image source={logo}/>

                <View style={styles.form}>

                    <Text style={styles.label}>Nome Completo</Text>
                    <TextInput
                        style={styles.input}
                        placeholder="Digite seu nome completo"
                        placeholderTextColor="#999"
                        keyboardType="default"
                        autoCapitalize="words"
                        autoCorrect={false}
                        onChangeText={setName}
                    >{name}</TextInput>

                    <TouchableOpacity onPress={handleSubmit} style={styles.button}>
                        <Text style={styles.buttonText}>{farm?'Salvar':'Criar usuário'}</Text>
                    </TouchableOpacity>
                </View>

            </View>
        </KeyboardAvoidingView>
    )
}
