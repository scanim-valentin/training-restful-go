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
    (element) => ({[element.label]: element.form})
    ));
  
  const switchMode = (label) => {
    setSelectedLabel(label)
    console.log("Switched to " + label)
    console.log(forms[label])
  }

  return (
    <div>
      <RadioSelectList labels={labels} onChange={switchMode} defaultOption={selectedLabel}/>
      <div>
        {
          forms[selectedLabel]
        }
      </div>
    </div>
  )
}

function ConnectionFrame() {
  const elements = [
    {'label' : 'Sign In',
    'form' : <form>
      <label htmlFor="usernameinput"> Username </label>
      <input type="text" id="usernamefield" /> <br />
      <input type="submit" id="signin_submit" value="Sign In" />
    </form> },
    {'label': 'Sign Up',
    'form': <form>
      <label htmlFor="idinput"> ID </label>
      <input type="text" id="ID" /> <br />
      <input type="submit" id="signup_submit" value="Sign Up" />
    </form>},
    {'label': 'Other option',
    'form': <form>
      <label htmlFor="other"> Other </label>
      <input type="text" id="ID" /> <br />
      <input type="submit" id="other_submit" value="Other" />
    </form>}
    ]

  return (
    <div className="ConnectionFrame">
      <SwitchableForm elements={elements} />
    </div>
  )
}

class UI extends React.Component {
  render() {
    return (
      <div className="ui">
        <ConnectionFrame />
      </div>
    );
  }
}

// ========================================

const root = ReactDOM.createRoot(document.getElementById("root"));
root.render(<UI />);
