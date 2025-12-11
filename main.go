package main

import (
	"flag"
	"fmt"
	"slices"
	"strings"
	"time"

	"github.com/go-vgo/robotgo"
)

var (
	pokemon string

	fastForwardSpeed int
	displayID        int

	keyA         string
	keyB         string
	keyUp        string
	keyDown      string
	keyLeft      string
	keyRight     string
	keyStart     string
	keySelect    string
	keyLShoulder string
	keyRShoulder string
)

const (
	layoutUS = "2006-01-02 15:04:05"
)

var (
	validPokemon = []string{
		"starter",
		"legendary",
		"eevee",
		"magikarp",
		"lapras",
		"hitmon",
		"casino1",
		"casino2",
		"casino3",
		"casino4",
		"casino5",
	}

	nonShinyHex      = ""
	colorShouldStay1 = ""
	colorShouldStay2 = ""
	colorShouldStay3 = ""

	shinyLocationX            = 0
	shinyLocationY            = 0
	colorShouldStay1LocationX = 0
	colorShouldStay1LocationY = 0
	colorShouldStay2LocationX = 0
	colorShouldStay2LocationY = 0
	colorShouldStay3LocationX = 0
	colorShouldStay3LocationY = 0

	saveCounter = 0
	realCounter = 0
)

