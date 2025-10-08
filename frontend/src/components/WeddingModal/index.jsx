import { Col, Modal, Row } from 'react-bootstrap';
import '../../index.css';
import ButtonCustom from '../ButtonCustom';
import './styles.css';

function WeddingModal({ showModal, handleCloseModal, selectedItem, whatsappGroom, whatsappBride, nameGroom, nameBride, cellPhoneGroom, cellPhoneBride }) {
  return (
    <Modal show={showModal} onHide={handleCloseModal} size="xl" centered>
      <Modal.Header className="text-white bg-primary">
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
            <p className={`mt-2 ${selectedItem?.available ? 'text-success' : 'text-danger'} fs-6`}>
              {selectedItem?.available ? "Disponível" : "Indisponível"}
            </p>
            <div className="mt-4 d-flex flex-column align-items-start">
              {selectedItem?.link && (
                <ButtonCustom href={selectedItem?.link} className="mb-2">Ver Produto na Loja</ButtonCustom>
              )}
              <ButtonCustom href={`${whatsappGroom}`} className="mb-2">WhatsApp ({nameGroom})</ButtonCustom>
              <ButtonCustom href={`${whatsappBride}`}>WhatsApp ({nameBride})</ButtonCustom>
            </div>
            <p className="mt-4 font-size-0-9em">
              ou o valor do presente pode ser transferido para a conta dos noivos via Pix
            </p>
            <div className="mt-4 d-flex flex-column">
              <h5 className="fs-5 fw-bold text-primary">Dados para transferência</h5>
              <div className="mt-2 fs-6 text-secondary">
                <strong>Pix ({nameGroom}):</strong> {cellPhoneGroom}
              </div>
              <div className="fs-6 text-secondary">
                <strong>Pix ({nameBride}):</strong> {cellPhoneBride}
              </div>
            </div>
          </Col>
        </Row>
      </Modal.Body>
      <Modal.Footer className="text-white bg-primary">
        <ButtonCustom onClick={handleCloseModal}>Fechar</ButtonCustom>
      </Modal.Footer>
    </Modal>
  );
}

export default WeddingModal;
