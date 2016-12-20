# 1.7.1 (2016-12-20)

## Features

- Add config to disable auto-comments on change [#78](https://gitlab.com//leanlabsio/kanban/issues/78) (cnam).
- GitLab flavoured markdown [#135](https://gitlab.com//leanlabsio/kanban/issues/135) (cnam).
- Mylestones, user, labels search for select [#31](https://gitlab.com//leanlabsio/kanban/issues/31) (cnam).

## Bugfixes:

- Assign is reset when an issue is updated [#222](https://gitlab.com//leanlabsio/kanban/issues/222) (Bethuys).

# 1.7.0 (2016-11-15)

## Features:

- Attach file for card, comment [#130](https://gitlab.com//leanlabsio/kanban/issues/130) (cnam).
- Show icon with comments count [#26](https://gitlab.com//leanlabsio/kanban/issues/26) (cnam).
- Add support for tasks [#183](https://gitlab.com//leanlabsio/kanban/issues/183) (cnam).
- Option to show only starred projects on list of boards [#177](https://gitlab.com//leanlabsio/kanban/issues/177) (cnam).
- Allow to set issue's due date (GitLab 8.7+) [#175](https://gitlab.com//leanlabsio/kanban/issues/175) (cnam).
- "Add Issue" dialog should set focus on title input [#72](https://gitlab.com//leanlabsio/kanban/issues/72) (cnam).

## Bugfixes:

- one of our issue is not appear in kanban view [#162](https://gitlab.com//leanlabsio/kanban/issues/162) [#89](https://gitlab.com//leanlabsio/kanban/issues/89) (cnam, Bethuys).
- 100+ projects, home page don't display all projects [#210](https://gitlab.com//leanlabsio/kanban/issues/210) (cnam).
- Fix modal close area hit detection [#203](https://gitlab.com//leanlabsio/kanban/issues/203) (cnam).

# 1.6.2 (2016-06-19)

## Features:

- Save a draft so you can close a not yet created issue [#118](https://gitlab.com//leanlabsio/kanban/issues/118) (cnam).
- Save a draft new comment [#117](https://gitlab.com//leanlabsio/kanban/issues/117) (cnam).
- Save a draft editable issue [#117](https://gitlab.com//leanlabsio/kanban/issues/117) (cnam).
- Milestones are sorted alphanumerically [#186](https://gitlab.com//leanlabsio/kanban/issues/186) (cnam).
- Inconsistent display of priority [#179](https://gitlab.com//leanlabsio/kanban/issues/179) (cnam).
- Date not displayed for Milestone in Swimlane view [#172](https://gitlab.com//leanlabsio/kanban/issues/172) (cnam).
- Allow use of Redis unix socket [#98](https://gitlab.com//leanlabsio/kanban/issues/98) (cnam).

## Bugfixes:

- Auto refresh to often, save all editable data as draft [#163](https://gitlab.com//leanlabsio/kanban/issues/163) (cnam).
- Priority with negative order is incorrectly handled [#199](https://gitlab.com//leanlabsio/kanban/issues/199) (cnam).
- Issue with long description or comments can't be scrolled [#157](https://gitlab.com//leanlabsio/kanban/issues/157) (cnam).

# 1.6.1 (2016-06-13)

## Features:

- Add favicon [#176](https://gitlab.com//leanlabsio/kanban/issues/176) (cnam, Toby Mole, Maël Lavault).

## Bugfixes:

- Problem white text on white label [#181](https://gitlab.com//leanlabsio/kanban/issues/181) (cnam).
- Add validation for board settings [#182](https://gitlab.com//leanlabsio/kanban/issues/182) (cnam).
- Regex preventing "moved issue" notes creation (Clément Bethuys).

## Maintanence:

- Update angular to 1.5 version [#195](https://gitlab.com//leanlabsio/kanban/issues/195) (cnam).
- Add contribution guides [#197](https://gitlab.com//leanlabsio/kanban/issues/197) (cnam).

# 1.6.0 (2016-05-04)

## Features:

- Cards prioritization [#119](https://gitlab.com/leanlabsio/kanban/issues/119) (TruongSinh Tran-Nguyen, cnam).
- Swimlanes by priority [#119](https://gitlab.com/leanlabsio/kanban/issues/119) (cnam).
- Configurable priorities from within board, settings [#119](https://gitlab.com/leanlabsio/kanban/issues/119) (cnam).
- Filter by priority [#119](https://gitlab.com/leanlabsio/kanban/issues/119) (cnam).

## Bugfixes:

- Fixed: Infinite redirects to non-existent board [#168](https://gitlab.com/leanlabsio/kanban/issues/168) (cnam).
- Fixed: Limits not updated when filtering by milestone [#167](https://gitlab.com/leanlabsio/kanban/issues/167) (V).

# 1.5.1 (2016-03-12)

## Bugfixes:

- Fixed: update of stage label causes logout (V).
- Fixed static assets versioning and configured build to get version from makefile [#164](https://gitlab.com/leanlabsio/kanban/issues/164) (V).
- Fixed stage regexp: make it non gready for not to get "][" symbols in stage name (V).

# 1.5.0 (2016-03-07)

## Features:

- Improved mobile experience [#110](https://gitlab.com/leanlabsio/kanban/issues/110) (Anton Popov).
- Improved cards filtering [#141](https://gitlab.com/leanlabsio/kanban/issues/141) (V).
- Implemented board settings, to manage stages from board [#15](https://gitlab.com/leanlabsio/kanban/issues/15) (V, Anton Popov).

## Maintanence:

- Removed ng-tags-input library, not used [#147](https://gitlab.com/leanlabsio/kanban/issues/147) (V).


# 1.4.7 (2016-01-31)

## Features:

- Redirect back to previous URL after the restoration of the authorization session [#156](https://gitlab.com/leanlabsio/kanban/issues/156) (Алексей Кукушкин).
- Allow to disable sign up [#68](https://gitlab.com/leanlabsio/kanban/issues/68) (V).

## Maintanence:

- Implemented data source adapter to enable pluggable storage backends [#154](https://gitlab.com/leanlabsio/kanban/issues/154) (V).
- Updated vendored packages (V).
- Compilation with go 1.5.3 (V).

# 1.4.6 (2016-01-16)

## Features:

- Show emoji [#144](https://gitlab.com/leanlabsio/kanban/issues/144) (V).

## Bugfixes:

- Issue description is not removed when closing issue [#145](https://gitlab.com/leanlabsio/kanban/issues/145) (V).
- Filter now shows only active milestones [#151](https://gitlab.com/leanlabsio/kanban/issues/151) (V).
- Card filter not duplicated [#149](https://gitlab.com/leanlabsio/kanban/issues/149) (V).

# 1.4.5 (2016-01-07)

## Features:

- Card counts against columns names [#52](https://gitlab.com/leanlabsio/kanban/issues/52) (V).
- Card assignee could be cleared [#61](https://gitlab.com/leanlabsio/kanban/issues/61) (V).
- Card milestone could be cleared [#136](https://gitlab.com/leanlabsio/kanban/issues/136) (V).
- Use two digits for default columns [#97](https://gitlab.com/leanlabsio/kanban/issues/97) (V).

## Bugfixes:

- Images uploaded via GitLab are displayed correctly [#90](https://gitlab.com/leanlabsio/kanban/issues/90) (V).
- Board state is local to every board [#134](https://gitlab.com/leanlabsio/kanban/issues/134) (V).
- Added new info comment patterns [#137](https://gitlab.com/leanlabsio/kanban/issues/137) (V).

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