func main() {
	flag.StringVar(&pokemon, "pokemon", "starter", "Select which pokemon to shinyhunt: "+strings.Join(validPokemon, ", "))
	flag.IntVar(&fastForwardSpeed, "ffspeed", 6, "Fast forward speed multiplier")
	flag.IntVar(&displayID, "displayID", 0, "Display ID for pixel color detection")
	flag.StringVar(&keyA, "keyA", "g", "Key mapping for A button")
	flag.StringVar(&keyB, "keyB", "f", "Key mapping for B button")
	flag.StringVar(&keyUp, "keyUp", "w", "Key mapping for Up button")
	flag.StringVar(&keyDown, "keyDown", "s", "Key mapping for Down button")
	flag.StringVar(&keyLeft, "keyLeft", "a", "Key mapping for Left button")
	flag.StringVar(&keyRight, "keyRight", "s", "Key mapping for Right button")
	flag.StringVar(&keyStart, "keyStart", "t", "Key mapping for Start button")
	flag.StringVar(&keySelect, "keySelect", "r", "Key mapping for Select button")
	flag.StringVar(&keyLShoulder, "keyLShoulder", "q", "Key mapping for L Shoulder button")
	flag.StringVar(&keyRShoulder, "keyRShoulder", "e", "Key mapping for R Shoulder button")
	flag.Parse()

	if !slices.Contains(validPokemon, pokemon) {
		fmt.Println("Invalid pokemon. Please use one of the following: " + strings.Join(validPokemon, ", "))
		return
	}

	hello := `
     _     _               _                 _            
    | |   (_)             | |               | |           
 ___| |__  _ _ __  _   _  | |__  _   _ _ __ | |_ ___ _ __ 
/ __| '_ \| | '_ \| | | | | '_ \| | | | '_ \| __/ _ \ '__|
\__ \ | | | | | | | |_| | | | | | |_| | | | | ||  __/ |   
|___/_| |_|_|_| |_|\__, | |_| |_|\__,_|_| |_|\__\___|_|   
                    __/ |                                 
                   |___/
    `
	fmt.Println(hello)

	instr1 := `
--------------- INSTRUCTION ------------------------------
select a pixel that changes for a shiny
----------------------------------------------------------
    `
	fmt.Println(instr1)

	time.Sleep(1 * time.Second)
	fmt.Println("(5)")
	time.Sleep(1 * time.Second)
	fmt.Println("(4)")
	time.Sleep(1 * time.Second)
	fmt.Println("(3)")
	time.Sleep(1 * time.Second)
	fmt.Println("(2)")
	time.Sleep(1 * time.Second)
	fmt.Println("(1)")
	time.Sleep(1 * time.Second)

	shinyLocationX, shinyLocationY = robotgo.Location()
	fmt.Printf("locationX: %d, locationY: %d\n", shinyLocationX, shinyLocationY)
	nonShinyHex = robotgo.GetPixelColor(shinyLocationX, shinyLocationY, displayID)
	fmt.Printf("nonShinyHex: %s\n", nonShinyHex)

	instr2 := `
--------------- INSTRUCTION ------------------------------
(1) select a pixel that should stay the same
----------------------------------------------------------
    `
	fmt.Println(instr2)

	time.Sleep(1 * time.Second)
	fmt.Println("(5)")
	time.Sleep(1 * time.Second)
	fmt.Println("(4)")
	time.Sleep(1 * time.Second)
	fmt.Println("(3)")
	time.Sleep(1 * time.Second)
	fmt.Println("(2)")
	time.Sleep(1 * time.Second)
	fmt.Println("(1)")
	time.Sleep(1 * time.Second)

	colorShouldStay1LocationX, colorShouldStay1LocationY = robotgo.Location()
	fmt.Printf("locationX: %d, locationY: %d\n", colorShouldStay1LocationX, colorShouldStay1LocationY)
	colorShouldStay1 = robotgo.GetPixelColor(colorShouldStay1LocationX, colorShouldStay1LocationY, displayID)
	fmt.Printf("colorShouldStay: %s\n", colorShouldStay1)

	instr3 := `
--------------- INSTRUCTION ------------------------------
(2) select a pixel that should stay the same
----------------------------------------------------------
    `
	fmt.Println(instr3)

	time.Sleep(1 * time.Second)
	fmt.Println("(5)")
	time.Sleep(1 * time.Second)
	fmt.Println("(4)")
	time.Sleep(1 * time.Second)
	fmt.Println("(3)")
	time.Sleep(1 * time.Second)
	fmt.Println("(2)")
	time.Sleep(1 * time.Second)
	fmt.Println("(1)")
	time.Sleep(1 * time.Second)

	colorShouldStay2LocationX, colorShouldStay2LocationY = robotgo.Location()
	fmt.Printf("locationX: %d, locationY: %d\n", colorShouldStay2LocationX, colorShouldStay2LocationY)
	colorShouldStay2 = robotgo.GetPixelColor(colorShouldStay2LocationX, colorShouldStay2LocationY, displayID)
	fmt.Printf("colorShouldStay2: %s\n", colorShouldStay2)

	instr4 := `
--------------- INSTRUCTION ------------------------------
(3) select a pixel that should stay the same
----------------------------------------------------------
    `
	fmt.Println(instr4)

	time.Sleep(1 * time.Second)
	fmt.Println("(5)")
	time.Sleep(1 * time.Second)
	fmt.Println("(4)")
	time.Sleep(1 * time.Second)
	fmt.Println("(3)")
	time.Sleep(1 * time.Second)
	fmt.Println("(2)")
	time.Sleep(1 * time.Second)
	fmt.Println("(1)")
	time.Sleep(1 * time.Second)

	colorShouldStay3LocationX, colorShouldStay3LocationY = robotgo.Location()
	fmt.Printf("locationX: %d, locationY: %d\n", colorShouldStay3LocationX, colorShouldStay3LocationY)
	colorShouldStay3 = robotgo.GetPixelColor(colorShouldStay3LocationX, colorShouldStay3LocationY, displayID)
	fmt.Printf("colorShouldStay3: %s\n", colorShouldStay3)

	doneStr := `
--------------- DONE ------------------------------
    `
	fmt.Println(doneStr)

	catchEmAll := `
Catch 'em All!
    `
	fmt.Println(catchEmAll)
	time.Sleep(1 * time.Second)

	for {
		softReset()
		goIntoGame()

		newAdditionalWaitTime := 20 * saveCounter // roughly 1.2 frames * saveCounter
		time.Sleep(time.Duration(newAdditionalWaitTime) * time.Millisecond)

		// after 100 runs, wait and save
		if saveCounter >= 100 {

			printStatus("saving...")
			save()

			saveCounter = 0
			continue
		}

		realCounter++
		saveCounter++

		switch pokemon {
		case "starter":
			takeStarter()
			goIntoPokedexStarter()
		case "legendary":
			takeLegendary()
		case "eevee":
			takeEevee()
			goIntoPokedexSixthPosition()
		case "magikarp":
			takeMagikarp()
			goIntoPokedexSixthPosition()
		case "lapras":
			takeLapras()
			goIntoPokedexSixthPosition()
		case "hitmon":
			takeHitmon()
			goIntoPokedexSixthPosition()
		case "casino1":
			takeCasino1()
			goIntoPokedexSixthPosition()
		case "casino2":
			takeCasino2()
			goIntoPokedexSixthPosition()
		case "casino3":
			takeCasino3()
			goIntoPokedexSixthPosition()
		case "casino4":
			takeCasino4()
			goIntoPokedexSixthPosition()
		case "casino5":
			takeCasino5()
			goIntoPokedexSixthPosition()
		default:
			fmt.Println("invalid pokemon")
			return
		}

		shiny := robotgo.GetPixelColor(shinyLocationX, shinyLocationY, displayID)
		stay1 := robotgo.GetPixelColor(colorShouldStay1LocationX, colorShouldStay1LocationY, displayID)
		stay2 := robotgo.GetPixelColor(colorShouldStay2LocationX, colorShouldStay2LocationY, displayID)
		stay3 := robotgo.GetPixelColor(colorShouldStay3LocationX, colorShouldStay3LocationY, displayID)
		if shiny == nonShinyHex {
			printStatus(fmt.Sprintf("no shiny, shinyColor: %s, addWaitTime:  %.3f, realCounter: %d, saveCounter: %d", shiny, float64(newAdditionalWaitTime)/1000, realCounter, saveCounter))
			continue
		}

		if stay1 != colorShouldStay1 || stay2 != colorShouldStay2 || stay3 != colorShouldStay3 {
			printStatus(fmt.Sprintf("something went wrong i think. shinyColor changed but shouldStayColor changed too, shinyColor: %s ... reset", shiny))
			continue
		}

		fmt.Printf("holy shit a shiny!!!, nonShinyHex: %s, shinyColor: %s, realCounter: %d,  saveCounter: %d\n time: %s\n", nonShinyHex, shiny, realCounter, saveCounter, time.Now().Format(layoutUS))
		return
	}
}

