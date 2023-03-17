import React, { useState } from 'react';
import './index.css';
import * as model from './model.js'
const IP = 'localhost'
const Port = '3001'

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
        'Username': data.Username,
        'UserList': data.UserList
      }))
  }
  const handleChangeInput = (event) => {
    setIDInput(event.target.value)
  }
  return (
    <form>
      <label htmlFor="idinput"> ID </label>
      <input type="text" id="idfield" value={IDInput} onChange={handleChangeInput} /> <br />
      <input type="button" id="signin_submit" value="Sign In" onClick={handleSignIn} />
    </form>
  )
}

/**
 * A form to sign
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
        'Username': data.Username,
        'UserList': data.UserList
      }))
  }

  const handleChangeInput = (event) => {
    setNameInput(event.target.value)
  }

  return (
    <form>
      <label htmlFor="nameinput"> Username </label>
      <input type="text" id="namefield" value={nameInput} onChange={handleChangeInput} /><br />
      <input type="button" id="signup_submit" value="Sign up" onClick={handleSignUp} />
    </form>
  )
}

export function ConnectionFrame(props) {
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