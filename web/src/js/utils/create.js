function createElementObj({
  tagName,
  classNames,
  children,
  attrs,
  textContent,
  onChange,
  onClick,
}) {
  const element = document.createElement(tagName);
  if (classNames) {
    element.classList.add(...classNames.split(' '));
  }
  if (children && children.length) {
    children.forEach((child) => {
      element.append(child);
    });
  }
  if (attrs && attrs.length) {
    attrs.forEach(([attrName, attrVal]) => {
      element.setAttribute(attrName, attrVal);
    });
  }
  if (textContent) {
    element.innerHTML = textContent;
  }

  if (onChange) {
    element.onchange = onChange;
  }

  if (onClick) {
    element.onclick = onClick;
  }

  return element;
}

function createElement(tagName, classNames, children, attrs, textContent) {
  return createElementObj({ tagName, classNames, children, attrs, textContent });
}

export { createElement, createElementObj };