func softReset() {
	keyPressDown(keyStart)
	keyPressDown(keySelect)
	keyPressDown(keyA)
	keyPressDown(keyB)
	time.Sleep(100 * time.Millisecond)
	keyPressUp(keyStart)
	keyPressUp(keySelect)
	keyPressUp(keyA)
	keyPressUp(keyB)
	time.Sleep(10 * time.Millisecond)
}

func save() {
	time.Sleep(1 * time.Second / time.Duration(fastForwardSpeed))
	keyStroke(keyStart)
	time.Sleep(1 * time.Second / time.Duration(fastForwardSpeed))
	keyStroke(keyDown)
	time.Sleep(1 * time.Second / time.Duration(fastForwardSpeed))
	keyStroke(keyDown)
	time.Sleep(1 * time.Second / time.Duration(fastForwardSpeed))
	keyStroke(keyDown)
	time.Sleep(1 * time.Second / time.Duration(fastForwardSpeed))
	keyStroke(keyDown)
	time.Sleep(1 * time.Second / time.Duration(fastForwardSpeed))
	keyStroke(keyA)
	time.Sleep(2 * time.Second / time.Duration(fastForwardSpeed))
	keyStroke(keyA)
	time.Sleep(2 * time.Second / time.Duration(fastForwardSpeed))
	keyStroke(keyA)
	time.Sleep(9 * time.Second / time.Duration(fastForwardSpeed))
}

func goIntoGame() {
	time.Sleep(time.Second * 4 / time.Duration(fastForwardSpeed))
	keyStroke(keyA)
	time.Sleep(time.Second * 5 / time.Duration(fastForwardSpeed))
	keyStroke(keyA)
	time.Sleep(time.Second * 4 / time.Duration(fastForwardSpeed))
	keyStroke(keyA)
	time.Sleep(time.Second * 4 / time.Duration(fastForwardSpeed))
	keyStroke(keyB)
	time.Sleep(time.Second * 3 / time.Duration(fastForwardSpeed))
}

func goIntoPokedexSixthPosition() {
	keyStroke(keyStart)
	time.Sleep(1 * time.Second / time.Duration(fastForwardSpeed))
	keyStroke(keyDown)
	time.Sleep(1 * time.Second / time.Duration(fastForwardSpeed))
	keyStroke(keyA)
	time.Sleep(2 * time.Second / time.Duration(fastForwardSpeed))
	keyStroke(keyUp)
	time.Sleep(1 * time.Second / time.Duration(fastForwardSpeed))
	keyStroke(keyUp)
	time.Sleep(1 * time.Second / time.Duration(fastForwardSpeed))
	keyStroke(keyA)
	time.Sleep(1 * time.Second / time.Duration(fastForwardSpeed))
	keyStroke(keyA)
	time.Sleep(3 * time.Second / time.Duration(fastForwardSpeed))
}

func keyPressDown(key string) {
	err := robotgo.KeyDown(key)
	if err != nil {
		fmt.Printf("error pressing key %s in keyPressDown: %v\n", key, err)
	}
}

