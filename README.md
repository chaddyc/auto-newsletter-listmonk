Newsletter builder/compiler software for Ghost Blog CMS and Listmonk Newsletter Manager. This [Blog Article](https://opensourcegeeks.net/how-i-built-an-autonomous-newsletter-compiler-tool-with-golang-github-actions-and-listmonk/) display the steps how this software was developed and examples of how it works.

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

### Docker CLI:

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

### Ansible:

In progress...

### Github Actions:

The first step is to create a [GitHub Secret](https://docs.github.com/en/actions/security-guides/encrypted-secrets) in the repository that you are going to utilize to run a workflow for this tool. Call the secret `ENV_FILE` and add your env vars values and save the secret. Add the following workflow below under `.github/workflows/campaign.yml`.

`ENV_FILE` secret template to be added under Github Actions secrets for your repository.(See `sample.env`)

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

This workflow should be executed manually in the `Github Actions` tab whenever you want to compile your latest 5 articles links into a newsletter campaign for Listmonk.

```
name: Create Listmonk Campaign

on:
  workflow_dispatch

jobs:
  campaign:
    runs-on: ubuntu-latest
    steps:
    - name: Create .env
      run: echo "${{ secrets.ENV_FILE }}" > .env

    - name: Docker run
      run: docker run --rm --name auto-listmonk --env-file .env chaddyc/auto-listmonk-newsletter:latest
```

This workflow below will execute every Sunday and create a campaign in Listmonk with the latest 5 articles.

```
name: Create Listmonk Campaign

on:
  schedule:
    # Runs every Sunday at 18:30
    - cron: '5 19 * * Sun'

jobs:
  campaign:
    runs-on: ubuntu-latest
    steps:
    - name: Create .env
      run: echo "${{ secrets.ENV_FILE }}" > .env

    - name: Docker run
      run: docker run --rm --name auto-listmonk --env-file .env chaddyc/auto-listmonk-newsletter:latest
```

## Contributors

## How To Contribute

Contact me via [Twitter](https://twitter.com/fossgeek)

## License

For license information, please see [LICENSE](https://github.com/chaddyc/auto-newsletter-listmonk/blob/main/LICENSE)