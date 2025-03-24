import React from 'react';
import { Card, Col, Container, Row } from 'react-bootstrap';
import usePhotos from '../../hooks/usePhotos';

function WeddingPhoto() {
  const photos = usePhotos();

  return (
    <Container>
      <Row>
        {photos && photos.map((photo, index) => (
          <Col md={4} key={index}>
            <Card className="mb-4 shadow-sm">
              <Card.Img variant="top" src={photo} alt={`Foto ${index + 1}`} />
            </Card>
          </Col>
        ))}
      </Row>
    </Container>
  );
}

export default WeddingPhoto;
