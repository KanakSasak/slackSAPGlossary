# **SlackSAPGlossary**
[![made-with-Go](https://img.shields.io/badge/Made%20with-Go-1f425f.svg)](http://golang.org)

###### **_SAP Knowledge Base Bot Slack_**

### step step for preparing slack :

1. Create slack apps
2. Active slack with socket mode (no need endpoint when this method active)
3. Add scope Oauth & Permission (BOT Scope):
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
1. Open GCP console
2. Open API & Services -> Library
3. Search **_Google Sheets API_** and enable
4. Open API & Services -> Credentials
5. Create Credentials -> Service Account (fill necessary data)
6. Open IAM & Admin -> Service Accounts
7. Search your service account click on 3 dots on the right to get option and open Manage Key
8. Add your json key and download.
9. Back to API & Services scroll to bottom (Service Accounts)
10. Copy your service account email
11. Create Gsheet document
12. Add your service account email as an editor at gsheet doc
