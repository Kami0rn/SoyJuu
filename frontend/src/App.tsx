import React from 'react';
import logo from './logo.svg';
import './App.css';
import Nav from './Nav/Nav'
import {

  BrowserRouter as Router,
  
  Routes,
  
  Route,
  
  Link,
  
  } from "react-router-dom";

function App() {
  return (
    <div className="App">
      <header className="App-header">
        <Nav />
        <img src={logo} className="App-logo" alt="logo" />
        <p>
          Edit <code>src/App.tsx</code> and save to reload.
        </p>
        <a
          className="App-link"
          href="https://reactjs.org"
          target="_blank"
          rel="noopener noreferrer"
        >
          Learn React
        </a>
      </header>
    </div>
  );
}

export default App;