import React, { Component } from 'react';
import { BrowserRouter as Router } from 'react-router-dom'
import AppNavbar from './components/AppNavbar'
import AppFooter from './components/AppFooter'
import AppLoader from './components/AppLoader'
import AppNotify from './components/AppNotify'
import AppBody from './components/AppBody'


class App extends Component {
  render() {
    return (
      <Router>
        <main>
          <AppLoader />
          <AppNavbar />
          <AppNotify />
          <AppBody />
        </main>
        <AppFooter />
      </Router>
    );
  }
}

export default App;
