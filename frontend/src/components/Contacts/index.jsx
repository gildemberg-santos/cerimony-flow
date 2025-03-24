import React from 'react';
import { Button } from 'react-bootstrap';

function Contacts({ whatsappGroom, whatsappBride, nameGroom, nameBride }) {
  return (
    <div className="bg-light mb-4 p-4 d-flex flex-column align-items-center">
      <h2 className="my-4 body-text text-center font-cursive">Contatos</h2>
      <Button variant="outline-olive" href={`${whatsappGroom}`} target="_blank" rel="noopener noreferrer" className="mb-2">WhatsApp ({nameGroom})</Button>
      <Button variant="outline-olive" href={`${whatsappBride}`} target="_blank" rel="noopener noreferrer">WhatsApp ({nameBride})</Button>
    </div>
  );
}

export default Contacts;
