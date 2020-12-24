import React, { useState, useEffect } from 'react';
import { View,
         AsyncStorage,
         KeyboardAvoidingView,
         Alert,
         Image,
         Text,
         TextInput,
         TouchableOpacity,
         StyleSheet } from 'react-native';
import logo from '../assets/logo.png';
import { Platform } from '@unimodules/core';

import axios from "axios";
import api from '../services/api'

import styles from '../Styles';

export default function Login({navigation}) {
    const [doc, setDoc] = useState('');
    const [pass, setPass] = useState('');

    useEffect(()=>{
        AsyncStorage.getItem('curr_user').then(user=>{
            if (user) {
                navigation.navigate('List');
            }
        });
    },[]);

    async function handleSubmit() {
        try {
            const response = await api.post('/auth/login', {
                data: {
                    document: doc,
                    password: pass,
                }
            })

            const {person_id, token} = response.data.data;

            if (person_id && token) {
                const r = await api.post(`/user/${person_id}`)
                await AsyncStorage.setItem('curr_user', JSON.stringify({
                    ...r.data.data,
                    token: token,
                }));
                AsyncStorage.getItem('curr_user').then(user=> {
                    navigation.navigate('List');
                });
            } else {
                Alert.alert(`CPF ou Senha estão errados!!`);
            }
        } catch (e){
            Alert.alert(`CPF ou Senha estão errados!!`);
        }

        // setPass('');
    }

    return (
        <KeyboardAvoidingView
         enabled={Platform.OS=='ios'}
         behavior="padding" style={styles.container}>
            <View style={styles.login}>
                <Image source={logo}/>

                <View style={styles.form}>
                    <Text style={styles.label}>CPF</Text>
                    <TextInput 
                        style={styles.input}
                        placeholder="Digite seu CPF"
                        placeholderTextColor="#999"
                        keyboardType="default"
                        autoCapitalize="none"
                        autoCorrect={false}
                        value={doc}
                        onChangeText={setDoc}
                    >
                    </TextInput>
                    <Text style={styles.label}>SENHA</Text>
                    <TextInput 
                        style={styles.input}
                        placeholder="Digite sua senha"
                        placeholderTextColor="#999"
                        keyboardType={"default"}
                        secureTextEntry={true}
                        onChangeText={setPass}
                    >{pass}</TextInput>

                    <TouchableOpacity onPress={handleSubmit} style={styles.button}>
                        <Text style={styles.buttonText}>Entrar</Text>
                    </TouchableOpacity>
                </View>
            </View>
        </KeyboardAvoidingView>
    )
}
