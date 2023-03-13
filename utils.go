package main

import(
	"os"
	"log"
	"math"
	"bytes"
	"runtime"
	"strconv"
	"math/rand"
	"io/ioutil"
	"path/filepath"
	"image"
	"image/png"
	"image/color"
	"encoding/json"
)

const GAMEID = "com.organization.project"
const GAMENAME = "Ebiten Template"
const GAMEWIDTH = 640
const GAMEHEIGHT = 520

type Data struct {
	//High float64 `json:"high"`
}

var WHITE color.Color

//Color Setup
func init() {
	WHITE = color.NRGBA{R: 254, G: 219, B: 99, A: 255}
}

//Randoms
func getRandomInt(min int, max int) int {
    return min + rand.Intn(max-min)
}

func getRandomRotation() float64 {
	num := getRandomInt(0, 4)
	deg := float64(num * 90)
	return deg * (math.Pi/180.0)
}

func getRandomBool() bool {
	return rand.Intn(2) == 1
}

//Casting
func float2String(input float64) string {
	return strconv.Itoa(int(input))
}

//File/Image Manipulation
func getFile(path string) []byte {
	f, _ := ioutil.ReadFile(path)
	return f
}

func imageFromBytes(byt []byte) image.Image {
	r := bytes.NewReader(byt)
	i, err := png.Decode(r)
	if err != nil {
		log.Fatal("Utils Byt2Img - " + err.Error())
	}
	return i
}

//Stored App Data
func getAppSupportFolder() string {
	var saveDir string
	homeDir, _ := os.UserHomeDir()
	switch runtime.GOOS {
		case "windows":
			saveDir = filepath.Join(filepath.Join(filepath.Join(homeDir, "AppData"), "Roaming"), GAMEID)
			_, err := os.Stat(saveDir)
			if os.IsNotExist(err) {
				dirErr := os.Mkdir(saveDir, 0755)
				if dirErr != nil {
					panic("settings: log file - " + dirErr.Error())
				}
			}
		case "darwin":			
			saveDir = filepath.Join(filepath.Join(filepath.Join(homeDir, "Library"), "Application Support"), GAMEID)
			_, statErr := os.Stat(saveDir)
			if os.IsNotExist(statErr) {
				dirErr := os.Mkdir(saveDir, 0755)
				if dirErr != nil {
					panic("settings: " + dirErr.Error())
				}
			}
		case "linux":
			saveDir = "/var/lib/" + GAMEID
			_, statErr := os.Stat(saveDir)
			if os.IsNotExist(statErr) {
				dirErr := os.Mkdir(saveDir, 0755)
				if dirErr != nil {
					panic("settings: log file - " + dirErr.Error())
				}
			}
	}
	return saveDir
}

func loadData() *Data {
	folderPath := getAppSupportFolder()
	path := folderPath + "/DATA.sun"
	
	_, fileErr := os.Stat(path)
	if os.IsNotExist(fileErr) {
		os.Create(path)
	}
	
	jsonData := getFile(path)
	
	var data *Data
	if len(jsonData) < 1 {
		data = new(Data)
		saveData(data)
	} else {
		jsonErr := json.Unmarshal(jsonData, &data)
		if jsonErr != nil {
			log.Fatal(jsonErr)
		}
	}
	
	return data
}

func saveData(d *Data) {
	json, err := json.MarshalIndent(d, "", "	")
	if err != nil {
		log.Fatal("saveData - ", err)
	}
	
	folderPath := getAppSupportFolder()
	path := folderPath + "/DATA.sun"
	os.WriteFile(path, json, 0755);
}

//Force Quit
func forceQuit() {
	os.Exit(3)
}