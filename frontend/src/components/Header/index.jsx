import React from 'react';

function Header({ title, description }) {
  return (
    <div className="bg-light mb-4 p-4">
      <h1 className="my-4 wedding-text text-center font-cursive">{title}</h1>
      <p className="body-text text-center font-roboto font-size-1-2em">{description}</p>
    </div>
  );
}

export default Header;
