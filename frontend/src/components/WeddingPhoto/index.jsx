import { useState } from 'react';
import { Card, Col, Row } from 'react-bootstrap';
import Skeleton from 'react-loading-skeleton';
import 'react-loading-skeleton/dist/skeleton.css';
import Lightbox from 'yet-another-react-lightbox';
import 'yet-another-react-lightbox/styles.css';
import usePhotos from '../../hooks/usePhotos';
import SubTitleCustom from '../SubTitleCustom';
import './styles.css';

function WeddingPhoto({ title }) {
  const photos = usePhotos();
  const [loadingStates, setLoadingStates] = useState(Array(photos?.length || 0).fill(true));
  const [isOpen, setIsOpen] = useState(false);
  const [photoIndex, setPhotoIndex] = useState(0);

  const handleImageLoad = (idx) => {
    setLoadingStates((prev) => {
      const updated = [...prev];
      updated[idx] = false;
      return updated;
    });
  };

  const handleOpenLightbox = (idx) => {
    setPhotoIndex(idx);
    setIsOpen(true);
  };

  return (
    <div className="mb-4 p-4">
      <SubTitleCustom className="text-primary">{title}</SubTitleCustom>
      <Row className="justify-content-center">
        {photos && photos.map((photo, index) => (
          <Col key={index} xs={12} sm={6} md={4} lg={3} className="mb-4">
            <Card className="shadow-sm border-0 d-flex align-items-center justify-content-center" style={{ minHeight: '220px' }}>
              {loadingStates[index] ? (
                <Skeleton height={200} width={200} style={{ borderRadius: '12px' }} />
              ) : null}
              <Card.Img
                variant="top"
                src={photo}
                alt={`Foto ${index + 1}`}
                style={loadingStates[index] ? { display: 'none' } : { maxHeight: '200px', maxWidth: '200px', objectFit: 'cover', borderRadius: '12px', cursor: 'pointer' }}
                onLoad={() => handleImageLoad(index)}
                onClick={() => handleOpenLightbox(index)}
              />
            </Card>
          </Col>
        ))}
      </Row>
      {isOpen && (
        <Lightbox
          open={isOpen}
          close={() => setIsOpen(false)}
          slides={photos.map((src, idx) => ({ src, alt: `Foto ${idx + 1}` }))}
          index={photoIndex}
          on={{
            view: ({ index }) => setPhotoIndex(index)
          }}
        />
      )}
    </div>
  );
}

export default WeddingPhoto;
