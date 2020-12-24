import React from 'react';
import {Text,
    View,
    ScrollView,
    StyleSheet,
    Image} from 'react-native';
import {Button} from 'react-native-paper';

function FarmList(props) {
    return (<ScrollView styles={styles.container}>
        <Button
            onPress={props.onRefresh}
            style={{
                ...styles.title,
                ...styles.box,
                flex: 0,
                tintColor: '#23B185',
                color: '#23B185',
            }}
            icon={({ size, color }) => (
                <Image
                    source={require("../assets/refresh.png")}
                    style={{
                        width: size,
                        height: size,
                        tintColor: '#23B185',
                    }}/>
            )}>
            <Text style={{
                ...styles.title,
                fontSize: 16
            }}>Fazendas</Text>
        </Button>
        {props.farms?(props.farms.map(farm => {
            console.log(farm);
            return (<View key={farm.id} style={styles.box}>

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
                }}>{farm.name}</Text>

                {farm.addr?(
                    <Text style={{
                        color: '#555',
                        fontSize: 16,
                    }}>
                        Rua {farm.addr.street}, {farm.addr.number}, {farm.addr.city} - {farm.addr.state}
                    </Text>
                ):null}

            </View>);
        } )) :(<Text style={{
            color: '#555',
            fontSize: 17,
            padding: 15,
            margin: 15,
        }}>Nenhuma fazenda foi encontrada.</Text>
        )
        }</ScrollView>);
}

const styles = StyleSheet.create({
    container: {
        flex: 1,
        alignItems: 'center',
        justifyContent: 'center',
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
        flex: 1,
        padding: 15,
        width: '100%',
        minWidth: 300,
        borderRadius: 2,
        borderColor: '#23B185',
        borderWidth: 1,
        marginTop: 15,
    },
});

export default FarmList;
