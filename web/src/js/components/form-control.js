import { createElementObj } from '../utils/create';
{
  /* <form class="row g-3">
  <div class="col-auto">
    <label for="staticEmail2" class="visually-hidden">Email</label>
    <input type="text" readonly class="form-control-plaintext" id="staticEmail2" value="email@example.com">
  </div>
  <div class="col-auto">
    <label for="inputPassword2" class="visually-hidden">Password</label>
    <input type="password" class="form-control" id="inputPassword2" placeholder="Password">
  </div>
  <div class="col-auto">
    <button type="submit" class="btn btn-primary mb-3">Confirm identity</button>
  </div>
</form> */
}

function Form(onSubmit, ...children) {
  const formS = createElementObj({
    tagName: 'div',
    classNames: 'row g-3',
    children: children,
    // attrs: [['onsubmit', onSubmit]]
  });

  return {
    init: () => {},
    get: () => {
      return formS;
    },
  };
}

{
  /* <div class="input-group input-group-sm mb-3">
  <span class="input-group-text" id="inputGroup-sizing-sm">Small</span>
  <input type="text" class="form-control" aria-label="Sizing example input" aria-describedby="inputGroup-sizing-sm">
</div> */
}

function Input(type = 'text') {
  let value = '';
  const handleChange = (e) => {
    value = e.target.value;
  };

  const i = createElementObj({
    tagName: 'input',
    classNames: 'form-control',
    onChange: handleChange,
    attrs: [['type', type]],
  });

  const div = createElementObj({
    tagName: 'div',
    classNames: 'input-group input-group-sm mb-3',
    children: [i],
  });

  return {
    init: () => {},
    get: () => {
      return div;
    },
    getValue: () => {
      return value;
    },
  };
}

export { Input, Form };
