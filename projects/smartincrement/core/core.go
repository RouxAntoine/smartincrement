package core

import (
	"time"
	"os"
	"io/ioutil"
	"log"
	"encoding/json"
	"strconv"
)

const DEFAULT = 10

const ctLayout = "2006-01-02T15:04:05Z07:00"

type SmartIncrement struct {
	data jsonData
	databaseFile string
	interval float64
	slow int
	fast int
}

type jsonData struct {
	Date  time.Time	`json:"date"`
	Value int		`json:"value"`
	IsInc bool		`json:"is_inc"`
}

// New constructeur
func New(file string, interval float64, fast int, slow int, initValue ...int) (e SmartIncrement) {
	e = SmartIncrement{}
	e.slow = slow
	e.fast = fast
	e.interval = interval

	e.databaseFile = file

	valueBytes := e.readCurrent()
	e.parse(valueBytes)

	if len(initValue) > 0 && initValue[0] > -1 {
		e.data.Value = initValue[0]
	}

	//valueInt, errConv := strconv.Atoi(string(valueBytes))
	//if errConv != nil {
	//	fmt.Printf("erreur durant la récupération de la valeur actuel définition a %d\n", DEFAULT)
	//}
	//e.Value = valueInt

	return
}

// before : avant
// after : après
func (si *SmartIncrement) NextValue() int {
	intervalLastRun := time.Now().Local().Sub(si.data.Date).Seconds()

	log.Printf("next: valeur now - time en sec : %f", intervalLastRun)
	// plus de 2 seconds que l'on a appuié
	if  si.data.IsInc && intervalLastRun > si.interval {
		si.data.Value += si.slow
	} else {
		si.data.Value += si.fast
	}

	si.data.IsInc = true
	si.data.Date = time.Now().Local()
	return si.data.Value
}

// before : avant
// after : après
func (si *SmartIncrement) PrevValue() int {
	intervalLastRun := time.Now().Local().Sub(si.data.Date).Seconds()

	log.Printf("prev: valeur now - time en sec: %f", intervalLastRun)

	// plus de 2 seconds que l'on a appuié
	if !si.data.IsInc && intervalLastRun > si.interval {
		si.data.Value -= si.slow
	} else {
		si.data.Value -= si.fast
	}

	si.data.IsInc = false
	si.data.Date = time.Now().Local()
	return si.data.Value
}

func (si *SmartIncrement) Persist() {
	err := ioutil.WriteFile(si.databaseFile, si.stringify(), 0644)
	if err != nil {
		log.Output(2, err.Error())
	}
}

func FormatDate(d time.Time) string {
	return d.Local().Format(ctLayout)
}

func (si *SmartIncrement) stringify() (b []byte) {
	b, err := json.Marshal(si.data)

	if err != nil {
		log.Output(2, "Error object to json " + err.Error())
	}
	return
}

func (si *SmartIncrement) parse(datum []byte) {
	err := json.Unmarshal(datum, &si.data)
	if err != nil {
		log.Output(2, "Error parsing json database " + err.Error())
	}
}

// read the database file for get the current value and date
func (si *SmartIncrement) readCurrent() (v []byte) {
	v, err := ioutil.ReadFile(si.databaseFile) // just pass the file name
	if err != nil {
		log.Output(2, "Error : failed to open file : "+err.Error())
		if os.IsNotExist(err) {
			v = []byte("{\"date\": \"" + FormatDate(time.Now().Local()) + "\", \"value\": "+strconv.Itoa(DEFAULT)+", \"is_inc\": false}")
		}
	}

	return
}
