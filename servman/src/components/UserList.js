import React from 'react';
import {Text,
    View,
    ScrollView,
    StyleSheet,
    Image} from 'react-native';
import {Button} from 'react-native-paper';
import styles from '../Styles'

function UserList(props) {
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
            }}>Usuários</Text>
        </View>


        {props.users && props.users.length>0 ?(props.users.map(user => {
            console.log('PRINTING, ', user);
            return user.roles.indexOf('root')<0 || props.curr.roles.indexOf('root')>=0?(
                <View key={user.id} style={{
                    ...styles.box,
                    ...styles.border,
                }}>

                    <View style={{...styles.row,
                        justifyContent: 'flex-end',
                    }}>
                        <Button
                            onPress={() => props.onEdit(user)}
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
                                props.onDetail(user);
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
                            onPress={() => props.onRemove(user)}
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
                    }}>{user.name}</Text>

                    <Text style={{
                        color: '#555',
                        fontSize: 16,
                    }}>CPF: {user.document}</Text>

                    <Text style={{
                        color: '#555',
                        fontSize: 16,
                    }}>Telefone: ({user.phone.slice(0, 2)}) {user.phone.slice(2, user.phone.length-4)}-{user.phone.slice(user.phone.length-4, user.phone.length)}</Text>

                    <View style={styles.row}>
                        {user.roles?(
                            user.roles.map(
                                role => (
                                    <Text key={role} style={{
                                        color: '#F55',
                                        fontSize: 16,
                                    }}>
                                        {' '+role+' '}
                                    </Text>
                                )
                            )):null}
                    </View>

                </View>):(
                    <View key={user.id} style={{
                        ...styles.box,
                        ...styles.border,
                    }}>
                        <Button icon={({ size, color }) => (
                            <Image
                                source={require("../assets/lock.png")}
                                style={{
                                    width: size,
                                    height: size,
                                    tintColor: '#23B185',
                                }}/>)}>
                            <Text style={{
                                color: '#F55',
                                fontSize: 16,
                            }}>{user.name} (admin)</Text>
                        </Button>
                    </View>);
        } )) :(<Text style={{
            color: '#555',
            fontSize: 17,
            padding: 15,
            margin: 15,
        }}>{props.users?'carregando...':'Nenhum usuário foi encontrado.'}</Text>
        )
        }</ScrollView>);
}

export default UserList;
