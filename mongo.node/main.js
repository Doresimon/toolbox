const MongoClient = require('mongodb').MongoClient;
const assert = require('assert');
 
// Connection URL
const url = 'mongodb://localhost:27017';
 
// Database Name
const dbName = 'halo';
 
// Use connect method to connect to the server
MongoClient.connect(url, function(err, client) {
  assert.equal(null, err);
  console.log("Connected successfully to server");
 
  const db = client.db(dbName);

  const collection = db.collection('docs');

  collection.insertMany([
    {a : {
        x:'xxx',
        y:[1,2,'ssss']
    }}, {a : 2}, {a : 3}
  ], function(err, result) {
    console.log("Inserted 3 documents into the collection");
    console.log(result);
  });

  collection.find({}).toArray(function(err, docs) {
    assert.equal(err, null);
    console.log("Found the following records");
    console.log(docs)
  });
 
  client.close();
});