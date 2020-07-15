var redisPool = require('redis-connection-pool')('myRedisPool', {
    host: '127.0.0.1', // default
    port: 6379, //default
    // optionally specify full redis url, overrides host + port properties
    // url: "redis://username:password@host:port"
    max_clients: 30, // defalut
    perform_checks: false, // checks for needed push/pop functionality
    database: 0, // database number to use
    options: {
      auth_pass: ''
    } //options for createClient of node-redis, optional
  });
 
redisPool.set('test-key', 'foobar', function (err) {
  redisPool.get('test-key', function (err, reply) {
    console.log(reply); // 'foobar'

    redisPool.del('test-key', null);
  });
});
