import { Card, Col, Row } from 'react-bootstrap';
import ButtonCustom from '../ButtonCustom';
import SubTitleCustom from '../SubTitleCustom';

function WeddingList({ title, weddingList, handleShowModal }) {
  return (
    <div className="mb-4 p-4">
      <SubTitleCustom className="text-primary">{title}</SubTitleCustom>
      <Row className="justify-content-center">
        {weddingList.filter(item => item.available).map((item, index) => (
          <Col key={index} xs={12} sm={6} md={4} lg={3} className="mb-4">
            <Card className="shadow-sm border-0 p-4">
              <div className="d-flex justify-content-center align-items-center overflow-hidden p-4" style={{ width: '200px', height: '200px', margin: '0 auto' }}>
                <Card.Img
                  variant="top"
                  src={item.picture}
                  alt={item.title}
                  className="p-4 rounded img-fluid"
                  style={{ objectFit: 'contain' }}
                />
              </div>
              <Card.Body className="text-center">
                <Card.Title className="text-olive">{item.title}</Card.Title>
                <ButtonCustom onClick={() => handleShowModal(item)} className="btn btn-primary">Detalhes</ButtonCustom>
              </Card.Body>
            </Card>
          </Col>
        ))}
      </Row>
    </div>
  );
}

export default WeddingList;