func keyPressUp(key string) {
	err := robotgo.KeyUp(key)
	if err != nil {
		fmt.Printf("error pressing key %s in keyPressUp: %v\n", key, err)
	}
}

func keyStroke(key string) {
	err := robotgo.KeyDown(key)
	if err != nil {
		fmt.Printf("error pressing key %s in keyStroke: %v\n", key, err)
		return
	}
	time.Sleep(80 * time.Millisecond)
	err = robotgo.KeyUp(key)
	if err != nil {
		fmt.Printf("error releasing key %s in keyStroke: %v\n", key, err)
		return
	}
	time.Sleep(10 * time.Millisecond)
}

func printStatus(msg string) {
	fmt.Println(msg)
}

// ################ Starter ################

func takeStarter() {
	keyStroke(keyA)
	time.Sleep(2 * time.Second / time.Duration(fastForwardSpeed))
	keyStroke(keyA)
	time.Sleep(2 * time.Second / time.Duration(fastForwardSpeed))
	keyStroke(keyA)
	time.Sleep(2 * time.Second / time.Duration(fastForwardSpeed))
	keyStroke(keyA)
	time.Sleep(5 * time.Second / time.Duration(fastForwardSpeed))
	keyStroke(keyB)
	time.Sleep(3 * time.Second / time.Duration(fastForwardSpeed))
	keyStroke(keyA)
	time.Sleep(2 * time.Second / time.Duration(fastForwardSpeed))
	keyStroke(keyA)
	time.Sleep(2 * time.Second / time.Duration(fastForwardSpeed))
}

func goIntoPokedexStarter() {
	keyStroke(keyStart)
	time.Sleep(1 * time.Second / time.Duration(fastForwardSpeed))
	keyStroke(keyA)
	time.Sleep(1 * time.Second / time.Duration(fastForwardSpeed))
	keyStroke(keyA)
	time.Sleep(1 * time.Second / time.Duration(fastForwardSpeed))
	keyStroke(keyA)
	time.Sleep(3 * time.Second / time.Duration(fastForwardSpeed))
}

// ################ Legendary ################

func takeLegendary() {
	keyStroke(keyA)
	time.Sleep(2 * time.Second / time.Duration(fastForwardSpeed))
	keyStroke(keyA)
	time.Sleep(6 * time.Second / time.Duration(fastForwardSpeed))
}

// ################ Eevee ################

func takeEevee() {
	keyStroke(keyA)
	time.Sleep(3 * time.Second / time.Duration(fastForwardSpeed))
	keyStroke(keyB)
	time.Sleep(1 * time.Second / time.Duration(fastForwardSpeed))
}

// ################ Magikarp ################

func takeMagikarp() {
	keyStroke(keyA)
	time.Sleep(2 * time.Second / time.Duration(fastForwardSpeed))
	keyStroke(keyA)
	time.Sleep(2 * time.Second / time.Duration(fastForwardSpeed))
	keyStroke(keyA)
	time.Sleep(2 * time.Second / time.Duration(fastForwardSpeed))
	keyStroke(keyA)
	time.Sleep(2 * time.Second / time.Duration(fastForwardSpeed))
	keyStroke(keyA)
	time.Sleep(3 * time.Second / time.Duration(fastForwardSpeed))
	keyStroke(keyB)
	time.Sleep(1 * time.Second / time.Duration(fastForwardSpeed))
}

// ################ Lapras ################

func takeLapras() {
	keyStroke(keyA)
	time.Sleep(2 * time.Second / time.Duration(fastForwardSpeed))
	keyStroke(keyA)
	time.Sleep(2 * time.Second / time.Duration(fastForwardSpeed))
	keyStroke(keyA)
	time.Sleep(2 * time.Second / time.Duration(fastForwardSpeed))
	keyStroke(keyA)
	time.Sleep(4 * time.Second / time.Duration(fastForwardSpeed))
	keyStroke(keyB)
	time.Sleep(2 * time.Second / time.Duration(fastForwardSpeed))
	keyStroke(keyA)
	time.Sleep(2 * time.Second / time.Duration(fastForwardSpeed))
	keyStroke(keyA)
	time.Sleep(2 * time.Second / time.Duration(fastForwardSpeed))
	keyStroke(keyA)
	time.Sleep(2 * time.Second / time.Duration(fastForwardSpeed))
	keyStroke(keyA)
	time.Sleep(2 * time.Second / time.Duration(fastForwardSpeed))
	keyStroke(keyA)
	time.Sleep(1 * time.Second / time.Duration(fastForwardSpeed))
}

// ################ Hitmon ################

