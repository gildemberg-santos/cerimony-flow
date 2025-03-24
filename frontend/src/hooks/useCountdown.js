import { useEffect, useState } from 'react';

function calculateTimeLeft(date) {
  const difference = +new Date(date) - +new Date();
  let timeLeft = {};

  if (difference > 0) {
    timeLeft = {
      dias: Math.floor(difference / (1000 * 60 * 60 * 24)),
      horas: Math.floor((difference / (1000 * 60 * 60)) % 24),
      minutos: Math.floor((difference / 1000 / 60) % 60)
    };
  };

  return timeLeft;
}

function useCountdown(date) {
  const [timeLeft, setTimeLeft] = useState(calculateTimeLeft(date));

  useEffect(() => {
    const timer = setInterval(() => {
      setTimeLeft(calculateTimeLeft(date));
    }, 30000);

    return () => clearInterval(timer);
  }, [date]);

  return timeLeft;
}

export default useCountdown;
