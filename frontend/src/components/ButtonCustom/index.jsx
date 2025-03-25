import { Button } from 'react-bootstrap';
import './styles.css';

function ButtonCustom({href, className, children, onClick}) {
  return (
    <Button href={href} target="_blank" rel="noopener noreferrer" className={`${className} button-custom`} onClick={onClick}>{children}</Button>
  );
}

export default ButtonCustom;