func takeHitmon() {
	keyStroke(keyA)
	time.Sleep(3 * time.Second / time.Duration(fastForwardSpeed))
	keyStroke(keyA)
	time.Sleep(4 * time.Second / time.Duration(fastForwardSpeed))
	keyStroke(keyB)
	time.Sleep(1 * time.Second / time.Duration(fastForwardSpeed))
}

// ################ Casino ################

func takeCasino1() {
	keyStroke(keyA)
	time.Sleep(2 * time.Second / time.Duration(fastForwardSpeed))
	keyStroke(keyA)
	time.Sleep(2 * time.Second / time.Duration(fastForwardSpeed))

	keyStroke(keyA)
	time.Sleep(2 * time.Second / time.Duration(fastForwardSpeed))
	keyStroke(keyA)
	time.Sleep(3 * time.Second / time.Duration(fastForwardSpeed))
	keyStroke(keyB)
	time.Sleep(1 * time.Second / time.Duration(fastForwardSpeed))
}

func takeCasino2() {
	keyStroke(keyA)
	time.Sleep(2 * time.Second / time.Duration(fastForwardSpeed))
	keyStroke(keyA)
	time.Sleep(2 * time.Second / time.Duration(fastForwardSpeed))

	keyStroke(keyDown)
	time.Sleep(1 * time.Second / time.Duration(fastForwardSpeed))

	keyStroke(keyA)
	time.Sleep(2 * time.Second / time.Duration(fastForwardSpeed))
	keyStroke(keyA)
	time.Sleep(3 * time.Second / time.Duration(fastForwardSpeed))
	keyStroke(keyB)
	time.Sleep(1 * time.Second / time.Duration(fastForwardSpeed))
}

func takeCasino3() {
	keyStroke(keyA)
	time.Sleep(2 * time.Second / time.Duration(fastForwardSpeed))
	keyStroke(keyA)
	time.Sleep(2 * time.Second / time.Duration(fastForwardSpeed))

	keyStroke(keyDown)
	time.Sleep(1 * time.Second / time.Duration(fastForwardSpeed))
	keyStroke(keyDown)
	time.Sleep(1 * time.Second / time.Duration(fastForwardSpeed))

	keyStroke(keyA)
	time.Sleep(2 * time.Second / time.Duration(fastForwardSpeed))
	keyStroke(keyA)
	time.Sleep(3 * time.Second / time.Duration(fastForwardSpeed))
	keyStroke(keyB)
	time.Sleep(1 * time.Second / time.Duration(fastForwardSpeed))
}

func takeCasino4() {
	keyStroke(keyA)
	time.Sleep(2 * time.Second / time.Duration(fastForwardSpeed))
	keyStroke(keyA)
	time.Sleep(2 * time.Second / time.Duration(fastForwardSpeed))

	keyStroke(keyDown)
	time.Sleep(1 * time.Second / time.Duration(fastForwardSpeed))
	keyStroke(keyDown)
	time.Sleep(1 * time.Second / time.Duration(fastForwardSpeed))
	keyStroke(keyDown)
	time.Sleep(1 * time.Second / time.Duration(fastForwardSpeed))

	keyStroke(keyA)
	time.Sleep(2 * time.Second / time.Duration(fastForwardSpeed))
	keyStroke(keyA)
	time.Sleep(3 * time.Second / time.Duration(fastForwardSpeed))
	keyStroke(keyB)
	time.Sleep(1 * time.Second / time.Duration(fastForwardSpeed))
}

func takeCasino5() {
	keyStroke(keyA)
	time.Sleep(2 * time.Second / time.Duration(fastForwardSpeed))
	keyStroke(keyA)
	time.Sleep(2 * time.Second / time.Duration(fastForwardSpeed))

	keyStroke(keyDown)
	time.Sleep(1 * time.Second / time.Duration(fastForwardSpeed))
	keyStroke(keyDown)
	time.Sleep(1 * time.Second / time.Duration(fastForwardSpeed))
	keyStroke(keyDown)
	time.Sleep(1 * time.Second / time.Duration(fastForwardSpeed))
	keyStroke(keyDown)
	time.Sleep(1 * time.Second / time.Duration(fastForwardSpeed))

	keyStroke(keyA)
	time.Sleep(2 * time.Second / time.Duration(fastForwardSpeed))
	keyStroke(keyA)
	time.Sleep(3 * time.Second / time.Duration(fastForwardSpeed))
	keyStroke(keyB)
	time.Sleep(1 * time.Second / time.Duration(fastForwardSpeed))
}
