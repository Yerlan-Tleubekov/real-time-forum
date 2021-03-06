const instance = (url = '', method = '', data = {}, headers = {}, body) => {
  const b = {
    method, // *GET, POST, PUT, DELETE, etc.
    mode: 'cors', // no-cors, *cors, same-origin
    cache: 'no-cache', // *default, no-cache, reload, force-cache, only-if-cached
    credentials: 'same-origin', // include, *same-origin, omit
    headers: {
      'Content-Type': 'application/json',
      ...headers,
    },
    ...data,
  };

  if (body) {
    b.body = JSON.stringify(body);
  }

  return fetch(url, b);
};

export default instance;
