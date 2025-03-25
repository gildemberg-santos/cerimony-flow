import React, { useState } from 'react';
import { Container } from 'react-bootstrap';
import Contacts from './components/Contacts';
import Header from './components/Header';
import WeddingDate from './components/WeddingDate';
import WeddingList from './components/WeddingList';
import WeddingLocation from './components/WeddingLocation';
import WeddingModal from './components/WeddingModal';
import WeddingPhoto from './components/WeddingPhoto';
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
    <Container fluid className="bg-texture p-0 m-0">
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
        src="https://www.google.com/maps/embed?pb=!1m18!1m12!1m3!1d260.1550136434166!2d-38.45137562854816!3d-3.798811071446984!2m3!1f0!2f0!3f0!3m2!1i1024!2i768!4f13.1!3m3!1m2!1s0x7c745af2d77f751%3A0x33210c6d90df141!2sR.%20Sabiaguaba%2C%201270%20-%20Edson%20Queiroz%2C%20Fortaleza%20-%20CE%2C%2060835-750!5e1!3m2!1spt-BR!2sbr!4v1742851282611!5m2!1spt-BR!2sbr"
      />
      <WeddingPhoto
        title="Galeria"
      />
      <Contacts
        title="Contatos"
        whatsappGroom={whatsappGroom}
        whatsappBride={whatsappBride}
        nameGroom={nameGroom}
        nameBride={nameBride}
      />
      <WeddingList
        title="Lista de Presentes do Casamento"
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
