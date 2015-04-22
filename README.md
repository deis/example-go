# Go Quick Start Guide

This guide will walk you through deploying a Go application on Deis.

## Usage

```
$ deis create
Creating application... done, created rising-yearbook
Git remote deis added
$ git push deis master
Counting objects: 17, done.
Delta compression using up to 8 threads.
Compressing objects: 100% (9/9), done.
Writing objects: 100% (17/17), 8.20 KiB | 0 bytes/s, done.
Total 17 (delta 5), reused 17 (delta 5)
-----> Go app detected
-----> Installing go1.2.1... done

       Tired of waiting for bzr and hg?
       Try github.com/kr/godep for faster deploys.

       Installing Virtualenv... done
       Installing Mercurial... done
       Installing Bazaar... done
-----> Running: go get -tags heroku ./...
-----> Discovering process types
       Procfile declares types -> web
-----> Compiled slug size is 1.7M
remote: -----> Building Docker image
remote: Uploading context 1.803 MB
remote: Uploading context
remote: Step 0 : FROM deis/slugrunner
remote:  ---> 5567a808891d
remote: Step 1 : RUN mkdir -p /app
remote:  ---> Using cache
remote:  ---> 4096b5c0b838
remote: Step 2 : ADD slug.tgz /app
remote:  ---> 84cf8072cc65
remote: Removing intermediate container 6f0a2985332c
remote: Step 3 : ENTRYPOINT ["/runner/init"]
remote:  ---> Running in cf364904c7df
remote:  ---> b0685acc120c
remote: Removing intermediate container cf364904c7df
remote: Successfully built b0685acc120c
remote: -----> Pushing image to private registry
remote:
remote:        Launching... done, v2
remote:
remote: -----> rising-yearbook deployed to Deis
remote:        http://rising-yearbook.local.deisapp.com
remote:
remote:        To learn more, use `deis help` or visit http://deis.io
remote:
To ssh://git@local.deisapp.com:2222/rising-yearbook.git
 * [new branch]      master -> master
$ curl http://rising-yearbook.local.deisapp.com
Release v3 Powered by Deis on c6f9ffdeda29
$ deis scale web=4
$ curl http://rising-yearbook.local.deisapp.com
Release v3 Powered by Deis on c6f9ffdeda29
$ curl http://rising-yearbook.local.deisapp.com
Release v3 Powered by Deis on f1c20002e957
$ curl http://rising-yearbook.local.deisapp.com
Release v3 Powered by Deis on 58dd1772b8ef
```

## Additional Resources

* [Get Deis](http://deis.io/get-deis/)
* [GitHub Project](https://github.com/deis/deis)
* [Documentation](http://docs.deis.io/)
* [Blog](http://deis.io/blog/)
