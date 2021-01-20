import React from 'react';
import {Text,
    View,
    Image } from 'react-native';
import { Button } from 'react-native-paper';
import styles from '../Styles'

import api from '../services/api';
import logo from '../assets/logo.png';
import trans from '../Translate';

export default function TagList({tags, onRemove}) {
    console.log('LISTING TAGS: ', tags);
    return (<View style={{
        ...styles.input,
        flex: 1,
        alignItem: 'center',
        justifyContent: 'space-around',
        padding: 15,
        height: '100%',
    }}>
        <View style={{
            flex: 1,
            alignItem: 'center',
            justifyContent: 'space-around',
            flexOverflow: 'wrap',
        }}>
            {tags.map(r => {
                var i = tags.indexOf(r);
                return (<Button
                    key={r}
                    style={{
                        ...styles.tag,
                        tintColor: '#FFF',
                        color: '#FFF',
                        justifyContent: 'space-around',
                        alignItems: 'center',
                        height: 30,
                        margin: 5,
                    }}
                    onPress={() => onRemove(i)}
                    icon={({ size, color }) => (
                        <Image
                            source={require("../assets/trash.png")}
                            style={{
                                padding: 5,
                                width: size,
                                height: size,
                                tintColor: '#FFF',
                            }}/>
                    )}>
                    <Text style={{
                        color: '#FFF',
                        fontSize: 10,
                    }}>{trans[r]?trans[r]:r}</Text>
                </Button>);

            })}
        </View>
    </View>);
}
