# Golang Slackbot

### Slack-bot. Golang. Nlopes/slack. Say "Hello".

## Step 1

### Install
For the Slack API Client use:
```
go get github.com/nlopes/slack
```

## Step 2

Slack has two different kinds of bot users: app bots and custom bots.
Recommended sticking with app bots.
#### The first way
In order to create Bot as a SlackApp:
https://api.slack.com/apps
### OR
#### Second way
In order to create custom bots:
Head to https://$YOUR_ORG.slack.com/apps/A0F7YS25R-bots to get to the Bots app page and get Slack setup properly.

## Step 3

In the token.json, replace "YOUR_TOKEN" with your token "xxxx-.......-..............."

## Step 4

```
go run main.go
```

## Step 5
Write to your bot "@your_bot hello" in Slack.
If you write to bot in private messages you don`t need to use "@your_bot".
Just write to him and he will understand you! :)

Also possible to use "hi", "Hi", "hello", "Hello", "hey", "Hey"