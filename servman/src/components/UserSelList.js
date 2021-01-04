import React, {useState} from 'react';
import {Text,
    View,
    ScrollView,
    TouchableOpacity,
    StyleSheet,
    Image} from 'react-native';
import {Button} from 'react-native-paper';
import styles from '../Styles'

export default function UserSelList(props) {
    const [selected, setSel] = useState(null);
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
                }}>Selecione o dono da fazenda:</Text>
            </Button>
        </View>


        <ScrollView>
            {props.users && props.users.length>0 ?(props.users.map(user => {
                console.log('PRINTING, ', user);
                return user.roles.indexOf('root')<0 || props.curr.roles.indexOf('root')>=0?(
                    <View key={user.id} style={{
                        ...styles.box,
                        ...styles.border,
                        ...{
                            backgroundColor: (user == selected ? '#23B18522': '#FFF'),
                        }
                    }}>

                        <TouchableOpacity onPress={() => {
                            setSel(user);
                            props.onSelect(user);
                        }}>

                            <View style={styles.box}>
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
                                                <Text style={{
                                                    color: '#F55',
                                                    fontSize: 16,
                                                }}>
                                                    {' '+role+' '}
                                                </Text>
                                            )
                                        )):null}
                                </View>
                            </View>
                        </TouchableOpacity>

                    </View>):null;
            } )) :(<Text style={{
                color: '#555',
                fontSize: 17,
                padding: 15,
                margin: 15,
            }}>{props.users?'carregando...':'Nenhum usuário foi encontrado.'}</Text>
            )
            }</ScrollView>
    </View>);
};
