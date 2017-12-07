# commango-slackbot

This is slack-bot which uses unix command in slack.
To run this bot, you need to create a bot as [slack app](https://api.slack.com/slack-apps).
Create a new app from [here](https://api.slack.com/apps).

## Usage

To run this bot you need to a Docker, in your environment.

1. edit config.*.toml (development/production)
2. edit shell file
	- docker-build.sh (tag)
	- docker-push.sh (default: public repository on docker hub)
	- docker-run.sh (volume mount the directory you want to execute the command)
3. make dev(pro)
4. make docker-run

### Dependencies

Dependencies are managed by [dep](https://github.com/golang/dep)