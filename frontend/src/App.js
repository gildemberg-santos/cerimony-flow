import axios from 'axios';
import 'bootstrap/dist/css/bootstrap.min.css';
import React, { useEffect, useState } from 'react';
import { Button, Card, Col, Container, Modal, Row } from 'react-bootstrap';
import './index.css'; // Import the CSS file

function App() {
  const [title, setTitle] = useState('');
  const [description, setDescription] = useState('');
  const [whatsappGroom, setWhatsappGroom] = useState('');
  const [whatsappBride, setWhatsappBride] = useState('');
  const [nameGroom, setNameGroom] = useState('');
  const [nameBride, setNameBride] = useState('');
  const [cellPhoneGroom, setCellPhoneGroom] = useState('');
  const [cellPhoneBride, setCellPhoneBride] = useState('');
  const [weddingList, setWeddingList] = useState([]);
  const [selectedItem, setSelectedItem] = useState(null);
  const [showModal, setShowModal] = useState(false);

  const loadMessage = () => {
    const url = `${process.env.REACT_APP_API_BASE || ''}/hello`;
    axios.get(url, {
      headers: {
        'Content-Type': 'application/json',
      }
    }).then(({ data }) => {
      const { title, description, whatsapp_groom, whatsapp_bride, name_groom, name_bride, cell_phone_groom, cell_phone_bride } = data;
      setTitle(title);
      setDescription(description);
      setWhatsappGroom(whatsapp_groom);
      setWhatsappBride(whatsapp_bride);
      setNameGroom(name_groom);
      setNameBride(name_bride);
      setCellPhoneGroom(cell_phone_groom);
      setCellPhoneBride(cell_phone_bride);
    })
      .catch(error => console.error(error));
  };

  const loadWeddingList = () => {
    const url = `${process.env.REACT_APP_API_BASE || ''}/wedding-list`;
    axios.get(url, {
      headers: {
        'Content-Type': 'application/json',
      }
    }).then(({ data }) => setWeddingList(data.wedding_list))
      .catch(error => console.error(error));
  };

  const handleShowModal = (item) => {
    setSelectedItem(item);
    setShowModal(true);
  };

  const handleCloseModal = () => {
    setShowModal(false);
    setSelectedItem(null);
  };

  useEffect(() => {
    loadMessage();
    loadWeddingList();
  }, []);

  return (
    <Container fluid className="bg-light text-light p-0 m-0">
      <div className="bg-light mb-4 p-4">
        <h1 className="my-4 wedding-text text-center font-cursive">{title}</h1>
        <p className="body-text text-center font-roboto font-size-1-2em">{description}</p>
      </div>

      <div className="bg-olive mb-4 p-4">
        <h2 className="my-4 body-text text-center font-cursive text-light">Data de Casamento</h2>
        <p className="body-text text-center font-roboto font-size-1-2em text-light">Dom, 09 de Novembro de 2025 às 16:00</p>
      </div>

      <div className="bg-light mb-4">
        <h2 className="my-4 body-text text-center font-cursive text-light">Localização do Casamento</h2>
        <div className="map-container mb-4">
          <iframe
            src="https://www.google.com/maps/embed?pb=!1m18!1m12!1m3!1d3153.019019726204!2d-122.419415484681!3d37.7749292797596!2m3!1f0!2f0!3f0!3m2!1i1024!2i768!4f13.1!3m3!1m2!1s0x8085809c5b0b0b0b%3A0x0!2zMzfCsDQ2JzI5LjgiTiAxMjLCsDI1JzEwLjkiVw!5e0!3m2!1sen!2sus!4v1633045678901!5m2!1sen!2sus"
            className="shadow-sm"
            style={{ border: 0, margin: 0, width: "100%", height: "400px" }}
            allowFullScreen=""
            loading="lazy"
            title="Localização do Casamento"
          ></iframe>
        </div>
      </div>

      <div className="bg-light mb-4 p-4 d-flex flex-column align-items-center">
        <h2 className="my-4 body-text text-center font-cursive">Contatos</h2>
        <Button variant="outline-olive" href={`${whatsappGroom}`} target="_blank" rel="noopener noreferrer" className="mb-2">WhatsApp ({nameGroom})</Button>
        <Button variant="outline-olive" href={`${whatsappBride}`} target="_blank" rel="noopener noreferrer">WhatsApp ({nameBride})</Button>
      </div>

      <div className="bg-light mb-4 p-4">
        <h2 className="my-4 body-text text-center font-cursive">Lista de Casamento</h2>
        <Row className="justify-content-center">
          {weddingList.map((item, index) => (
            <Col key={index} xs={12} sm={6} md={4} lg={3} className="mb-4">
              <Card className="shadow-sm border-0">
                <Card.Img variant="top" src={item.picture} alt={item.title} className="p-4 rounded" />
                <Card.Body className="text-center">
                  <Card.Title className="body-text font-roboto font-size-1-1em text-olive">{item.title}</Card.Title>
                  <Button variant="outline-olive body-text" onClick={() => handleShowModal(item)}>Detalhes</Button>
                  <p className={`mt-2 body-text ${item.available ? 'text-success' : 'text-danger'} font-size-0-9em`}>
                    {item.available ? "Disponível" : "Indisponível"}
                  </p>
                </Card.Body>
              </Card>
            </Col>
          ))}
        </Row>
      </div>


      <Modal show={showModal} onHide={handleCloseModal} size="xl" centered>
        <Modal.Header className="text-white bg-olive">
          <Modal.Title>{selectedItem?.title}</Modal.Title>
        </Modal.Header>
        <Modal.Body className="bg-light">
          <Row>
            <Col xs={12} md={4}>
              <div className="modal-image-container p-2">
                <img src={selectedItem?.picture} alt={selectedItem?.title} className="img-fluid mb-3 modal-image rounded shadow-sm" />
              </div>
            </Col>
            <Col xs={12} md={8}>
              <p>{selectedItem?.description}</p>
              <p className={`mt-2 ${selectedItem?.available ? 'text-success' : 'text-danger'} font-size-0-9em`}>
                {selectedItem?.available ? "Disponível" : "Indisponível"}
              </p>
              <div className="mt-4 d-flex flex-column align-items-start">
                <Button variant="outline-olive" href={selectedItem?.link} target="_blank" rel="noopener noreferrer" className="mb-2">Ver Produto na Loja</Button>
                <Button variant="outline-olive" href={`${whatsappGroom}`} target="_blank" rel="noopener noreferrer" className="mb-2">WhatsApp ({nameGroom})</Button>
                <Button variant="outline-olive" href={`${whatsappBride}`} target="_blank" rel="noopener noreferrer">WhatsApp ({nameBride})</Button>
              </div>
              <p className="mt-4 body-text font-size-0-9em">
                ou o valor do presente pode ser transferido para a conta dos noivos via Pix
              </p>
              <div className="mt-4 d-flex flex-column">
                <h5 className="body-text font-cursive">Dados para transferência</h5>
                <div className="body-text mt-2">
                  <strong>Pix ({nameGroom}):</strong> {cellPhoneGroom}
                </div>
                <div className="body-text">
                  <strong>Pix ({nameBride}):</strong> {cellPhoneBride}
                </div>
              </div>
            </Col>
          </Row>
        </Modal.Body>
        <Modal.Footer className="text-white bg-olive">
          <Button variant="outline-light" onClick={handleCloseModal}>Fechar</Button>
        </Modal.Footer>
      </Modal>
    </Container>
  );
}

export default App;
