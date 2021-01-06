import React, { useState, useEffect } from 'react';
import Moment from 'moment';
import {Text,
    View,
    BackHandler,
    ScrollView,
    StyleSheet,
    Image} from 'react-native';
import {Button} from 'react-native-paper';
import styles from '../Styles'
import trans from '../Translate';
import BackButton from '../components/BackButton';
import Footer from '../components/Footer';

export default function Detail({navigation}) {
    const [items, setItems] = useState([]);

    function handleBackButtonClick() {
        console.log("backing...");
        navigation.navigate(navigation.getParam('back'));
        return true;
    }

    useEffect(() => {
        setItems(navigation.getParam('items'));
    }, []);

    useEffect(() => {
        BackHandler.addEventListener('hardwareBackPress', handleBackButtonClick);
        return () => {
            BackHandler.removeEventListener('hardwareBackPress', handleBackButtonClick);
        };
    }, []);

    return (<View style={styles.root}>
        <BackButton
            navigation={navigation}
            back={navigation.getParam('back')}/>
        <ScrollView style={{margin: 15}}>
            <View style={styles.container}>
                <View style={{
                    ...styles.box,
                }}>
                </View>

                {items.map((i) => {
                    if (i.key && i.value) {
                        if (['created_at',
                            'updated_at',
                            'started_at',
                            'finished_at',].indexOf(i.key)>=0){
                            console.log(i);
                            Moment.locale('pt-BR');
                            i.value = Moment(i.value).format('DD MMM YYYY HH:SS');
                        }

                        if (i.key.toLowerCase() == 'cep')
                            i.value = `${i.value.slice(0,5)}-${i.value.slice(5,9)}`

                        if (i.key.toLowerCase() == 'phone')
                            i.value = `(${i.value.slice(0,2)}) ${i.value.slice(2,i.value.length-4)} - ${i.value.slice(i.value.length-4, i.value.length)}`

                        if (trans[i.key])
                            i.key = trans[i.key];

                        return (<View key={`${i.parent?i.parent:''}${i.key}`} style={{
                            ...styles.row,
                            ...styles.box,
                        }}>
                            <Text>{i.key} </Text>
                            <Text>{i.value}</Text>
                        </View>);
                    } else if (i.title) {
                        return (<View key={i.title} style={styles.title}>
                            <Text style={styles.title}>{i.title}</Text>
                        </View>);
                    }
                })}

            </View>
        </ScrollView>
        <Footer/>
    </View>);
}
