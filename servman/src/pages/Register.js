import React, { useState, useEffect } from 'react';
import { View, AsyncStorage, KeyboardAvoidingView ,Image, Text, TextInput, TouchableOpacity, StyleSheet} from 'react-native';
import logo from '../assets/logo.png';
import { Platform } from '@unimodules/core';

import api from '../services/api'
export default function Register({navigation}) {
    const [email, setEmail] = useState('');
    const [techs, setTechs] = useState('');

    /*
    useEffect(()=>{
        AsyncStorage.getItem('user').then(user=>{
            if(user) {
                navigation.navigate('List')
            }
        }
        )
    },[]);
    */

    async function handleSubmit() {
        navigation.navigate('List');
        /*
        const response = await api.post('/sessions', {
            email
        })
        const {_id} = response.data;
        await AsyncStorage.setItem('user', _id);
        await AsyncStorage.setItem('techs', techs);

        navigation.navigate('List');
        */
    }

    async function goToRegister() {
        navigation.navigate('Register');
    }

    return (
        <KeyboardAvoidingView
         enabled={Platform.OS== 'ios'} 
         behavior="padding" style={styles.container}>
            <View style={styles.box}>
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
                        onChangeText={setEmail}
                    >{email}</TextInput>
                    <Text style={styles.label}>Número de Telefone</Text>
                    <TextInput
                        style={styles.input}
                        placeholder="Digite seu telefone"
                        placeholderTextColor="#999"
                        keyboardType="default"
                        autoCapitalize="words"
                        autoCorrect={false}
                        onChangeText={setEmail}
                    >{email}</TextInput>
                    <Text style={styles.label}>CPF</Text>
                    <TextInput
                        style={styles.input}
                        placeholder="Digite seu CPF"
                        placeholderTextColor="#999"
                        keyboardType="email-address"
                        autoCapitalize="none"
                        autoCorrect={false}
                        value={email}
                        onChangeText={setEmail}
                    >
                    </TextInput>
                    <Text style={styles.label}>SENHA</Text>
                    <TextInput 
                        style={styles.input}
                        placeholder="Digite sua senha"
                        placeholderTextColor="#999"
                        keyboardType={"default"}
                        secureTextEntry={true}
                        onChangeText={setTechs}
                    >{techs}</TextInput>

                    <TouchableOpacity onPress={handleSubmit} style={styles.button}>
                        <Text style={styles.buttonText}>Entrar</Text>
                    </TouchableOpacity>
                </View>

                <Text style={styles.linkText} onPress={goToRegister}>Não tem conta? Registre-se</Text>
            </View>
        </KeyboardAvoidingView>
    )
}
const styles = StyleSheet.create({
   container: {
       flex:1,
       justifyContent:'center',
       alignItems:'center'
   }, 
   box:{
        alignSelf: 'stretch',
        paddingHorizontal: 30,
        alignItems: 'center',
        justifyContent: 'space-between',
        marginTop: 30,
   },
   form:{
        alignSelf: 'stretch',
        paddingHorizontal: 30,
        marginTop: 30,
        marginBottom: 30,
   },
   label: {
       fontWeight: 'bold',
       color:'#444',
       marginBottom:8
   },
   input: {
       borderWidth:1,
       borderColor: '#ddd',
       paddingHorizontal:20,
       fontSize: 16,
       color:'#444',
       height: 44,
       marginBottom: 20,
       borderRadius: 2
   },
   button: {
       height: 42,
       backgroundColor: '#23B185',
       justifyContent: 'center',
       alignItems:'center',
       borderRadius:2,
   },
   buttonText:{
       color: '#FFF',
       fontWeight:'bold',
       fontSize:16,
   },
   linkText: {
       color: '#23B185',
       fontWeight:'bold',
       fontSize:16,
   },
});
