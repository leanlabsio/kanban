# 1.4.4 (2016-01-02)

## Features:

- New filter system, by multiple tags (Michael Filonenko, V)
- Collapsible groups when grouping enabled (Michael Filonenko)

# 1.4.3 (2015-12-03)

## Features:

- Swimlanes by milestones, resolves [#109](https://gitlab.com/leanlabsio/kanban/issues/109) (Don Mahurin).
- Create milestones from kanban, resolves [#105](https://gitlab.com/leanlabsio/kanban/issues/105) (Anton Popov).

## Bugfixes:

- Solved infinite page reload, resolves [#112](https://gitlab.com/leanlabsio/kanban/issues/112) (V).
- Fixed modal window history navigation, resolves [#108](https://gitlab.com/leanlabsio/kanban/issues/108) (Anton Popov).

# 1.4.2 (2015-11-22)

## Bugfixes:

- Projects with dot in name now works, resolves [#91](https://gitlab.com/leanlabsio/kanban/issues/91) (Anton Popov).
- Show group and project members in "Assign" menu, resolves [#94](https://gitlab.com/leanlabsio/kanban/issues/94) (Anton Popov).
- Updated GitLab logo in signin form (Michael Filonenko).
- Fixed comments spam, when moving card in same column (Michael Filonenko).
- Changed confusing text in sign in form to be more user friendly (Anton Popov)

# 1.4.0 (2015-11-01)

## Features:

- Rewritten server side to enable installation without docker

# 1.3.0 (2015-09-06)

## Features:
- Started complete rewrite in golang.

## Bugfixes:

- Added versioning for static assets [#252](https://gitlab.com/kanban/client/issues/252).
- Use project avatar if not empty [#248](https://gitlab.com/kanban/client/issues/248).
- Project description is styled correctly [#247](https://gitlab.com/kanban/client/issues/247)
- <code> block styles [#241](https://gitlab.com/kanban/client/issues/241).

# 1.2.2 (2015-08-17)

## Features:

- Enabled plugins integration

# 1.2.1 (2015-08-13)

## Bugfixes

- assets versioning

# 1.2.0 (2015-08-08)

## Features:

- Implemented swimlanes by users [#243](https://gitlab.com/kanban/client/issues/243).
- Added ability to unassign issues [#61](https://gitlab.com/leanlabsio/kanban/issues/61).
- Redesigned boards

# 1.1.4 (2015-07-05)

## Features:

- Implemented andon
- Card movement notifications
- Markdown preview in textareas

## Bugfixes:

- Cannot remove last todo. Closes [#44](https://gitlab.com/leanlabsio/kanban/issues/44).

# 1.1.3 (2015-06-21)

## Features:

- Mobile friendly card view, card create.
- Redesigned comments.

## Bugfixes:

- Do not add empty TODO items. Closes [#53](https://gitlab.com/leanlabsio/kanban/issues/53).
- Every TODO item adds a new line to the beginning of the description. Closes [#46](https://gitlab.com/leanlabsio/kanban/issues/46).
- Every second TODO item has a newline. Closes [#45](https://gitlab.com/leanlabsio/kanban/issues/45).
- margin of columns lets board overflow. Closes [#35](https://gitlab.com/leanlabsio/kanban/issues/35).
- The horizontal scroll bar is only visible when you scroll down, instead of always showing in the viewport. Closes [#6](https://gitlab.com/leanlabsio/kanban/issues/6).

# 1.1.2 (2015-06-12)

## Features:

- Filter cards by label. Closes [#233](https://gitlab.com/kanban/client/issues/233).
- Textarea in comment form. Closes [#40](https://gitlab.com/leanlabsio/kanban/issues/40).

# 1.1.1 (2015-06-12)

## Refactoring:

- Reduced wsserver image size

# 1.1.0 (2015-06-03)

## Features:

- Add "cancel" button for edit issue functionality. Closes [#39](https://gitlab.com/leanlabsio/kanban/issues/39).
- Rename "Edit card" to "Edit issue" and "/boards/.../cards/..." to "/boards/.../issues/...". Closes [#41](https://gitlab.com/leanlabsio/kanban/issues/41).

## Bugfixes:

- Every TODO item adds a new line to the beginning of the description. Closes [#46](https://gitlab.com/leanlabsio/kanban/issues/46).
- Modifying the issue title updates title in board without safing. Closes [#42](https://gitlab.com/leanlabsio/kanban/issues/42).
- Fixed base docker images versions. Fixed Redis container version. Fixed proxy container version.

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
