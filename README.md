# Go Quick Start Guide

This guide will walk you through deploying a Go application on [Deis Workflow][].

## Usage

```console
$ git clone https://github.com/deis/example-go.git
$ cd example-go
$ deis create
Creating Application... done, created allied-question
Git remote deis added
remote available at ssh://git@deis-builder.deis.rocks:2222/allied-question.git
$ git push deis master
Counting objects: 89, done.
Delta compression using up to 4 threads.
Compressing objects: 100% (67/67), done.
Writing objects: 100% (89/89), 19.23 KiB | 0 bytes/s, done.
Total 89 (delta 37), reused 46 (delta 17)
Starting build... but first, coffee!
-----> Go app detected
-----> Checking Godeps/Godeps.json file.
-----> Installing go1.6.2... done
-----> Running: go install -v -tags heroku .
       github.com/deis/example-go
-----> Discovering process types
       Procfile declares types -> web
-----> Compiled slug size is 2.2M
Build complete.
Launching App...
Done, allied-question:v2 deployed to Deis

Use 'deis open' to view this application in your browser

To learn more, use 'deis help' or visit https://deis.com/

To ssh://git@deis-builder.deis.rocks:2222/allied-question.git
 * [new branch]      master -> master
$ curl http://allied-question.deis.rocks
Powered by Deis
Release v2 on allied-question-v2-web-wudcx
```

## Additional Resources

* [GitHub Project](https://github.com/deis/workflow)
* [Documentation](https://deis.com/docs/workflow/)
* [Blog](https://deis.com/blog/)

[Deis Workflow]: https://github.com/deis/workflow#readme
