import { createElementObj } from '../utils/create';

const Button = (type = '', text = '', onClick = () => {}) => {
  const button = createElementObj({
    tagName: 'button',
    classNames: 'btn button',
    textContent: text,
    attrs: [['type', type]],
    onClick: onClick,
  });

  return {
    get: () => {
      return button;
    },
  };
};

export default Button;
