# **SlackSAPGlossary**
[![made-with-Go](https://img.shields.io/badge/Made%20with-Go-1f425f.svg)](http://golang.org)

###### **_SAP Knowledge Base Bot Slack_**

### step step for preparing slack :

1. create slack apps
2. active slack with socket mode (no need endpoint when this method active)
3. add scope Oauth & Permission (BOT Scope):
   - app_mentions:read
   - channels:history
   - chat:write
   - commands
   - groups:history
   - im:history
   - mpim:history
   - reactions:write
4. Event and subscription (Subscribe to bot events)
   - app_mention
   - message.channels
   - message.groups
   - message.im
   - message.mpim
5. Get app token(xoxp***) and bot token (xoxb***)

### step step for preparing get GCP json credential :
1. open GCP console
2. open API & Services -> Library
3. search **_Google Sheets API_** and enable
4. open API & Services -> Credentials
5. Create Credentials -> Service Account (fill necessary data)
6. back to API & Services scroll to bottom (Service Accounts)
7. copy your service account email
8. create Gsheet document
9. add your service account email as an editor at gsheet doc
