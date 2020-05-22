import React from 'react';
import './App.css';
import Customers from "./components/Customers";

const App = () => {
  return (
    <div className="App">
      <header className="App-header">
        <h1>Customers</h1>
        <Customers />
      </header>
    </div>
  );
}

export default App;
