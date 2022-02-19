import React from 'react';
import Calculator from './components/Calculator';
import { Reports } from './components/logs';

class App extends React.Component {
  render() {
    return (
      <div className="App">
        <Calculator/>
        <Reports />
      </div>

    );
  }
}

export default App;
