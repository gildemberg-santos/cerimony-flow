import React, { useEffect, useState } from 'react';

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
      <h2 className="my-4 wedding-text text-center font-cursive">{title}</h2>
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
          <button onClick={handleShare} className="btn btn-primary">Compartilhar</button>
        </div>
      )}
    </div>
  );
}

export default WeddingLocation;
