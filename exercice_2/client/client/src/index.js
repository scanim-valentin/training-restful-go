import React from 'react';
import ReactDOM from 'react-dom/client';
import './index.css';

class AddApp extends React.Component {
    render() {
      return (
        <button className="addapp">
          {/* TODO */}
        </button>
      );
    }
  }

  var callAdd = function() {
    console.log('Calling backend . . .');
                fetch("http://localhost:3001/add")
                .then((response) => {
                    if(response.status === 200){
                        console.log("SUCCESSS")
                        response.blob()
                        .then(array => array.text())
                        .then(text => console.log(text))
                        // This shouldn't be so weird
                    }else if(response.status === 408){
                        console.log("SOMETHING WENT WRONG")
                        this.setState({ requestFailed: true })
                    }
                })                
}

  class AddButton extends React.Component {
    render() {
        return (
            <button onClick={() => callAdd()}> 
                Add 
            </button>
            )
        }
    }
  
  class UI extends React.Component {
    render() {
      return (
        <div className="ui">
          <AddButton/>
        </div>
      );
    }
  }
  
  // ========================================
  
  const root = ReactDOM.createRoot(document.getElementById("root"));
  root.render(<UI />);
  