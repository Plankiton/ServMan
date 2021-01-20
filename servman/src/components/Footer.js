import React from 'react';
import {Text, View} from 'react-native';
import styles from '../Styles'

export default function Footer(props) {
    return (<View style={{
            margin: 15,
            alignSelf: 'center',
            ...styles.center,
        }}>
        <Text>CopyRight &copy; Plankiton &lt;pl4nk1ton@gmail.com&gt;</Text>
    </View>);
}
