import React, { useState, useEffect } from 'react';
import './index.css';
import * as model from './model.js'
import * as database from './database.js'
const IP = 'localhost'
const Port = '3001'
const CreatedStatusCode = 201

/**
 * select call rate to update conversation (ms)
 */
// const RefreshRate = 1000*3600
const RefreshRate = 200
/**
 * - {array} messages
 * } props
 * @param props
 */
function ChatArea(props) {

    function Conversation(props) {
        const messages = props.messages

        return (
            <div className='Conversation'>
                {messages.map((value, index) =>
                    <div key={index}>
                        <small>{value.Time}</small><br />
                        <b>{value.Source.toString() === props.userid.toString() ? props.destination.name : props.source.name}    </b>{value.Content}
                    </div>
                )}
            </div>
        )
    }

    const [message, setMessage] = useState({})
    const [newconversation, setNewConversation] = useState([])
    const destination = props.destination
    const source = props.source

    const handleSend = () => {
        if (message.Content) {
            // HTTP Query
            let headers = {
                'Content-Types': 'application/json; charset=UTF-8'
            }
            let body = JSON.stringify(message)
            fetch('http://' + IP + ':' + Port + '/send', { method: 'POST', headers: headers, body: body })
                .then(response => handleServerResponse(response))
        }
    }
    const handleServerResponse = (response) => {
        if (response.status !== CreatedStatusCode) {
            console.log('Server Error \n')
            console.log(response)
        } else {
            setNewConversation(newconversation.concat(message))
        }
    }
    const handleOnChange = (event) => {
        setMessage(new database.Message(0, source.id, destination.id, event.target.value))
    }


    const [init_conv, setConv] = useState([])
    useEffect(() => {
        /*For some reasons this is necessary to reset the state values when switching from one users to another*/
        setMessage({})
        setNewConversation([])
        setConv([])

        const interval = setInterval(() => {
            // HTTP request
            fetch('http://' + IP + ':' + Port + '/select?user=' + source.id + '&other=' + destination.id)
                .then(response => response.json())
                .then(data => {
                    if (data.Messages !== null) {
                        setConv(data.Messages)
                    }
                })
        }, RefreshRate)
        return () => clearInterval(interval);
    }, [destination, source])

    return (
        <div className='ChatArea'>
            <b>You're talking to {destination.name}</b> <br />
            <div>
                <Conversation userid={props.userid} messages={/*init_conv.concat(newconversation)*/ init_conv} source={source} destination={destination} />
            </div>
            <textarea onChange={handleOnChange} defaultValue="" />
            <input id="sendbutton" type="button" value=">" onClick={handleSend} />
        </div>
    )
}

export function ChatFrame(props) {

    useEffect(() => {
        const handleTabClose = () => {
            fetch('http://' + IP + ':' + Port + '/logout?id=' + props.source.id)
        };

        window.addEventListener('beforeunload', handleTabClose);

        return () => {
            window.removeEventListener('beforeunload', handleTabClose);
        };
    }, [props.source.id]);

    const users = Object.assign({}, ...props.UserList.map(
        (user) => ({ [user.ID]: { 'label': user.Name, 'form': < ChatArea userid={user.ID} destination={{ 'id': user.ID, 'name': user.Name }} source={{ 'id': props.source.id, 'name': props.source.name }} /> } })
    ));
    return (
        <div>
            <b id="idPanel" name={props.source.name} uniqueid={props.source.id}>Connected as {props.source.name} (unique ID: {props.source.id})</b>
            <model.SwitchableForm elements={users} />
        </div>
    )
}