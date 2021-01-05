import React, { useState, useEffect } from 'react';
import { View,
    AsyncStorage,
    KeyboardAvoidingView,
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

import api from '../services/api'
export default function Serv({navigation}) {
    const [description, setDesc] = useState('');
    const [price,       setPric] = useState('');

    const [token, setToke] = useState('');

    const [serv,  setServ] = useState(null);
    const [user,  setUser] = useState(null);
    const [farm,  setFarm] = useState(null);

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
        const lserv = navigation.getParam('serv');
        const luser = navigation.getParam('user');
        const lfarm = navigation.getParam('farm');
        AsyncStorage.getItem('curr_user').then(curr=>{
            if (curr) {
                curr = JSON.parse(curr);
                setToke(curr.token);
            }
        });

        if (luser && lfarm) {
            setUser(luser);
            setFarm(lfarm);
        } else if (lserv) {
            setDesc(lserv.description);

            setServ(lserv);
        } else {
            navigation.navigate('List');
        }

    },[]);

    async function handleSubmit() {
        var url = '';
        if (user) {
            url += `/user/${user.id}`
        }

        url += '/serv';
        if (serv) {
            url += `/${serv.id}`;
        }

        console.log(url,' ',token);
        try {
            var text = price.replace('-', '');
            text = text.replace(' ', '');
            const response = await api.post(url, {token, data: {
                description,
                farm,
                price,
            }})

            navigation.navigate('List');
        } catch (e) {
            console.log(e);
            Alert.alert(`Não foi possível ${serv?'editar':'criar'} ${description}`);
        }

    }

    return (
        <KeyboardAvoidingView
         enabled={Platform.OS== 'ios'} 
         behavior="padding" style={styles.container}>
            <View style={styles.centerBox}>
                <Image source={logo}/>

                <ScrollView style={styles.form}>
                    <Text style={styles.label}>Descrição</Text>
                    <TextInput
                        multiline = {true}
                        numberOfLines = {4}
                        style={{...styles.input,
                            height: 100,
                        }}
                        placeholder="Digite a descrição do serviço"
                        placeholderTextColor="#999"
                        keyboardType="default"
                        autoCapitalize="words"
                        autoCorrect={false}
                        onChangeText={setDesc}
                    >{description}</TextInput>
                </ScrollView>

                <ScrollView style={styles.form}>
                    <Text style={styles.label}>Preço</Text>
                    <TextInput
                        style={styles.input}
                        placeholder="Digite o preço (por hora) do serviço"
                        placeholderTextColor="#999"
                        keyboardType="numeric"
                        autoCapitalize="words"
                        autoCorrect={false}
                        onChangeText={setPric}
                    >{price}</TextInput>
                </ScrollView>

                <View style={styles.form}>
                    <TouchableOpacity onPress={handleSubmit} style={styles.button}>
                        <Text style={styles.buttonText}>{serv?'Salvar':'Criar fazenda'}</Text>
                    </TouchableOpacity>
                </View>
            </View>
        </KeyboardAvoidingView>
    )
}
