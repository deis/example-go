# Go Quick Start Guide

This guide will walk you through deploying a Go application on [Deis Workflow][].

## Usage

```
$ deis create
Creating application... done, created luxury-waxworks
Git remote deis added
$ git push deis master
Counting objects: 49, done.
Delta compression using up to 4 threads.
Compressing objects: 100% (44/44), done.
Writing objects: 100% (49/49), 12.60 KiB | 0 bytes/s, done.
Total 49 (delta 18), reused 0 (delta 0)
-----> Go app detected
-----> Installing go1.4.2... done
-----> Running: godep go install -tags heroku ./...
-----> Discovering process types
       Procfile declares types -> web
-----> Compiled slug size is 1.7M

-----> Building Docker image
remote: Sending build context to Docker daemon 1.723 MB
remote: build context to Docker daemon
Step 0 : FROM deis/slugrunner
 ---> 553ece96de1d
Step 1 : RUN mkdir -p /app
 ---> Running in 79a14e4fa092
 ---> dacb2ee0e3f3
Removing intermediate container 79a14e4fa092
Step 2 : WORKDIR /app
 ---> Running in c5dcb81816d8
 ---> c19b88f2cff7
Removing intermediate container c5dcb81816d8
Step 3 : ENTRYPOINT /runner/init
 ---> Running in 2badb1515dfc
 ---> d2bf6069fa3a
Removing intermediate container 2badb1515dfc
Step 4 : ADD slug.tgz /app
 ---> 774507a8209f
Removing intermediate container 17cc7ac67a0c
Step 5 : ENV GIT_SHA 3800239fbd11c2ac79e22c1db82aeff0d631dc2b
 ---> Running in d4029c60d514
 ---> 269bf440d703
Removing intermediate container d4029c60d514
Successfully built 269bf440d703
-----> Pushing image to private registry

-----> Launching...
       done, luxury-waxworks:v2 deployed to Deis

       http://luxury-waxworks.local3.deisapp.com

       To learn more, use `deis help` or visit http://deis.io

To ssh://git@deis.local3.deisapp.com:2222/luxury-waxworks.git
 * [new branch]      master -> master
$ curl http://luxury-waxworks.local3.deisapp.com
Powered by Deis
Release v2 on 2ce9ed10d21d
$ deis scale web=3
Scaling processes... but first, coffee!
done in 15s
=== luxury-waxworks Processes

--- web:
web.1 up (v2)
web.2 up (v2)
web.3 up (v2)

$ curl http://luxury-waxworks.local3.deisapp.com
Powered by Deis
Release v2 on 2ce9ed10d21d
$ curl http://luxury-waxworks.local3.deisapp.com
Powered by Deis
Release v2 on 3a1234b1b2b2
$ curl http://luxury-waxworks.local3.deisapp.com
Powered by Deis
Release v2 on 47ebe13e1c09
```

## Additional Resources

* [GitHub Project](https://github.com/deis/workflow)
* [Documentation](https://deis.com/docs/workflow/)
* [Blog](https://deis.com/blog/)

[Deis Workflow]: https://github.com/deis/workflow#readme
