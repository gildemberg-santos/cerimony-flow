import React from 'react';
import ButtonCustom from '../ButtonCustom';
import SubTitleCustom from '../SubTitleCustom';

function Contacts({ title, whatsappGroom, whatsappBride, nameGroom, nameBride }) {
  return (
    <div className="mb-4 p-4 d-flex flex-column align-items-center">
      <SubTitleCustom className="text-primary">{title}</SubTitleCustom>
      <ButtonCustom href={whatsappGroom} className="mb-2">WhatsApp ({nameGroom})</ButtonCustom>
      <ButtonCustom href={whatsappBride}>WhatsApp ({nameBride})</ButtonCustom>
    </div>
  );
}

export default Contacts;
