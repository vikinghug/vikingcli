vikingcli
=========

Small cli utility to maintain your Viking repositories. This will clone all of the repositories into your WildStar addon directory (%APPDATA%/NCSOFT/WildStar/Addons).

Install
=======

If you have Go, simply use go get to install:

    go get github.com/jamal/vikingcli

Usage
=====

    vikingcli command [arguments]
    
The commands are:

    clone       clone all of the viking repositories to the current directory.
    branch      switch all repositories to a specific branch
    pull        pull the latest revision from GitHub
