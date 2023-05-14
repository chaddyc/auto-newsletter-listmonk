# auto-newsletter-listmonk
Autonomous newsletter builder tool for Listmonk. This GoLang App compiles a newsletter from an RSS feed and posts it to Listmonk via API  to create the new campaign.

### How To Run This Software

Docker CLI

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



docker run --rm --name auto-listmonk --env-file .env auto-listmonk:latest