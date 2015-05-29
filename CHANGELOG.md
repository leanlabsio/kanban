# 1.0.11 (2015-05-29)

## Features:

- Board shows 100 issues. Closes [#27](https://gitlab.com/leanlabsio/kanban/issues/27).

## Bugfixes:

- Reload the page on websocket connection fail. Closes [#33](https://gitlab.com/leanlabsio/kanban/issues/33).

# 1.0.10 (2015-05-28)

## Features:

- Allow to change the title and description of an issue. Closes [#23](https://gitlab.com/leanlabsio/kanban/issues/23).
- Add links to GitLab issues from card view. Closes [#21](https://gitlab.com/leanlabsio/kanban/issues/21).

# 1.0.9 (2015-05-25)

## Features:

- Hide archived projects from boards list. Closes [#36](https://gitlab.com/leanlabsio/kanban/issues/36).

## Bugfixes:

- TODO can be added when issue description is empty. Closes [#29](https://gitlab.com/leanlabsio/kanban/issues/29).
- Fixed websocket connection issue, thanks to Markus Zimmermann. Closes [#14](https://gitlab.com/leanlabsio/kanban/issues/14).

# 1.0.8 (2015-05-11)

## Features:

- Filter by milestones with "^" notation. Part of [#12](https://gitlab.com/leanlabsio/kanban/issues/12).
- Filter autocomplete for users and milestones. Part of [#12](https://gitlab.com/leanlabsio/kanban/issues/12).

# 1.0.7 (2015-05-07)

## Features:

- The description on creating an issue is now optional. Closes [#19](https://gitlab.com/leanlabsio/kanban/issues/19).
- Display the project name in browser window title. Closes [#17](https://gitlab.com/leanlabsio/kanban/issues/17).
- Filter cards by assignee with "@" notation. Part of [#12](https://gitlab.com/leanlabsio/kanban/issues/12).

# 1.0.6 (2015-05-06)

## Features:

- Use the same URI routing scheme as GitLab. Closes [#18](https://gitlab.com/leanlabsio/kanban/issues/18).

# 1.0.5 (2015-04-27)

## Deprecations:

- Deprecated fig.yml in favour of docker-compose.yml.

# 1.0.4 (2015-04-04)

## Bugfixes:

- Fixed broken authorization.

# 1.0.3 (2015-03-25)

## Bugfixes:

- Fixed RabbitMQ docker image version.

# 1.0.1 (2015-03-21)

## Deprecations:

- Removed contstraint to run only on HTTPS.

# 1.0.0 (2015-02-27)

## Features:

- Initial stable release
