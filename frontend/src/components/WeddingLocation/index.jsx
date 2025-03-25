import React, { useEffect, useState } from 'react';
import ButtonCustom from '../ButtonCustom';
import SubTitleCustom from '../SubTitleCustom';

function WeddingLocation({title, src}) {
  const [isShareSupported, setIsShareSupported] = useState(false);

  useEffect(() => {
    if (navigator.share) {
      setIsShareSupported(true);
    }
  }, []);

  const handleShare = async () => {
    if (navigator.share) {
      try {
        await navigator.share({
          title: 'Local do Casamento',
          url: window.location.href,
        });
        console.log('Compartilhamento bem-sucedido');
      } catch (error) {
        console.error('Erro ao compartilhar', error);
      }
    } else {
      console.log('Compartilhamento não suportado neste navegador');
    }
  };

  return (
    <div className="mb-4">
      <SubTitleCustom>{title}</SubTitleCustom>
      <div className="text-center font-roboto mb-4 fs-5">
        R. Sabiaguaba, 1270 - Edson Queiroz, Fortaleza - CE, 60835-750
      </div>
      <div className="map-container mb-4">
        <iframe
          src={src}
          className="shadow-sm"
          style={{ border: 0, width: "100%", height: "400px" }}
          allowFullScreen=""
          loading="lazy"
          title={title}
        ></iframe>
      </div>
      {isShareSupported && (
        <div className="text-center">
          <ButtonCustom onClick={handleShare} className="btn-primary">Compartilhar</ButtonCustom>
        </div>
      )}
    </div>
  );
}

export default WeddingLocation;
