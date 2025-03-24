import React from 'react';
import { Button, Col, Modal, Row } from 'react-bootstrap';

function WeddingModal({ showModal, handleCloseModal, selectedItem, whatsappGroom, whatsappBride, nameGroom, nameBride, cellPhoneGroom, cellPhoneBride }) {
  return (
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
  );
}

export default WeddingModal;
