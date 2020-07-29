import React, { Component } from 'react';
import { BrowserRouter } from 'react-router-dom'
import './App.css';

import Header from '../components/Templates/Header';
import Routes from './Routes'

class App extends Component {
  render() {
    return (
      <BrowserRouter>
        <div className="app">
          <Header />
          <Routes />
        </div>
      </BrowserRouter>
    )
  }
}

export default App;