import React from 'react';
import useCountdown from '../../hooks/useCountdown';

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
    <div className="bg-olive p-4">
      <h2 className="wedding-text text-center font-cursive text-light">{title}</h2>
      <p className="body-text text-center font-roboto font-size-1-2em text-light">{formatDate(date)}</p>
      <div className="body-text text-center font-roboto font-size-1-2em text-light">
        {timerComponents.length ? timerComponents : <span>Já chegou!</span>}
      </div>
    </div>
  );
}

export default WeddingDate;
