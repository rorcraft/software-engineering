### Process

__1. Story Kick off__
* Full team at kick off meetings. 
* Ideally batch multiple stories at the same time
* Explain requirement and acceptance criteria, try to nail down the scope.
* If stories are straight forward, the team can provide estimate. 
* Complex stories, set a time box for designing solution and review - output is estimation and implementation plan(s).
* QE should start planning for test plan and write up acceptance criteria.
* QE should identify risk areas as input to implementation plan.

__1a. Time boxed spike (for complex stories)__
* Aim to design 2-3 solutions
* Write exploratory code (throw away) or prototype code outside of the app.
* Identify technical risk areas and attempt to confirm assumptions (proof of concept).
* If more time is required for particular areas, this should spawn other specific spike with separate goals.
* Implementation plan - break down solution into subtasks/stories and estimates for review. (see. <"Definition of Done">)

__1b. Implementation Review (for complex stories)__
* Review solution(s), compare pro/cons of each solution based on time to complete, code quality improvement/damage, user value of feature.
* During time box, engineers should already be clarifying assumptions and contraints with Product, but this meeting will be the final confirmation before implementation.
* QE should review the plan and confirm that risk areas are covered.
* TL or Infrastructure or domain experts should be involved in review potential impact.
* Output is a story with subtasks or multiple stories required to achieve the acceptance criteria.
* Try to break down story as much as possible that is deployable, ie deliver small amount of value.

__2. Implementation__
* Engineers should pair on complex stories, e.g. changes to storage, integrating external api, complicate UI interactions.
* If engineer feel confident on executing implementation plan, that work can be done in solo.
* Any blockers should be raised to team lead and/or at standup.
* For UI related stories, the look and feel should be prototyped first and confirm with Design and Product Manager early, before integrating the actual functionality.
* TDD is ideal, analyze the acceptance criteria and break it down into steps, involve QE if necessary.
* Check the <"Definition of Done"> list.

__3. Acceptance__
* Product Manager should review
* QE should test and verify according to acceptance criteria
* Team Lead or other engineer of the team should code review (see <Code Review Checklist>)
* Code review is mandatory for solo or paired work, reviewer should merge to latest or hitched in the merge commit.
* Also check the <"Definition of Done"> again.

__4. Ready to Deploy__
* Check <deployment checklist>, check your sh!t on staging, run blackbox smoke, turn on experiment(s), look for errors, check mixpanel.

__5. Deploy, Learning__
* Check <deployment checklist>
* Watch airbrake, newsrelic
* Turn on experiment(s), check mixpanel
