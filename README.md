Newsletter builder/compiler software for Ghost Blog CMS and Listmonk Newsletter Manager.

## Table of Contents

- [About](#about)
- [Getting Started](#how_to_run_this_software)
- [Contributors](#contributors)
- [How To Contribute](#how_to_contribute)
- [License](#license)

## About

This is a newsletter builder tool for [Ghost CMS](https://opensourcegeeks.net/how-to-setup-ghost-blog-cms-with-docker/) and [Listmonk](https://opensourcegeeks.net/how-to-install-listmonk-with-docker/). This software allows you to compile a newsletter via Ghost CMS RSS Feed and creates your newsletter campaign in Listmonk via API with one command and can be executed manually or autonomously via the methods below.

## How To Run This Software

This software tool can be utilized via the `Docker CLI`, `GitHub Actions` or by creating an `Ansible Playbook`.

### Docker CLI

Create a folder and call it `auto-listmonk-newsletter` and cd to the directory. Create a `.env` file and add the following vars in the `.env`, add your values and save it:

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

Run the following command to execute the docker container which will then build out the container and run the go binary once and remove the container once done.

```
docker run --rm --name auto-listmonk --env-file .env chaddyc/auto-listmonk-newsletter:latest
```

To automate this process you can set up a `cron` job or you can build an ansible playbook like in the next method.

### Ansible

### Github Actions

## Contributors

## How To Contribute

Contact me via [Twitter](https://twitter.com/fossgeek)

## License

For license information, please see [LICENSE](./LICENSE)