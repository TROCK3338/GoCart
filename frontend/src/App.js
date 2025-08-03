import React, { useState, useEffect } from 'react';
import Login from './components/Login';
import ListItems from './components/ListItems'; // Import the new component
import './App.css';

function App() {
  const [isLoggedIn, setIsLoggedIn] = useState(false);

  useEffect(() => {
    const token = localStorage.getItem('token');
    if (token) {
      setIsLoggedIn(true);
    }
  }, []);

  const handleLoginSuccess = () => {
    setIsLoggedIn(true);
  };

  const handleLogout = () => {
    localStorage.removeItem('token');
    setIsLoggedIn(false);
  }

  return (
    <div className="App">
      <header className="App-header">
        <h1>GoCart</h1>
        {isLoggedIn && <button onClick={handleLogout}>Logout</button>}
      </header>
      {isLoggedIn ? (
        <ListItems /> // Render the ListItems component
      ) : (
        <Login onLoginSuccess={handleLoginSuccess} />
      )}
    </div>
  );
}

export default App;