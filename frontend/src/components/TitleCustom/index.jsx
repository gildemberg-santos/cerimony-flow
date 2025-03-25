import './styles.css';

function TitleCustom({ children }) {
  return (
    <h1 className="my-4 text-center text-title display-1 text-primary">{children}</h1>
  );
}

export default TitleCustom;
