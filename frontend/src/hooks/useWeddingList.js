import axios from 'axios';
import { useEffect, useState } from 'react';

const useWeddingList = () => {
  const [weddingList, setWeddingList] = useState([]);

  const loadWeddingList = () => {
    const url = `${process.env.REACT_APP_API_BASE || ''}/wedding-list`;
    axios.get(url, {
      headers: {
        'Content-Type': 'application/json',
      }
    }).then(({ data }) => setWeddingList(data.wedding_list))
      .catch(error => console.error(error));
  };

  useEffect(() => {
    loadWeddingList();
  }, []);

  return weddingList;
};

export default useWeddingList;
