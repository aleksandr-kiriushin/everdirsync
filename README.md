![](./assets/jenkins.svg)
[![Build Status](https://ci.everalerta.com/buildStatus/icon?job=everdirsync%2Fmain)](https://ci.everalerta.com/job/everdirsync/job/main/)

# EverDirSync
EverDirSync is a client-server tool for syncing directories over network


# Install

## Using Github releases
* Download the latest github release
* Move archive to the new folder: `cd ~/Downloads; mkdir everdirsync; mv everdirsync_0.0.1_Linux_x86_64.tar.gz everdirsync; cd everdirsync`
* Extract binary file from the archive: `tar xvf everdirsync_0.0.1_Linux_x86_64.tar.gz`
* Add execution permission to this file: `chmod +x everdirsync`
* Run the program: `./everdirsync`
* Optionally, copy binary file to `/usr/local/bin/` to run program using `everdirsync`

## Using go get
* Install everdirsync: `go get -v github.com/aleksandr-kiriushin/everdirsync`
