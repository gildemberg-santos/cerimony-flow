import React from 'react';
import brasao from '../../brasao.png';
import background from '../../cerimonia.jpeg';
import TitleCustom from '../TitleCustom';
import './styles.css';

function Header({ title, description }) {
  return (
    <div className="mb-4 p-4">
      <div className="crest text-center mb-2">
        <img src={brasao} alt="Brasão" width={200} height={200} />
      </div>
      <TitleCustom>{title}</TitleCustom>
      <div className="image-container mb-4">
        <img src={background} alt="Cerimônia" className="img-fluid w-100 cropped-image filter-image" />
      </div>
      <p className="text-center font-roboto font-size-1-2em">{description}</p>
    </div>
  );
}

export default Header;
