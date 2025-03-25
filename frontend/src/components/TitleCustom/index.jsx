import './styles.css';

function TitleCustom({ children, className }) {
  return (
    <h1 className={`my-4 text-center text-title display-1 text-primary ${className}`}>{children}</h1>
  );
}

export default TitleCustom;
