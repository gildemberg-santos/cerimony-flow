import './styles.css';

function SubTitleCustom({ children, className }) {
  return (
    <h2 className={`my-4 text-center text-title display-5 text-primary ${className}`}>{children}</h2>
  );
}

export default SubTitleCustom;
