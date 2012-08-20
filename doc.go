/*
tsuru is a command line tool for application developers.

It provide some commands that allow a developer to register himself/herself,
manage teams, apps and services.

Usage:

	% tsuru <command> [args]

The currently available commands are (grouped by subject):

    target            changes or retrive the current tsuru server

    user-create       creates a new user
    login             authenticates the user with tsuru server
    logout            finishes the session with tsuru server
    key-add           adds a public key to tsuru deploy server
    key-remove        removes a public key from tsuru deploy server

    team-create       creates a new team (adding the current user to it automatically)
    team-list         list teams that the user is member
    team-user-add     adds a user to a team
    team-user-remove  removes a user from a team

    app-create        creates an app
    app-remove        removes an app
    app-list          lists apps that the user has access (see app-grant and team-user-add)
    app-grant         allows a team to have access to an app
    app-revoke        revokes access to an app from a team
    log               shows log for an app
    run               runs a command in all units of an app
	restart           restarts the app's application server

    env-get           display environment variables for an app
    env-set           set environment variable(s) to an app
    env-unset         unset environment variable(s) from an app

    bind              binds an app to a service instance
    unbind            unbinds an app from a service instance

    service-list      list all services, and instances of each service
    service-add       creates a new instance of a service
    service-remove    removes a instance of a service
    service-status    checks the status of a service instance
    service-info      list instances of a service, and apps binded to each instance
    service-doc       displays documentation for a service

Use "tsuru help <command>" for more information about a command.


Change/retrieve remote tsuru server

Usage:

	% tsuru target [target]

This command should be used to get current tsuru target, or retrieve current
target.

The target is the tsuru server to which all operations will be directed to.


Create a user

Usage:

	% tsuru user-create <email>

user-create creates a user within tsuru remote server. It will ask for the
password before issue the request.


Authenticate within remote tsuru server

Usage:

	% tsuru login <email>

Login will ask for the password and check if the user is successfully
authenticated. If so, the token generated by the tsuru server will be stored in
${HOME}/.tsuru_token.

All tsuru actions require the user to be authenticated (except login and
user-create, obviously).


Logout from remote tsuru server

Usage:

	% tsuru logout

Logout will delete the token file and terminate the session within tsuru
server.


Add SSH public key to tsuru's git server

Usage:

	% tsuru key-add [${HOME}/.ssh/id_rsa.pub]

key-add sends your public key to tsuru's git server. By default, it will try
send a public RSA key, located at ${HOME}/.ssh/id_rsa.pub. If you want to send
other file, you can call it with the path to the file. For example:

	% tsuru key-add /etc/my-keys/id_dsa.pub

The key will be added to the current logged in user.


Remove SSH public key from tsuru's git server

Usage:

	% tsuru key-remove [${HOME}/.ssh/id_rsa.pub]

key-remove removes your public key from tsuru's git server. By default, it will
try to remove a key that match you public RSA key located at
${HOME}/.ssh/id_rsa.pub. If you want to remove a key located somewhere else,
you can pass it as parameter to key-remove:

	% tsuru key-remove /etc/my-keys/id_dsa.pub

The key will be removed from the current logged in user.


Create a new team for the user

Usage:

	% tsuru team-create <teamname>

team-create will create a team for the user. Tsuru requires a user to be a
member of at least one team in order to create an app or a service instance.

When you create a team, you're automatically member of this team.


List teams that the user is member of

Usage:

	% tsuru team-list

team-list will list all teams that you are member of.


Add a user to a team

Usage:

	% tsuru team-user-add <teamname> <useremail>

team-user-add adds a user to a team. You need to be a member of the team to be
able to add a user to it.


Remove a user from a team

Usage:

	% tsuru team-user-remove <teamname> <useremail>

team-user-remove removes a user from a team. You need to be a member of the
team to be able to remove a user from it.

A team can never have 0 users. If you are the last member of a team, you can't
remove yourself from it.


Create an app

Usage:

	% tsuru app-create <appname> <platform>

app-create will create a new app using the given name and platform. For tsuru,
a platform is a Juju charm. To check the available platforms/charms, check this
URL: https://github.com/timeredbull/charms/tree/master/centos.

In order to create an app, you need to be member of at least one team. All
teams that you are member (see "tsuru team-list") will be able to access the
app.


Remove an app

Usage:

	% tsuru app-remove <appname>

app-remove removes an app. If the app is binded to any service instance, it
will be unbinded before be removed (see "tsuru unbind"). You need to be a
member of a team that has access to the app to be able to remove it (you are
able to remove any app that you see in "tsuru app-list").


List apps that the user has access to

Usage:

	% tsuru app-list

app-list will list all apps that you have access to. App access is controlled
by teams. If your team has access to an app, then you have access to it.


Allow a team to access an app

Usage:

	% tsuru app-grant <appname> <teamname>

app-grant will allow a team to access an app. You need to be a member of a team
that has access to the app to allow another team to access it.


Revoke from a team access to an app

Usage:

	% tsuru app-revoke <appname> <teamname>

app-revoke will revoke the permission to access an app from a team. You need to
have access to the app to revoke access from a team.

An app cannot be orphaned, so it will always have at least one authorized team.


See app's logs

Usage:

	% tsuru log <appname>

Log will show log entries for an app. These logs are not related to the code of
the app itself, but to actions of the app in tsuru server (deployments,
restarts, etc.).


Run an arbitrary command in the app machine

Usage:

	% tsuru run <appname> <command> [commandarg1] [commandarg2] ... [commandargn]

Run will run an arbitrary command in the app machine. Base directory for all
commands is the root of the app. For example, in a Django app, "tsuru run" may
show the following output:


	% tsuru run polls ls -l
	app.conf
	brogui
	deploy
	foo
	__init__.py
	__init__.pyc
	main.go
	manage.py
	settings.py
	settings.pyc
	templates
	urls.py
	urls.pyc


Restart the app's application server

Usage:

	% tsuru restart <appname>

Restart will call the restart hook from the app platform (the "restart" hook
from the Juju charm).
*/
package documentation
