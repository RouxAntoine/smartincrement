package main

import (
	"smartincrement/core"
	"flag"
	"fmt"
	"github.com/spf13/viper"
	"path/filepath"
	"log"
	"strconv"
)

const TIME_INTERVAL = 2
const SLOW = 1
const FAST = 5

type flags struct {
	FileDB string
	configFile string
	inc bool
	dec bool
	initValue int
}

func main() {
	flags := parseFlag()

	var interval float64
	var slow, fast int

	if flags.configFile != "" {
		interval, slow, fast = readconfig(flags.configFile)
	} else {
		interval = TIME_INTERVAL
		slow = SLOW
		fast = FAST
	}

	var smartInc core.SmartIncrement
	if f := flag.CommandLine.Lookup("init"); f != nil {
		log.Output(2, "init : " + f.Value.String())
		smartInc = core.New(flags.FileDB, interval, fast, slow, flags.initValue)
	} else {
		smartInc = core.New(flags.FileDB, interval, fast, slow)
	}

	//log.Printf("Current value : %d\n", smartInc.Value)
	//log.Printf("Date %s\n", core.FormatDate(smartInc.Date))
	//log.Printf("New value : %d\n", smartInc.NextValue())

	if flags.inc {
		fmt.Print(smartInc.NextValue())
	} else if flags.dec {
		fmt.Print(smartInc.PrevValue())
	}

	smartInc.Persist()

}

func readconfig(configFile string) (inter float64, slow, fast int) {

	viper.SetConfigType("toml")

	filename :=filepath.Base(configFile)
	extension := filepath.Ext(filename)
	name := filename[0:len(filename)-len(extension)]
	viper.SetConfigName(name) // name of config file (without extension)

	// path to look for the config file in
	base, _ := filepath.Abs(filepath.Dir(configFile))
	viper.AddConfigPath(base)
	viper.AddConfigPath(".")

	err := viper.ReadInConfig() // Find and read the config file
	if err != nil { // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	inter = viper.GetFloat64("interval")
	fast = viper.GetInt("fast")
	slow = viper.GetInt("slow")

	log.Printf("fast : " + strconv.Itoa(fast))
	log.Printf("slow : " + strconv.Itoa(slow))
	log.Printf("interval : " + strconv.FormatFloat(inter, 'g', 1, 64))
	return
}

func parseFlag() (f flags) {
	f = flags{}

	flag.StringVar(&f.configFile, "config", "", "fichier de configuration interval, increment min, increment max")
	flag.StringVar(&f.FileDB, "db", "./smartincrement.db", "fichier vers le chemin de la base de donnée")
	flag.IntVar(&f.initValue, "init", -1, "valeur par defaut à initialiser dans le fichier de base de donnée")
	flag.BoolVar(&f.inc, "inc", false, "on incremente le compteur")
	flag.BoolVar(&f.dec, "dec", false, "on decremente le compteur")

	flag.Parse()

	return
}