const channels = require('./dormitories/client');

const client = channels.Client('http://localhost:8080');

//Scenario 1: Display available Dormitory.
client.findDormitory('biology')
  .then((list) => {
    console.log('----- Scenario 1 -----');
    console.log('Available dormitory:');
    console.log(`Dormitory ID: ${list.id}`);
    console.log(`Name: ${list.name}`);
    console.log("Specialities: ", list.studentsCount)
    console.log('----------------------');
  })
  .catch((e) => {
    console.log(`Problem listing available dormitory: ${e.message}`);
  });

//Scenario 2: Add Student to Dormitory
client.addStudent('Dan', 'biology', 3)
.then((resp) => {
  console.log('----- Scenario 2 -----');
  console.log('New student added: ');
  console.log(resp);
  console.log('----------------------');
})
.catch(e => {
  console.log(`Problem adding student: ${e.message}`);
});