import React, { useState } from 'react';
import axios from 'axios';
import './Login.css'; // New CSS file for styling

const Login = ({ onLoginSuccess }) => {
  const [username, setUsername] = useState('');
  const [password, setPassword] = useState('');

  const handleLogin = async (e) => {
    e.preventDefault();
    try {
      const response = await axios.post('http://localhost:8080/users/login', {
        username,
        password,
      });
      const { token } = response.data;
      localStorage.setItem('token', token);
      onLoginSuccess();
    } catch (error) {
      window.alert('Invalid username/password');
    }
  };

  return (
    <div className="login-container">
      <div className="login-card">
        <h2>Login to GoCart</h2>
        <form onSubmit={handleLogin}>
          <div className="form-group">
            <label>Username:</label>
            <input
              type="text"
              value={username}
              onChange={(e) => setUsername(e.target.value)}
            />
          </div>
          <div className="form-group">
            <label>Password:</label>
            <input
              type="password"
              value={password}
              onChange={(e) => setPassword(e.target.value)}
            />
          </div>
          <button type="submit" className="login-button">Login</button>
        </form>
      </div>
    </div>
  );
};

export default Login;