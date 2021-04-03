# Discord Join Page

This is a small Go program to serve a webpage that will generate on-demand one use Discord invites.
It is inspired on the many Slack invite pages.

The goal is to add an hCaptcha challange, Discord currently has a mix of hCapthca and reCaptcha themselves. This was designed to add an extra boundry before getting an invite link.

## Configuration
All configuration is done via envvvars
```
# A Discord Bot token
DISCORDJOINPAGE_TOKEN

# Channel ID in the guild to join
DISCORDJOINPAGE_CHANNEL_ID

# hCaptcha keys
DISCORDJOINPAGE_HCAPTCHA_SITE_KEY
DISCORDJOINPAGE_HCAPTCHA_SITE_SECRET
```

# Using the official Docker image

We provide a prebuilt Docker image to use, you might however want to overwrite our WWW directory.

```bash
docker run \
  -e DISCORDJOINPAGE_TOKEN=<token> \
  -e DISCORDJOINPAGE_CHANNEL_ID=<id> \
  -e DISCORDJOINPAGE_HCAPTCHA_SITE_KEY=<site key> \
  -e DISCORDJOINPAGE_HCAPTCHA_SITE_SECRET=<site secret> \
  -v www:/go/src/github.com/meyskens/discord-join-page/www \ # override the website with your own template 
  ghcr.io/meyskens/discord-invite-page:v0.1.0
```

## Real world examples
* [IT Factory Discord](https://discord.itf.to)