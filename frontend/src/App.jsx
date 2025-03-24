import React, { useState } from 'react';
import { Container } from 'react-bootstrap';
import Contacts from './components/Contacts';
import Header from './components/Header';
import WeddingDate from './components/WeddingDate';
import WeddingList from './components/WeddingList';
import WeddingLocation from './components/WeddingLocation';
import WeddingModal from './components/WeddingModal';
import useSettings from './hooks/useSettings';
import useWeddingList from './hooks/useWeddingList';

function App() {
  const { title, description, whatsappGroom, whatsappBride, nameGroom, nameBride, cellPhoneGroom, cellPhoneBride } = useSettings();
  const weddingList = useWeddingList();
  const [selectedItem, setSelectedItem] = useState(null);
  const [showModal, setShowModal] = useState(false);

  const handleShowModal = (item) => {
    setSelectedItem(item);
    setShowModal(true);
  };

  const handleCloseModal = () => {
    setShowModal(false);
    setSelectedItem(null);
  };

  return (
    <Container fluid className="bg-light text-light p-0 m-0">
      <Header
        title={title}
        description={description}
      />
      <WeddingDate
        title="Data de Casamento"
        date="2025-11-09T16:00:00"
      />
      <WeddingLocation
        title="Localização do Casamento"
        src="https://www.google.com/maps/embed?pb=!1m18!1m12!1m3!1d3153.019019726204!2d-122.419415484681!3d37.7749292797596!2m3!1f0!2f0!3f0!3m2!1i1024!2i768!4f13.1!3m3!1m2!1s0x8085809c5b0b0b0b%3A0x0!2zMzfCsDQ2JzI5LjgiTiAxMjLCsDI1JzEwLjkiVw!5e0!3m2!1sen!2sus!4v1633045678901!5m2!1sen!2sus"
      />
      <Contacts
        whatsappGroom={whatsappGroom}
        whatsappBride={whatsappBride}
        nameGroom={nameGroom}
        nameBride={nameBride}
      />
      <WeddingList
        weddingList={weddingList}
        handleShowModal={handleShowModal}
      />
      <WeddingModal
        showModal={showModal}
        handleCloseModal={handleCloseModal}
        selectedItem={selectedItem}
        whatsappGroom={whatsappGroom}
        whatsappBride={whatsappBride}
        nameGroom={nameGroom}
        nameBride={nameBride}
        cellPhoneGroom={cellPhoneGroom}
        cellPhoneBride={cellPhoneBride}
      />
    </Container>
  );
}

export default App;
