import React from 'react';
import { Card, Col, Row } from 'react-bootstrap';
import usePhotos from '../../hooks/usePhotos';
import SubTitleCustom from '../SubTitleCustom';

function WeddingPhoto({title}) {
  const photos = usePhotos();

  return (
    <div className="mb-4 p-4">
      <SubTitleCustom>{title}</SubTitleCustom>
      <Row className="justify-content-center">
        {photos && photos.map((photo, index) => (
          <Col key={index} xs={12} sm={6} md={4} lg={3} className="mb-4">
            <Card className="shadow-sm border-0">
              <Card.Img variant="top" src={photo} alt={`Foto ${index + 1}`} />
            </Card>
          </Col>
        ))}
      </Row>
    </div>
  );
}

export default WeddingPhoto;
