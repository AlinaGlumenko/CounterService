var express = require('express');
var app = express();

var path = require('path')
app.use(express.static(__dirname + '/static/'));
app.set('views', path.join(__dirname, '/static/'));
app.engine('html', require('ejs').renderFile);
app.set('view engine', 'html');

app.get('/', function (req, res) {
  res.render("index");
});

app.get('/catpage', function (req, res) {
  res.render("catpage");
});

app.get('/unicornpage', function (req, res) {
  res.render("unicornpage");
});

app.listen(3000, function () {
  console.log('App listening on port 3000!');
});