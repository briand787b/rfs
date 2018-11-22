# RFS - A Distributed File Server [WIP]
#### (Probably should've been named dfs)

## How to use it
Because RFS is written in Go, it can be compiled for multiple operating systems and/or 
architectures on a single machine.  For example, my router runs a Debian-derived
distribution of Linux on an ArmV7 architecture.  Installation on this platform was not
an issue for RFS since I was able to cross compile an ARMv7-compatible binary from my
Mac and run it on the server without need for a virtual machine or pre-installed runtime. 

The orchestration of RFS is controlled through a single master node that maintains a 
master list of all the media files that should exist on each server.  Although RFS
maintains the designated media files on every server it is told to manage, it does
NOT delete files.  If a file is found that is not expected in the server's manifest,
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

Although RFS can be installed locally, it has been designed from the ground up to
be installed on the cloud.  It will be built to run on AWS initially but may be
extended to other providers in the future.  Because the master node will be 
installed on the cloud, it will be accessible from anywhere, thus extending the
reach of RFS to deliver content wherever it is demanded.

## How it works
### Media
This application defines the term 'media' (ignoring correct singularization) to
represent a grouping of any number of entertainment artifacts that can be 
logically organized into a single entity.  Organization of media is hierarchical
and one media may be considered a 'parent' to zero or more media and a 'child' 
of exactly one media (unless it is basal to its hierarchy tree).  Although the 
logical grouping of digitial files is arbitrary, most media will be represented 
exactly as they are obtained from their source. For example, a TV show can be 
represented in the data layer of this application like so:
```
TV Shows -> Show X -> Season 1 -> Episode 1
                 |           | -> Episode 2
                 |            \-> Episode 3
                  \-> Season 2 -> Episode 1
                              \-> Episode 2
```
The root media are
    1. Movies
    2. TV Shows
    3. Books
    4. Photos
    5. Miscellaneous
Again, this is completely arbitrary but should be intuitive.  This logical organization
of media manifests itself in the organization of the files managed by RFS.  The
filesystem on which 'Show X' resides should have its directory structure be identical
to the logical representation above.  For example, Episode 2 of Season 1 should have 
this path:
```
${RFS_WORKING_DIR}/TV_Shows/Show_X/Season_1/Episode_2
```
### File
A file is a concrete artifact that relates to exactly one media.  It has a 1:1
relationship with a file in the traditional sense.  Although a single media may
contain many files, only one of its files may be considered the 'feature' file.
This file should be the main purpose of the media if it is not merely an 
organizational unit like the root level media.  A checksum is performed against
this feature file, which is later used for validation of data integrity.
Although it is possible for a media to not contain a feature file, it must not
contain more than one.  For example, the correct location of the playable file
that contains episode 1 of Show X's first season would be in the Episode_1 folder:
```
${RFS_WORKING_DIR}/TV_Shows/Show_X/Season_1/Episode_1/episode_1.mp4
```
'Episode_1', and so Episode_1 would contain the checksum of this file for data.
The file 'episode_1.mp4' would obviously be the feature file for the media named
validation.  Other files may be associated with this media, including subtitles
(.srt files) or artwork, but these files must not hold equal importance to the
feature file.  

As an implementation detail, the database must ensure that no two files of the
same media share the same name.  This would cause one of the files to be
overwritten when saved to the node's disk.
### Master Node - Client Node Communication
RFS's master node will be initially hosted on AWS.  Client nodes, however, can
be installed wherever the content is required to exist.  This means that the 
two nodes must be able to communicate even when on different networks, with 
one potentially being behind a firewall.  Although initial plans were to require
the opening of a firewall on the router, doing so was deemed to open up a potential
security vulnerability and also significantly limit the potential range of applications
since the end user may not have the authority to do so on the networks the hosts reside
on.  For this reason, all communication between master and client will follow the 
client-server model where communication must always be initiated by the client.  

The requirement of client initiation will be accomodated by the use of web sockets.
A client will register itself with the master server.  It will then wait until
it recieves an instruction from the master node.  Transfers of media will be
communicated to the server in this way.  Media transers are unique, however, in
that they ultimately need to communicate data from one client to another.  To
achieve this, the master server acts as an intermediary that streams the outgoing
bytes from the sending client to the receiving client.  In later iterations, the
master server may be able to detect if the two clients are on the same network 
and instruct them to speak to each other and not use the master node as an
intermediary (this may become part of the MVP depending on the cost of routing
traffic through the cloud provider).

Client-Master communication will occur on port 2269 through grpc, while all 
administrative communication for the master server will occur on port 80 and
administrative communication with each individual client will occur on port 8080.   
### Service Discovery
All clients will listen on port 2269

### Media Transmission
Media 

## How to install it
### Client
Each client will be run as a daemon on its respective host.  Upon startup,
it will run a web server through which administrative tasks can be
performed.  The first will be to configure the client.  This involves
providing the client with this information:
