import React, {useState} from 'react';
import {Text,
    View,
    ScrollView,
    TouchableOpacity,
    StyleSheet,
    Image} from 'react-native';
import {Button} from 'react-native-paper';
import styles from '../Styles'

export default function FarmSelList(props) {
    const [selected, setSel] = useState(null);
    console.log('ALL farms, ', props.farms);
    return (<View style={styles.container}>

        <View style={{
            ...styles.box,
            flex: 0,
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
                <Text style={{
                    ...styles.title,
                    fontSize: 16
                }}>Selecione a fazenda:</Text>
            </Button>
        </View>


        <ScrollView>

            {props.farms && props.farms.length>0 ?(props.farms.map(farm => {
                console.log(farm);
                return (
                    <View key={farm.id} style={{
                        ...styles.box,
                        ...styles.border,
                        ...{
                            backgroundColor: (farm == selected ? '#23B18522': '#FFF'),
                        }
                    }}>

                        <TouchableOpacity onPress={() => {
                            setSel(farm);
                            props.onSelect(farm);
                        }}>


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
                        </TouchableOpacity>

                    </View>);
            } )) :(<Text style={{
                color: '#555',
                fontSize: 17,
                padding: 15,
                margin: 15,
            }}>{props.farms?'carregando...':'Nenhuma fazenda foi encontrada.'}</Text>
            )
            }</ScrollView>
    </View>);
};
