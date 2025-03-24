import axios from 'axios';
import React, { useEffect, useState } from 'react';
import './App.css';

function App() {
  const [message, setMessage] = useState('')

  useEffect(() => {
    const url = `${process.env.REACT_APP_API_BASE || ''}/hello`;
    axios.get(url, {
      headers: {
        'Content-Type': 'application/json',
      },
      mode: 'no-cors',
    }).then(({ data }) => setMessage(data.message))
      .catch(error => console.error(error));
  }, []);

  return (
    <>
      <h1>Message from backend: {message}</h1>
    </>
  );
}

export default App;
