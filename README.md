# rssbot

Telegram bot that allows you to receive news from RSS channels.
You can use an existing bot - [t.me/forest33_rss_bot](https://t.me/forest33_rss_bot) or deploy your own.

### Installation

Create database and user
```
CREATE DATABASE rssbot;
CREATE USER rssbot WITH LOGIN PASSWORD 'change-me';
GRANT ALL ON DATABASE rssbot TO rssbot;
```

```
git clone https://github.com/Forest33/rssbot.git
cd rssbot
```

Edit the config file config/rssbot.json

```
docker-compose up -d
```
