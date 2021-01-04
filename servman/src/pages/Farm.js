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
export default function User({navigation}) {
    const [neighborhood, setNei] = useState('');
    const [name,   setNam] = useState('');
    const [street, setStr] = useState('');
    const [state,  setSta] = useState('');
    const [number, setNum] = useState('');
    const [cep,    setCep] = useState('');
    const [city,   setCit] = useState('');

    const [token, setToke] = useState('');

    const [farm,  setFarm] = useState(null);
    const [user,  setUser] = useState(null);

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
        const luser = navigation.getParam('user');
        AsyncStorage.getItem('curr_user').then(curr=>{
            if (curr) {
                curr = JSON.parse(curr);
                setToke(curr.token);
            }
        });

        if (luser) {
            setUser(luser);
        } else if (lfarm) {
            setNam(lfarm.name);
            setStr(lfarm.addr.street);
            setSta(lfarm.addr.state);
            setNum(lfarm.addr.number);
            setCep(lfarm.addr.cep);
            setCit(lfarm.addr.city);
            setNei(lfarm.addr.neighborhood);

            setFarm(lfarm);
        }

    },[]);

    async function handleSubmit() {
        var url = '';
        if (user) {
            url += `/user/${user.id}`
        }

        url += '/farm';
        if (farm) {
            url += `/${farm.id}`;
        }

        console.log(url,' ',token);
        try {
            var text = cep.replace('-', '');
            text = text.replace(' ', '');
            const response = await api.post(url, {token,data: {
                neighborhood,
                cep: text,
                street,
                number,
                state,
                name,
                city,
            }})

            navigation.navigate('List');
        } catch {
            Alert.alert(`Não foi possível ${farm?'editar':'criar'} ${name}`);
        }

    }

    return (
        <KeyboardAvoidingView
         enabled={Platform.OS== 'ios'} 
         behavior="padding" style={styles.container}>
            <View style={styles.centerBox}>
                <Image source={logo}/>

                <ScrollView style={styles.form}>

                    <Text style={styles.label}>Nome Completo</Text>
                    <TextInput
                        style={styles.input}
                        placeholder="Digite seu nome completo"
                        placeholderTextColor="#999"
                        keyboardType="default"
                        autoCapitalize="words"
                        autoCorrect={false}
                        onChangeText={setNam}
                    >{name}</TextInput>

                    <Text style={styles.label}>Codigo de Endereço Postal</Text>
                    <TextInput
                        style={styles.input}
                        placeholder="Digite seu cep"
                        placeholderTextColor="#999"
                        keyboardType="default"
                        autoCapitalize="words"
                        autoCorrect={false}
                        onChangeText={async (text) => {
                            text = text.replace('-', '');
                            text = text.replace(' ', '');

                            try {
                                const r = await api.get(`/addr/${text}`);
                                const addr = r.data.data;

                                if (addr) {
                                    setStr(addr.street);
                                    setCit(addr.city);
                                    setSta(addr.state);
                                    setNei(addr.neighborhood);
                                    setCep(text);
                                }
                            } catch {}

                        }}
                    >{cep.slice(0, 5)} - {cep.slice(5, 9)}</TextInput>


                    <Text style={styles.label}>Rua</Text>
                    <TextInput
                        style={styles.input}
                        placeholder="Digite sua rua"
                        placeholderTextColor="#999"
                        keyboardType="default"
                        autoCapitalize="words"
                        autoCorrect={false}
                        onChangeText={setStr}
                    >{street}</TextInput>

                    <Text style={styles.label}>Número</Text>
                    <TextInput
                        style={styles.input}
                        placeholder="Digite o número da residência"
                        placeholderTextColor="#999"
                        keyboardType="default"
                        autoCapitalize="words"
                        autoCorrect={false}
                        onChangeText={setNum}
                    >{number}</TextInput>

                    <Text style={styles.label}>Bairro</Text>
                    <TextInput
                        style={styles.input}
                        placeholder="Digite o nome do bairro"
                        placeholderTextColor="#999"
                        keyboardType="default"
                        autoCapitalize="words"
                        autoCorrect={false}
                        onChangeText={setNei}
                    >{neighborhood}</TextInput>

                    <Text style={styles.label}>Cidade</Text>
                    <TextInput
                        style={styles.input}
                        placeholder="Digite o nome da cidade"
                        placeholderTextColor="#999"
                        keyboardType="default"
                        autoCapitalize="words"
                        autoCorrect={false}
                        onChangeText={setCit}
                    >{city}</TextInput>

                    <Text style={styles.label}>Estado</Text>
                    <TextInput
                        style={styles.input}
                        placeholder="Digite o a sigla do estado"
                        placeholderTextColor="#999"
                        keyboardType="default"
                        autoCapitalize="words"
                        autoCorrect={false}
                        onChangeText={setSta}
                    >{state}</TextInput>
                </ScrollView>

                <View style={styles.form}>
                    <TouchableOpacity onPress={handleSubmit} style={styles.button}>
                        <Text style={styles.buttonText}>{farm?'Salvar':'Criar fazenda'}</Text>
                    </TouchableOpacity>
                </View>
            </View>
        </KeyboardAvoidingView>
    )
}
