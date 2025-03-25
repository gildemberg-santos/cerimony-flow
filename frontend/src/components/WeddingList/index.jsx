import React from 'react';
import { Card, Col, Row } from 'react-bootstrap';
import ButtonCustom from '../ButtonCustom';
import SubTitleCustom from '../SubTitleCustom';

function WeddingList({ title, weddingList, handleShowModal }) {
  return (
    <div className="mb-4 p-4">
      <SubTitleCustom>{title}</SubTitleCustom>
      <Row className="justify-content-center">
        {weddingList.map((item, index) => (
          <Col key={index} xs={12} sm={6} md={4} lg={3} className="mb-4">
            <Card className="shadow-sm border-0">
              <Card.Img variant="top" src={item.picture} alt={item.title} className="p-4 rounded" />
              <Card.Body className="text-center">
                <Card.Title className="font-roboto font-size-1-1em text-olive">{item.title}</Card.Title>
                <ButtonCustom onClick={() => handleShowModal(item)}>Detalhes</ButtonCustom>
                <p className={`mt-2 ${item.available ? 'text-success' : 'text-danger'} fs-6`}>
                  {item.available ? "Disponível" : "Indisponível"}
                </p>
              </Card.Body>
            </Card>
          </Col>
        ))}
      </Row>
    </div>
  );
}

export default WeddingList;
