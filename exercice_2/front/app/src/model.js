import React, { useState } from 'react';

/**
 *   List of ConnectMode components allowing a single button to be selected at a time
 *   @param {props} map Should have:
 *   - labels: List of unique strings
 *   - onChange: Function to run when an element is selected
 *   - defaultOption: Which option should be checked by default
 */
export function RadioSelectList(props) {
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
          <label htmlFor={id}> { labels[id] } </label>
        </div>
      )
    }
    let listConnectModes = Object.entries(labels).map( (entry) => (<RadioSelect id={entry[0]} key={entry[0]} />))
  
    return (
      <div>
        {listConnectModes}
      </div>
    )
  }
  
  /**
   * Hook that allows the user to select a form in a list and render it dynamically
   * @param {map} props Should have:
   * - elements: list of pair id string: {'label': string, 'form': form }
   */
export function SwitchableForm(props) {
    const elements = props.elements
    const [selectedLabel, setSelectedLabel] = useState(Object.entries(elements)[0][0])
    
    const switchMode = (label) => {
      setSelectedLabel(label)
    }

    const labels = Object.assign({}, ...Object.entries(elements).map(
        (element) => ({ [element[0]]: element[1].label })
    ));
    return (
      <div>
        <RadioSelectList labels={labels} onChange={switchMode} defaultOption={selectedLabel} />
        <div>
          {
            elements[selectedLabel].form
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
export function SwitchableFrame(props) {
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
