import React from 'react';

function WeddingLocation({title, src}) {
  return (
    <div className="bg-light mb-4">
      <h2 className="my-4 body-text text-center font-cursive text-light">{title}</h2>
      <div className="map-container mb-4">
        <iframe
          src={src}
          className="shadow-sm"
          style={{ border: 0, margin: 0, width: "100%", height: "400px" }}
          allowFullScreen=""
          loading="lazy"
          title={title}
        ></iframe>
      </div>
    </div>
  );
}

export default WeddingLocation;
