import React from 'react';
import { Button, Card, Col, Row } from 'react-bootstrap';

function WeddingList({ weddingList, handleShowModal }) {
  return (
    <div className="mb-4 p-4">
      <h2 className="my-4 wedding-text text-center font-cursive">Lista de Casamento</h2>
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
  );
}

export default WeddingList;
