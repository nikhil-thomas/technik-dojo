# Telegram Bot Example
Credit: Soham Kamani Blog [How to make a responsive telegram bot](https://www.sohamkamani.com/blog/2016/09/21/making-a-telegram-bot/)

## Deployment steps

1. create telegram bot and get an api-key
2. add telegram api-key to app
3. deploy this app with a public domain name
4. add telegram web-hook using the doamin name
   curl -F "url=<app-domain-name>/new-message" https://api.telegram.org/bot<telegram-bot-api-key>/setWebhook

to deploy the app in public domain:
  - use [now](https://zeit.co/now)
