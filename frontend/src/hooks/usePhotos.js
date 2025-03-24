import axios from 'axios';
import { useEffect, useState } from 'react';

const usePhotos = () => {
  const [photos, setPhotos] = useState();

  useEffect(() => {
    const loadPhotos = () => {
      const url = `${process.env.REACT_APP_API_BASE || ''}/wedding-photos`;
      axios.get(url, {
        headers: {
          'Content-Type': 'application/json',
        }
      }).then(({ data }) => {
        const { pictures } = data;
        setPhotos(pictures);
      })
        .catch(error => console.error(error));
    };

    loadPhotos();
  }, []);

  return photos;
};

export default usePhotos;
