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

# Configuration

create a plain text file config.json with the following content (and adapt to your needs):
---
{
        "basedir": "/home/myself/Skripts/Goyts",
        "port":1234,
        "videodir": "/mnt/Recordings/Youtube/",
        "options": "%(title)s",
        "videoformat": "bestvideo[height<=1080]+bestaudio/best[height<=1080]"
}
---
