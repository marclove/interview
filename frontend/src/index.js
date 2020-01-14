import React from 'react';
import ReactDOM from 'react-dom';
import './index.css';
import App from './App';
import axios from 'axios';

ReactDOM.render(<App />, document.getElementById('root'));

axios.get("/api").then(r => console.log(r.data));
