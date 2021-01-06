import React from 'react';
import {Text,
    View,
    ScrollView,
    StyleSheet,
    Image} from 'react-native';
import {Button} from 'react-native-paper';
import styles from '../Styles'

function FarmList(props) {
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
            }}>Fazendas</Text>
        </View>


        {props.farms && props.farms.length>0 ?(props.farms.map(farm => {
            console.log(farm);
            return (
                <View key={farm.id} style={{
                    ...styles.box,
                    ...styles.border,
                }}>

                    <View style={{...styles.row,
                        justifyContent: 'flex-end',
                    }}>
                        <Button
                            onPress={() => props.onEdit(farm)}
                            icon={({ size, color }) => (
                                <Image
                                    source={require("../assets/pencil.png")}
                                    style={{
                                        width: size,
                                        height: size,
                                        tintColor: '#23B185',
                                    }}/>)}/>
                        <Button
                            onPress={() => {
                                props.onDetail(farm);
                            }}
                            icon={({ size, color }) => (
                                <Image
                                    source={require("../assets/more.png")}
                                    style={{
                                        width: size,
                                        height: size,
                                        tintColor: '#23B185',
                                    }}/>)}/>
                        <Button
                            onPress={() => props.onRemove(farm)}
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
                    }}>{farm.name}</Text>

                    {farm.addr?(
                        <Text style={{
                            color: '#555',
                            fontSize: 16,
                        }}>
                            {farm.addr.street}, {farm.addr.number}, {farm.addr.neighborhood}, {farm.addr.city} - {farm.addr.state}
                        </Text>
                    ):null}

                </View>);
        } )) :(<Text style={{
            color: '#555',
            fontSize: 17,
            padding: 15,
            margin: 15,
        }}>{props.farms?'carregando...':'Nenhuma fazenda foi encontrada.'}</Text>
        )
        }</ScrollView>);
}

export default FarmList;
