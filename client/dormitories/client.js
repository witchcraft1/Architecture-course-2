const http = require('../common/http');

const Client = (baseUrl) => {

  const client = http.Client(baseUrl);

  return {
    findDormitory: (specialty) => client.get('/getData', {specialty}),
    addStudent: (name, specialty, dormitoryId) => client.post('/postData', { name, specialty, dormitoryId })
  }

};

module.exports = { Client };