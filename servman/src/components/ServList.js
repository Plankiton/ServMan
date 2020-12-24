import React from 'react';
import {Text,
    View,
    ScrollView,
    Image} from 'react-native';
import {Button} from 'react-native-paper';

import styles from '../Styles'

function ServList (props) {
    return (<ScrollView>
        <View style={{
            ...styles.box,
        }}>
            <View style={{
                ...styles.row,
                flex: 1,
                alignSelf: 'stretch',
                justifyContent: 'flex-end',
            }}>
                <Button
                    onPress={props.onRefresh}
                    style={{
                        flex: 0,
                        tintColor: '#23B185',
                        color: '#23B185',
                    }}
                    icon={({ size, color }) => (
                        <Image
                            source={require("../assets/refresh.png")}
                            style={{
                                padding: 5,
                                width: size,
                                height: size,
                                tintColor: '#23B185',
                            }}/>
                    )}>
                </Button>

                <Button
                    onPress={props.onCreate}
                    style={{
                        flex: 0,
                        tintColor: '#23B185',
                        color: '#23B185',
                    }}
                    icon={({ size, color }) => (
                        <Image
                            source={require("../assets/plus.png")}
                            style={{
                                padding: 10,
                                width: size,
                                height: size,
                                tintColor: '#23B185',
                            }}/>
                    )}>
                </Button>
            </View>

            <Text style={{
                ...styles.title,
                fontSize: 16
            }}>Serviços</Text>
        </View>


        {props.servs && props.servs.length>0 ?(props.servs.map(serv => {
            var begin = new Date(serv.started_at)
            var end = new Date(serv.finished_at)
            var hours = Math.abs(
                end - begin
            );
            hours = hours/1000/60/60; // converting milisec to hours
            return (
                <View key={serv.id} style={{
                    ...styles.box,
                    ...styles.border,
                }}>

                    <View style={{...styles.row,
                        justifyContent: 'flex-end',
                    }}>
                        <Button
                            onPress={props.onEdit}
                            icon={({ size, color }) => (
                                <Image
                                    source={require("../assets/pencil.png")}
                                    style={{
                                        width: size,
                                        height: size,
                                        tintColor: '#23B185',
                                    }}/>)}/>
                        <Button
                            onPress={props.onRemove}
                            icon={({ size, color }) => (
                                <Image
                                    source={require("../assets/trash.png")}
                                    style={{
                                        width: size,
                                        height: size,
                                        tintColor: '#23B185',
                                    }}/>)}/>
                    </View>


                    <Text style={{
                        color: '#555',
                        fontWeight: 'bold',
                        fontSize: 16,
                    }}>{serv.description}</Text>


                    {hours>0?(<Text style={{
                        color: '#555',
                        fontSize: 16,
                    }}>Carga horária: {
                        Math.trunc(hours)>0? `${Math.trunc(hours)} hora`+(
                            Math.trunc(hours)>1?'s':''
                        ): '' } {  hours%1>0?((Math.trunc(hours)>0?'e':''
                        )+ ` ${Math.trunc((hours%1*100))} minuto`+(
                            Math.trunc(hours%1*100)>1?'s':'') ):''
                        }</Text>):null}


                    <Text style={{
                        color: '#555',
                        fontSize: 16,
                    }}>Preço: {(serv.price*hours).toFixed(2).replace('.',',')} R$</Text>



                </View>
            );
        } )): (<Text style={{
            color: '#555',
            fontSize: 17,
            padding: 15,
            margin: 15,
        }}>{props.servs?'carregando...':'Nenhum serviço foi encontrado.'}</Text>
        )}</ScrollView>);
}

export default ServList;
