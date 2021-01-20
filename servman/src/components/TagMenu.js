import React from 'react';
import {Text,
    View,
    TouchableOpacity,
    Image } from 'react-native';
import { Button } from 'react-native-paper';
import styles from '../Styles'

import api from '../services/api';
import logo from '../assets/logo.png';
import trans from '../Translate';
import Menu, {MenuItem, MenuDivider} from 'react-native-material-menu';

export default function TagMenu({tags, curr_tags, crols, onSelect}) {
    let _menu = null;
    return (<View style={{
        flex: 1,
        alignItem: 'center',
        justifyContent: 'space-around',
    }}>
        <Menu
            ref={(ref) => (_menu = ref)}
            button={(<TouchableOpacity style={{
                alignSelf: 'center',
                color: '#fff',
                backgroundColor: '#23B185',
                padding: 5,
                borderRadius: 5,
                marginVertical: 5,
            }} onPress={() => _menu.show()}>
                <Text style={{ color: '#fff' }}>Adicionar</Text>
            </TouchableOpacity>)}>
            {tags.map(o => {
                if (curr_tags)
                    if (curr_tags.indexOf(o)>=0)
                        return null;
                if (o == 'root')
                    if (crols.indexOf('root')<0)
                        return null;
                if (o == 'admin')
                    if (crols.indexOf('root')<0&&crols.indexOf('admin')<0)
                        return null;

                return (<MenuItem
                    key={o}
                    onPress={() => onSelect(o, _menu)}>
                    <Text style={{ color: '#444' }}>{trans[o]?trans[o]:o}</Text>
                </MenuItem>);
            })}
            <MenuItem onPress={() => _menu.hide()}>
                <Text style={{ color: '#444' }}>Cancelar</Text>
            </MenuItem>
        </Menu>
    </View>);
}
