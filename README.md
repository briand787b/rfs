# RFS - A Distributed File Server
#### (Probably should've been named dfs)

## How to use it
Because RFS is written in Go, it can be compiled for multiple operating systems and/or 
architectures on a single machine .  For example, my router runs a Debian-derived
distribution of Linux on an ArmV7 architecture.  Installation on this platform is not
an issue for RFS since a cross-compiled binary of itself can be copied onto it and 
run without need for a virtual machine or pre-installed runtime. 

The orchestration of RFS is controlled through a single master node that maintains a 
master list of all the media files that should exist on each server.  Although RFS
maintains the designated media files on every server it is told to manage, it does
NOT delete files.  If a file is found that is not expected in the servers manifest,
RFS will add that file to the server's expected list of files.  

The master list of all files and servers is stored in a SQL database on the master
node.  All additions and removals of media on RFS-controlled nodes (computers) 
should be performed through master node's interface.  Interaction with the master
node is achieved through two interfaces:
    1. *Web App* - An HTTP web server is hosted on the master node through which
        users can perform all functionality
    2. *Command Line Interface (CLI)* - Many of the functions of RFS can be
        performed through invocations of the application binary, passing it 
        correct sub commands and arguments.  
The web app will usually be more fully featured than the CLI.  Delivery and
manipulation of HTML will be accomplished through the use of the Go standard 
library's templating engine.  Presentation of file transfer progress will not be 
immediately available by RFS, although it is in the works!  The MVP of RFS aims to
merely have a binary file transfer status: in-transit or stationary.

## How it works
This application defines the term 'media' (ignoring correct singularization) to
represent any number of entertainment artifacts that can be logically grouped into
a single entity.  Organization of media is hierarchical and one media may be considered
a 'parent' to zero or more media and a 'child' of exactly one media (unless it is basal
to its hierarchy tree, in which case it has no parent).  Although the logical grouping 
of digitial files is arbitrary, most media will be represented exactly as they are 
obtained from their source.  For example, a TV show can be represented in the data
layer of this application like so:
```
Show X  -> Season 1 -> Episode 1
                   \-> Episode 2
                   \-> Episode 3
       \-> Season 2 -> Episode 1
                   \-> Episode 2
```
