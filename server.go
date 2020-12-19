package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strconv"
)

/* Declarations */

type Config struct {
	Basedir     string
	Port        int
	Videodir    string
	Options     string
	Videoformat string
}

// global config struct
var config Config

func main() {

	/* set up logfile, see https://www.honeybadger.io/blog/golang-logging/ */
	// If the file doesn't exist, create it or append to the file
	fmt.Println("Opening logfile")
	logfile, err := os.OpenFile("goyts.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	log.SetOutput(logfile)
	log.Println("=== Welcome to goyts - logging started ===")
	/* logging initialized */

	log.Println("Firing up goyts! Have fun!")

	/* init Config */
	readconfig()

	fmt.Println("config read!")

	/* starting up the webserver */
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/geturl", urlHandler)

	log.Println("Started server at port" + strconv.Itoa(config.Port))
	// fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":"+strconv.Itoa(config.Port), nil); err != nil {
		log.Fatal(err)
	}
}

func urlHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}
	fmt.Fprintf(w, "POST request successful\n")
	url := r.FormValue("url")
	// format := r.FormValue("format")
	fmt.Fprintf(w, "URL = %s\n", url)
	log.Println("Fetching URL " + url)
	log.Println("Using options " + config.Options)
	log.Println("Target Dir" + config.Videodir)
	optionstring := "-o" + config.Videodir + config.Options
	//optionstring := "-o" + config.Options
	log.Println("Full optionstring " + optionstring)
	// ytoptionf := "-f"
	// ytoptions := "bestvideo[height<=1080]+bestaudio/best[height<=1080]"
	formatoptions := "-f" + config.Videoformat

	// now get video using os-call
	//_, err := exec.Command("youtube-dl", optionstring, ytoptionf, ytoptions, url).Output()
	//_, err := exec.Command("youtube-dl", url).Output()
	cmd := exec.Command("youtube-dl", optionstring, formatoptions, url)
	// starting download concurrent
	err := cmd.Start()

	if err != nil {
		log.Printf("Error! %v", err)
	}
	log.Println("Download finished")

	// now move video(s) to videodir
	/*
		optionstring = "mp4 " + config.Videodir
		log.Println("Optionstring  for move: " + optionstring)
		_, err2 := exec.Command("mv", optionstring).Output()
		if err2 != nil {
			log.Printf("Error! %v", err2)
		}
		log.Println("Move videos finished")
		fmt.Fprintf(w, "Move videos finished")
	*/

}

func readconfig() {
	log.Println("started reading config")
	file, err := ioutil.ReadFile("config.json")
	if err != nil {
		log.Println("Error reading config file")
	}

	err2 := json.Unmarshal(file, &config)
	if err2 != nil {
		log.Println("Error parsing configfile")
		os.Exit(1)
	}

	log.Println("Basedir: ", config.Basedir)
	log.Println("Port: ", config.Port)
	log.Println("Videodir: ", config.Videodir)
	log.Println("Options: ", config.Options)
	log.Println("Videoformat:", config.Videoformat)

	log.Println("config reading OK")

}
