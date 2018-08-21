var express = require('express');
var app = express();
var bodyParser = require('body-parser');
const axios = require('axios');

app.use(bodyParser.json())
app.use(
    bodyParser.urlencoded({
        extended: true
    })
);

app.post('/new-message', function(req, res) {
    const message = req.body.message;

    if (!message || (message.text && message.text.toLowerCase().indexOf('technik') < 0)) {
        return res.end()
    }

    // if message text contains the word 'technik' reply the word 'dojo'
    axios
        .post(
            'https://api.telegram.org/bot<telegram-bot-api-token>/sendMessage',
            {
                chat_id: message.chat.id,
                text: "dojo: " + message.text
            }
        )
        .then(response => {
            console.log('Success: Message Posted');
            res.end('ok');
        })
        .catch(err => {
            console.log('Error:', err);
            res.end('Error :' + err);
        })
});

app.listen(3000, function() {
    console.log('Telegram app listening on port 3000!');
});
