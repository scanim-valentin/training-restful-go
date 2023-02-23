import React, { useState } from 'react';
import ReactDOM from 'react-dom/client';
import './index.css';



/**
 *   List of ConnectMode components allowing a single button to be selected at a time
 *   @param {props} map Should have:
 *   - labels: List of unique strings
 *   - onChange: Function to run when an element is selected
 *   - defaultOption: Which option should be checked by default
 */
function RadioSelectList(props) {
  const [id_checked, setChecked] = useState(props.defaultOption);
  const labels = props.labels
  const onChangeHandler = props.onChange

  /**
   *  A radio button hook that requires an id (Both unique identifier and label)
   *  @param {map} props Should have:
   *  - id: String
   */
  const RadioSelect = (props) => {
    const id = props.id
    const handleLoginChecked = () => {
      onChangeHandler(id)
      setChecked(id)
    }

    return (
      <div className="RadioSelect">
        <input type="radio" value={id} onClick={handleLoginChecked} defaultChecked={id === id_checked} />
        <label htmlFor={id}> {id} </label>
      </div>
    )
  }
  let listConnectModes = labels.map((label, index) => (<RadioSelect id={label} key={index} />))

  return (
    <div>
      {listConnectModes}
    </div>
  )
}

/**
 * Hook that allows the user to select a form in a list and render it dynamically
 * @param {map} props Should have:
 * - elements: array of structure {'label': string, 'form: form' }
 */
function SwitchableForm(props) {

  const labels = props.elements.map(
    (element) => element.label
  )
  const [selectedLabel, setSelectedLabel] = useState(labels[0])
  const forms = Object.assign({}, ...props.elements.map(
    (element) => ({ [element.label]: element.form })
  ));

  const switchMode = (label) => {
    setSelectedLabel(label)
  }

  return (
    <div>
      <RadioSelectList labels={labels} onChange={switchMode} defaultOption={selectedLabel} />
      <div>
        {
          forms[selectedLabel]
        }
      </div>
    </div>
  )
}

/**
 * Hook to dynamically switch from one screen to another
 * @param {map} props Should have:
 * - currentElement: current element label to be rendered 
 * - elements: array of structure: 
 *  'label': string
 *  'frame': frame
 */
function SwitchableFrame(props) {
  const currentLabel = props.currentElement
  const forms = Object.assign({}, ...props.elements.map(
    (element) => ({ [element.label]: element.frame })
  ));

  return (
    <div>
      <div>
        {
          forms[currentLabel]
        }
      </div>
    </div>
  )
}

/*******CLIENT*********/
/**
 * A form to sign in
 */

function FormSignIn(props) {
  const onSubmitHandler = props.onSubmit
  const [IDInput, setIDInput] = useState("00000")
  
  const handleSignIn = () => {
    fetch('http://localhost:3001/login?id=' + IDInput)
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
    fetch('http://localhost:3001/register?name=' + nameInput)
      .then(response => response.json() )
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
  const elements = [
    {
      'label': 'Sign In',
      'form': <FormSignIn onSubmit={SwitchToChat} />
    },
    {
      'label': 'Sign Up',
      'form': <FormSignUp onSubmit={SwitchToChat} />
    }
  ]

  return (
    <div className="ConnectionFrame">
      <SwitchableForm elements={elements} />
    </div>
  )
}


function ChatFrame(props) {

  const selectListElements = props.UserList.map(
    (element) => ({
      'label': [element.ID],
      'form': <textarea value={"You're talking to "+element.ID}/>
    }))
  return (
    <div>
      <b>PLEASE SELECT YOUR FRIEND</b>
      <SwitchableForm elements = {selectListElements} />
    </div>
  )
  

}

function App() {
  const [currentElement, setCurrentElement] = useState('Login')
  const [UserList, setUserlist] = useState({})
  const switchToChat = (props) => {
    setUserlist(props.UserList)
    console.log(props)
    setCurrentElement('Chat')
  }

  const elements = [
    {
      'label': 'Login',
      'frame': <ConnectionFrame onSubmit={switchToChat} />
    },
    {
      'label': 'Chat',
      'frame': <ChatFrame UserList={UserList}/>
    }
  ]

  return (
    <div>
      <SwitchableFrame elements={elements} currentElement={currentElement} />
    </div>
  )
}

class UI extends React.Component {
  render() {
    return (
      <div className="ui">
        <App />
      </div>
    );
  }
}

// ========================================

const root = ReactDOM.createRoot(document.getElementById("root"));
root.render(<UI />);
