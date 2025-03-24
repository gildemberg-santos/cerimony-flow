import axios from 'axios';
import { useEffect, useState } from 'react';

const useSettings = () => {
  const [settings, setSettings] = useState({
    title: '',
    description: '',
    whatsappGroom: '',
    whatsappBride: '',
    nameGroom: '',
    nameBride: '',
    cellPhoneGroom: '',
    cellPhoneBride: ''
  });

  useEffect(() => {
    const loadSettings = () => {
      const url = `${process.env.REACT_APP_API_BASE || ''}/settings`;
      axios.get(url, {
        headers: {
          'Content-Type': 'application/json',
        }
      }).then(({ data }) => {
        const { title, description, whatsapp_groom, whatsapp_bride, name_groom, name_bride, cell_phone_groom, cell_phone_bride } = data;
        setSettings({
          title,
          description,
          whatsappGroom: whatsapp_groom,
          whatsappBride: whatsapp_bride,
          nameGroom: name_groom,
          nameBride: name_bride,
          cellPhoneGroom: cell_phone_groom,
          cellPhoneBride: cell_phone_bride
        });
      })
        .catch(error => console.error(error));
    };

    loadSettings();
  }, []);

  return settings;
};

export default useSettings;
