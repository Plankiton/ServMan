import { StyleSheet } from 'react-native';

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
    },
    center: {
        alignItems:'center',
        justifyContent:'center',
    },
    title: {
        color: '#23B185',
        fontWeight: 'bold',
        fontSize: 16,
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
        width: '100%',
        minWidth: 300,
        marginTop: 10,
        alignSelf: 'stretch',
        paddingHorizontal: 30,
        paddingVertical: 10,
        justifyContent: 'space-between',
        marginTop: 30,
    },
    login: {
        alignSelf: 'stretch',
        alignItems: 'center',
        paddingHorizontal: 30,
        paddingVertical: 10,
        justifyContent: 'space-between',
        marginTop: 30,
    },
    border:{
        borderRadius: 2,
        borderTopWidth: 1,
        borderColor: '#23B185',
    },

    line:{
        borderRadius: 2,
        borderTopWidth: 1,
        borderColor: '#23B185',
        minWidth: "1000%",
   },

   form:{
        alignSelf: 'stretch',
        paddingHorizontal: 30,
        marginTop: 30,
        marginBottom: 30,
   },
   label: {
       fontWeight: 'bold',
       color:'#444',
       marginBottom:8
   },
   input: {
       borderWidth:1,
       borderColor: '#ddd',
       paddingHorizontal:20,
       fontSize: 16,
       color:'#444',
       height: 44,
       marginBottom: 20,
       borderRadius: 2
   },
   button: {
       height: 42,
       backgroundColor: '#23B185',
       justifyContent: 'center',
       alignItems:'center',
       borderRadius:2,
   },
   buttonText:{
       color: '#FFF',
       fontWeight:'bold',
       fontSize:16,
   },
   linkText: {
       color: '#23B185',
       fontWeight:'bold',
       fontSize:16,
   },

   container: {
       flex:1,
       justifyContent:'center',
       alignItems:'center'
   }, 

   centerBox:{
        alignSelf: 'stretch',
        paddingHorizontal: 30,
        alignItems: 'center',
        justifyContent: 'space-between',
        marginTop: 30,
   },
   form:{
        alignSelf: 'stretch',
        paddingHorizontal: 30,
        marginTop: 30,
        marginBottom: 30,
        maxHeight: "70%",
   },
   label: {
       fontWeight: 'bold',
       color:'#444',
       marginBottom:8
   },
   input: {
       borderWidth:1,
       borderColor: '#ddd',
       paddingHorizontal:20,
       fontSize: 16,
       color:'#444',
       height: 44,
       marginBottom: 20,
       borderRadius: 2
   },
   button: {
       height: 42,
       backgroundColor: '#23B185',
       justifyContent: 'center',
       alignItems:'center',
       borderRadius:2,
   },
   buttonText:{
       color: '#FFF',
       fontWeight:'bold',
       fontSize:16,
   },
   linkText: {
       color: '#23B185',
       fontWeight:'bold',
       fontSize:16,
   },
});

export default styles;
