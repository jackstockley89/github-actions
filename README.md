# Github Actions

## About
--- 
`github-actions` repository is a collection of Go scripts used to preform tasks within a CI/CD pipeline 

## Language Version
---
![GitHub go.mod Go version (subdirectory of monorepo)](https://img.shields.io/github/go-mod/go-version/jackstockley89/github-actions?filename=go.mod&style=for-the-badge)


## Workflow Activity
![GitHub Workflow Status](https://img.shields.io/github/actions/workflow/status/jackstockley89/github-actions/go-releaser.yaml?label=Release%20Build&style=for-the-badge)
![GitHub Workflow Status](https://img.shields.io/github/actions/workflow/status/jackstockley89/github-actions/nightly-build.yaml?label=Nightly%20Build&style=for-the-badge)


## Script Index
--- 

### GITHUB

![Docker Pulls](https://img.shields.io/docker/pulls/jackstock8904/request-review?label=request-review%20docker%20pulls%20&style=for-the-badge)

When called this action will assign a github user to a pull request if they are a Collaborator

Directory Link: [request-review](https://github.com/jackstockley89/github-actions/tree/main/request-review)

Workflow Example: [README](ttps://github.com/jackstockley89/github-actions/github/tree/main/request-review/README.md)

---
![Docker Pulls](https://img.shields.io/docker/pulls/jackstock8904/auth-check?label=auth-check%20docker%20pulls%20&style=for-the-badge)

When called this action will compare the user who raised the pull request with the Collaborator list

Directory Link: [auth-check](ttps://github.com/jackstockley89/github-actions/github/tree/main/auth-check)

Workflow Example: [README](ttps://github.com/jackstockley89/github-actions/github/tree/main/auth-check/README.md)

---
![Docker Pulls](https://img.shields.io/docker/pulls/jackstock8904/close-pr?label=close-pr%20docker%20pulls%20&style=for-the-badge)

When called this action will close a pull request 

Directory Link: [close-pr](ttps://github.com/jackstockley89/github-actions/github/tree/main/close-pr)

Workflow Example: [README](ttps://github.com/jackstockley89/github-actions/github/tree/main/close-pr/README.md)

---
![Docker Pulls](https://img.shields.io/docker/pulls/jackstock8904/version-check?label=version-check%20docker%20pulls&style=for-the-badge)

When called this action will check repository for the latest release tag

Directory Link: [version-check](ttps://github.com/jackstockley89/github-actions/github/tree/main/version-check)

---

## How to publish a new binary
---
Within this respoistory, the goreleaser tool is used to create a Go binary and push a image to Docker Hub. To publish a new binary once the Pull Request is approved and merged into the main branch, create a new release and the github action will automatically start and publish the new binary. No changes will be needed to the workflow that is calling the image aslong as the latest image is being pulled