import React from 'react';
import useCountdown from '../../hooks/useCountdown';
import SubTitleCustom from '../SubTitleCustom';

function formatDate(dateString) {
  const options = { weekday: 'long', year: 'numeric', month: 'long', day: 'numeric', hour: 'numeric', minute: 'numeric' };
  return new Date(dateString).toLocaleDateString('pt-BR', options);
}

function WeddingDate({ title, date }) {
  const timeLeft = useCountdown(date);

  const timerComponents = [];

  Object.keys(timeLeft).forEach(interval => {
    if (!timeLeft[interval]) {
      return;
    }

    timerComponents.push(
      <div key={interval} className="border border-light p-4 d-inline-block mx-2">
        <div className="d-flex flex-column align-items-center">
          <span className="timer-value">{timeLeft[interval]}</span>
          <span className="timer-label">{interval}</span>
        </div>
      </div>
    );
  });

  return (
    <div className="bg-primary p-4 pb-5">
      <SubTitleCustom className="text-white">{title}</SubTitleCustom>
      <p className="text-center font-roboto text-light fs-5">{formatDate(date)}</p>
      <div className="text-center font-roboto text-light">
        {timerComponents.length ? timerComponents : <span>Já chegou!</span>}
      </div>
    </div>
  );
}

export default WeddingDate;
