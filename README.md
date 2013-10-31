# Go Quick Start Guide

This guide will walk you through deploying a Go application on Deis.

## Prerequisites

* A [User Account](http://docs.deis.io/en/latest/client/register/) on a [Deis Controller](http://docs.deis.io/en/latest/terms/controller/).
* A [Deis Formation](http://docs.deis.io/en/latest/gettingstarted/concepts/#formations) that is ready to host applications

If you do not yet have a controller or a Deis formation, please review the [Deis installation](http://docs.deis.io/en/latest/gettingstarted/installation/) instructions.

## Setup your workstation

* Install [RubyGems](http://rubygems.org/pages/download) to get the `gem` command on your workstation
* Install [Foreman](http://ddollar.github.com/foreman/) with `gem install foreman`
* Install [Go](https://code.google.com/p/go/downloads/list). Use `brew install go` for Mac OS X

## Clone your Appication 

If you want to use an existing application, no problem.  You can also use the Deis sample application located at <https://github.com/bengrunfeld/example-go>.  Clone the example application to your local workstation:

    $ git clone https://github.com/bengrunfeld/example-go.git
    $ cd example-go

## Prepare your Application

To use a Go application with Deis, you will need to conform to 3 basic requirements:

 1. Manage dependencies via your `Go` source code
 2. Use [Foreman](http://ddollar.github.com/foreman/) to manage processes
 3. Use [Environment Variables](https://help.ubuntu.com/community/EnvironmentVariables) to manage configuration inside your application

If you're deploying the example application, it already conforms to these requirements.

#### 1. Manage dependencies via your Go source code

[Go](http://golang.org/) manages via import statements in your source code which can be pointed at [GitHub](http://github.com), [BitBucket](https://bitbucket.org/), or [Google Code](https://code.google.com/) repositories. E.g.

	import "github.com/stvp/gostatsd"
	
It's as simple as that.

#### 2. Use Foreman to manage processes

Deis relies on a [Foreman](http://ddollar.github.com/foreman/) `Procfile` that lives in the root of your repository.  This is where you define the command(s) used to run your application.  Here is an example `Procfile`:

    web: example-go

This tells Deis to run `web` workers using the command `example-go`. You can test this locally by running `foreman start`.

	$ foreman start
	13:41:38 web.1  | started with pid 2466
	13:41:38 web.1  | listening on 5000...
	
You should now be able to access your application locally at <http://localhost:5000>.

If you are having issues with Foreman, please refer to [Mark McGranaghan's](http://mmcgrana.github.io/2012/09/getting-started-with-go-on-heroku.html) excellent tutorial on getting Go set up locally.

#### 3. Use Environment Variables to manage configuration

Deis uses environment variables to manage your ¨application's configuration. For example, your application listener must use the value of the `PORT` environment variable. The following code snippet demonstrates how this can work inside your application:

    port := os.Getenv("PORT")

## Create a new Application

Per the prerequisites, we assume you have access to an existing Deis formation. If not, please review the Deis [installation instuctions](http://docs.deis.io/en/latest/gettingstarted/installation/).

Use the following command to create an application on an existing Deis formation.

    $ deis create --formation=<formationName> --id=<appName>
	Creating application... done, created <appName>
	Git remote deis added
    
If an ID is not provided, one will be auto-generated for you.

## Deploy your Application

Use `git push deis master` to deploy your application.

	$ git push deis master

	       Go app detected
	-----> Using Go 1.1.2
	-----> Running: go get -tags heroku ./...
	-----> Discovering process types
	       Procfile declares types -> web
	
	-----> Compiled slug size: 1.2 MB
	       Launching... 

Once your application has been deployed, use `deis open` to view it in a browser. To find out more info about your application, use `deis info`.

## Scale your Application

To scale your application's [Docker](http://docker.io) containers, use `deis scale` and specify the number of containers for each process type defined in your application's `Procfile`. For example, `deis scale web=8`.

	$ deis scale web=8
	Scaling containers... but first, coffee!
	done in 16s
	
	=== <appName> Containers
	
	--- web: `example-go`
	web.1 up 2013-10-29T16:30:24.690Z (goFormation-runtime-1)
	web.2 up 2013-10-29T16:46:25.719Z (goFormation-runtime-1)
	web.3 up 2013-10-29T16:46:25.738Z (goFormation-runtime-1)
	web.4 up 2013-10-29T16:46:25.754Z (goFormation-runtime-1)
	web.5 up 2013-10-29T16:46:25.771Z (goFormation-runtime-1)
	web.6 up 2013-10-29T16:46:25.790Z (goFormation-runtime-1)
	web.7 up 2013-10-29T16:46:25.810Z (goFormation-runtime-1)
	web.8 up 2013-10-29T16:46:25.831Z (goFormation-runtime-1)


## Configure your Application

Deis applications are configured using environment variables. The example application includes a special `POWERED_BY` variable to help demonstrate how you would provide application-level configuration. 

	$ curl -s http://yourapp.yourformation.com
	Powered by Deis
	$ deis config:set POWERED_BY=Go
	=== <appName>
	POWERED_BY: Go
	$ curl -s http://yourapp.yourformation.com
	Powered by Go

`deis config:set` is also how you connect your application to backing services like databases, queues and caches. You can use `deis run` to execute one-off commands against your application for things like database administration, initial application setup and inspecting your container environment.

	$ deis run ls -la
	total 52
	drwxr-xr-x  4 root root  4096 Oct 30 20:23 .
	drwxr-xr-x 57 root root  4096 Oct 30 20:25 ..
	-rw-r--r--  1 root root   252 Oct 30 20:23 .gitignore
	-rw-r--r--  1 root root    11 Oct 30 20:23 .godir
	drwxr-xr-x  2 root root  4096 Oct 30 20:23 .profile.d
	-rw-r--r--  1 root root    58 Oct 30 20:23 .release
	-rw-r--r--  1 root root 10273 Oct 30 20:23 LICENSE
	-rw-r--r--  1 root root    16 Oct 30 20:23 Procfile
	-rw-r--r--  1 root root  2001 Oct 30 20:23 README.md
	drwxr-xr-x  2 root root  4096 Oct 30 20:23 bin
	-rw-r--r--  1 root root   660 Oct 30 20:23 web.go

## Troubleshoot your Application

To view your application's log output, including any errors or stack traces, use `deis logs`.

    $ deis logs
	Oct 29 16:31:29 ip-172-31-27-8 gaslit-electron[web.1]: listening on 10001...
	Oct 29 16:43:55 ip-172-31-27-8 gaslit-electron[web.1]: listening on 10001...
	Oct 29 16:46:36 ip-172-31-27-8 gaslit-electron[web.2]: listening on 10002...

## Additional Resources

* [Get Deis](http://deis.io/get-deis/)
* [GitHub Project](https://github.com/opdemand/deis)
* [Documentation](http://docs.deis.io/)
* [Blog](http://deis.io/blog/)
