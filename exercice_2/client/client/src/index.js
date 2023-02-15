import React from 'react';
import ReactDOM from 'react-dom/client';
import './index.css';

class Square extends React.Component {
    render() {
      return (
        <button className="square">
          {/* TODO */}
        </button>
      );
    }
  }
  
  class Board extends React.Component {
    renderSquare(i) {
      return <Square />;
    }
  
    render() {
      const status = 'Next player: X';
  
      return (
        <div>
          <div className="status">{status}</div>
          <div className="board-row">
            {this.renderSquare(0)}
            {this.renderSquare(1)}
            {this.renderSquare(2)}
          </div>
          <div className="board-row">
            {this.renderSquare(3)}
            {this.renderSquare(4)}
            {this.renderSquare(5)}
          </div>
          <div className="board-row">
            {this.renderSquare(6)}
            {this.renderSquare(7)}
            {this.renderSquare(8)}
          </div>
        </div>
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
  