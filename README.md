# Goyts
a simple local webserver as a convenient frontend to youtube-dl

Based on a simple webserver written in golang, goyts offers a mega simple web frontend and allows you to prompt a URL which will then be handled over to youtube-dl. 
Goyts allows for a json based configuration where you can specify:
* dir where goyts is running
* webserver port
* dir where downloads should be written to (e.g. a mount directly pointing to your NAS)
* youtube options for filename
* youtube options for e.g specifying a max. video resolution (otherwise youtube-dl will try to get maximum resolution)

Goyts is open source.

# Installation

There is none: simply clone these files containing goyts to where you want it to reside using git clone.

But you have to build an executable for your platform:
* install go language
* run 'go build'
** executable should be built

The executable is self contained! So you can build the exe on one platform (eg. x86) and the run it on all your x86 machines without any runtime environment needed (one of the advantages of the go language).

# Configuration
create a plain text file config.json based on the example.

# Running goyts

on bash run: 'nohup Goyts &' to run and detach it from your current process.

# test ist and use it

Have a look at the logfile. Goyts does not record successful youtube-dl-executions, but it catches stderr-output of youtube-dl to the logfile.

# have fun!
