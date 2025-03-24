import axios from 'axios';
import React, { useEffect, useState } from 'react';
import './App.css';

function App() {
  const [message, setMessage] = useState('')

  useEffect(() => {
    axios.defaults.headers.post['Content-Type'] = 'application/json';
    axios.defaults.headers.post['Access-Control-Allow-Origin'] = '*';

    const url = process.env.REACT_APP_API_BASE || "" + '/hello';
    axios.get(url).then(({ data }) => setMessage(data.message));
  }, []);

  return (
    <>
      <h1>Message from backend: {message}</h1>
    </>
  );
}

export default App;
