# auto-newsletter-listmonk
Autonomous newsletter builder tool for Listmonk. This GoLang App compiles a newsletter from an RSS feed and posts it to Listmonk via API  to create the new campaign.

### How To Run This Software

Docker run

Docker Compose

Ansible

Github Actions

### Example Env Vars

```
# Listmonk Env Vars

LISTMONK_API=
USER=
PASSWORD=
FROM_EMAIL=
TEMPLATE_ID=
LISTS=

# RSS Feed Env Vars

RSS_FEED=
HEADER_MESSAGE=
FOOTER_MESSAGE=
```

### Docker CLI

```
docker run --rm \
  --name auto-listmonk \
  --env USER='magician' \
  --env PASSWORD='Magician.4sure' \
  --env LISTMONK_API='https://newsletter.opensourcegeeks.net/api/campaigns' \
  --env FROM_EMAIL='Chad at Opensource Geeks <chad@opensourcegeeks.net>' \
  --env TEMPLATE_ID='3' \
  --env LISTS='3' \
  --env RSS_FEED='https://opensourcegeeks.net/rss/' \
  --env HEADER_MESSAGE='Check Out This Week's Articles:' \
  --env FOOTER_MESSAGE='Thanks for being part of the Opensource Geeks community💻🐧' \
  auto-listmonk:latest
```

docker run --rm --name auto-listmonk --env-file .env auto-listmonk:latest