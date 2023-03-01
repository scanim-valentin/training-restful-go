import React, { useState, useEffect } from 'react';
import ReactDOM from 'react-dom/client';
import './index.css';
import * as model from './model.js'
import * as database from './database.js'
const IP = 'localhost'
const Port = '3001'
const CreatedStatusCode = 201
/*******CLIENT*********/
/**
 * A form to sign in
 */

function FormSignIn(props) {
  const onSubmitHandler = props.onSubmit
  const [IDInput, setIDInput] = useState("")

  const handleSignIn = () => {
    fetch('http://' + IP + ':' + Port + '/login?id=' + IDInput)
      .then(response => response.json())
      .then(data => onSubmitHandler({
        'ID': data.ID,
        'UserList': data.UserList
      }))
  }
  const handleChangeInput = (event) => {
    setIDInput(event.target.value)
  }
  return (
    <form>
      <label htmlFor="usernameinput"> ID </label>
      <input type="text" id="usernamefield" value={IDInput} onChange={handleChangeInput} /> <br />
      <input type="button" id="signin_submit" value="Sign In" onClick={handleSignIn} />
    </form>
  )
}

/**
 * A form to sign 
 * DON'T: hacked','6.6.6.6','666'),('hacked2','::1','666'),('hackedagain
 */
function FormSignUp(props) {
  const onSubmitHandler = props.onSubmit
  const [nameInput, setNameInput] = useState("Empty")

  const handleSignUp = () => {
    // HTTP Query
    fetch('http://' + IP + ':' + Port + '/register?name=' + nameInput)
      .then(response => response.json())
      .then(data => onSubmitHandler({
        'ID': data.ID,
        'UserList': data.UserList
      }))
  }

  const handleChangeInput = (event) => {
    setNameInput(event.target.value)
  }

  return (
    <form>
      <label htmlFor="idinput"> Username </label>
      <input type="text" id="idfield" value={nameInput} onChange={handleChangeInput} /><br />
      <input type="button" id="signup_submit" value="Sign up" onClick={handleSignUp} />
    </form>
  )
}

function ConnectionFrame(props) {
  const SwitchToChat = props.onSubmit
  const elements = {
    'sign_in': {
      'label': 'Sign In',
      'form': <FormSignIn onSubmit={SwitchToChat} />
    },
    'sign_up': {
      'label': 'Sign Up',
      'form': <FormSignUp onSubmit={SwitchToChat} />
    }
  }
  return (
    <div className="ConnectionFrame">
      <model.SwitchableForm elements={elements} />
    </div>
  )
}




/**
 * - {obj} destination : { {string} ID, {string} name}
 * - {string} source : source ID
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
            {value.Content}
          </div>
        )}
      </div>
    )
  }

  const [message, setMessage] = useState({})
  const [newconversation, setNewConversation] = useState([])
  const destination = props.destination
  const sourceid = props.sourceid

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
    setMessage(new database.Message(0, sourceid, destination.id, event.target.value))
  }


  const [init_conv, setConv] = useState([])
  useEffect(() => {
    /*For some reasons this is necessary to reset the state values when switching from one user to another*/
    setMessage({})
    setNewConversation([])
    setConv([])
    // HTTP request
    fetch('http://' + IP + ':' + Port + '/select?user=' + sourceid + '&other=' + destination.id)
      .then(response => response.json())
      .then(data => {
        if (data.Messages !== null) {
          setConv(data.Messages)
        }
      }
      )
  }, [destination, sourceid])

  return (
    <div className='ChatArea'>
      <b>You're talking to {destination.name}</b> <br />
      <div>
        <Conversation messages={init_conv.concat(newconversation)} sourceid={sourceid} destination={destination} />
      </div>
      <textarea onChange={handleOnChange} defaultValue="" />
      <input type="button" value=">" onClick={handleSend} />
    </div>
  )
}

function ChatFrame(props) {

  useEffect(() => {
    const handleTabClose = () => {
      fetch('http://' + IP + ':' + Port + '/logout?id=' + props.sourceid)
    };

    window.addEventListener('beforeunload', handleTabClose);

    return () => {
      window.removeEventListener('beforeunload', handleTabClose);
    };
  }, [props.sourceid]);


  const users = Object.assign({}, ...props.UserList.map(
    (user) => ({ [user.ID]: { 'label': user.Name, 'form': < ChatArea destination={{ 'id': user.ID, 'name': user.Name }} sourceid={props.sourceid} /> } })
  ));
  return (
    <div>
      <b>PLEASE SELECT YOUR FRIEND</b>
      <model.SwitchableForm elements={users} />
    </div>
  )
}

function App() {
  const [currentElement, setCurrentElement] = useState('Login')
  const [UserList, setUserlist] = useState({})
  const [UserID, setUserID] = useState({})
  const switchToChat = (props) => {
    setUserlist(props.UserList)
    setUserID(props.ID)
    setCurrentElement('Chat')
  }

  const elements = [
    {
      'label': 'Login',
      'frame': <ConnectionFrame onSubmit={switchToChat} />
    },
    {
      'label': 'Chat',
      'frame': <ChatFrame UserList={UserList} sourceid={UserID} />
    }
  ]

  return (
    <div>
      <model.SwitchableFrame elements={elements} currentElement={currentElement} />
    </div>
  )
}

function UI () {

    return (
      <div className="ui">
        <App />
      </div>
    );
}


// ========================================

const root = ReactDOM.createRoot(document.getElementById("root"));

root.render(<UI />);
