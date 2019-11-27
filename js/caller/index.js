const express = require('express');
const got = require('got');
const app = express();

app.get('/js', (req, res) => {

  got('https://jsapi-zlg45lmdqa-ew.a.run.app/').then(r => {
    res.send(`From JS API: Hello ${r.body}!`);
  } 
  )
  
});

app.get('/go', (req, res) => {

  got('https://projectemitter-zlg45lmdqa-ew.a.run.app/').then(r => {
    res.send(`From GO API: Hello ${r.body}!`);
  } 
  )
  
});

const port = process.env.PORT || 8080;
app.listen(port, () => {
  console.log('Hello world listening on port', port);
});
