import React from 'react';
import {Text,
    View,
    ScrollView,
    StyleSheet,
    Image} from 'react-native';
import {Button} from 'react-native-paper';

function ServList (props) {
    return (<ScrollView styles={styles.container}>{props.servs?(props.servs.map(serv => {
        var begin = new Date(serv.started_at)
        var end = new Date(serv.finished_at)
        var hours = Math.abs(
            end - begin
        );
        hours = hours/1000/60/60; // converting milisec to hours
        return (
            <View key={serv.id} style={styles.box}>

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
    }
    ))
    :(<Text style={{
        color: '#555',
        fontSize: 17,
        padding: 15,
        margin: 15,
    }}>Nenhum serviço foi encontrado.</Text>)}</ScrollView>);
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

export {ServList};
