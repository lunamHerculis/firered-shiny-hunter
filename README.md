# firered-shiny-hunter

A cli tool for automatically grinding shiny pokemon in Pokémon FireRed and LeafGreen.

Watch this tutorial video to see how it works: TODO

# Running the tool

Start the tool using this command:

For mac with arm processor:
```bash
./firered-shiny-hunter-mac-arm --pokemon "starter"
```

For windows with x64 processor:
```bash
firered-shiny-hunter-windows-x64.exe --pokemon "starter"
```

You set the emulator speed with the `--ffspeed` flag. It must match the speed set in the emulator settings.
The standard is --ffspeed 6. So if you want to run the emulator at normal speed, you must set the flag to 1:

```bash
./firered-shiny-hunter-mac-arm --pokemon "starter" --ffspeed 1
```

The tool supports the following pokémon:
* "starter": Bulbasaur, Charmander, or Squirtle
* "legendary": Articuno, Zapdos, Moltres, or Mewtu
* "eevee": Eevee gifted in Prismania City
* "magikarp" Magikarp bought from the salesman on Route 4
* "lapras" Lapras gifted in Silph Co.
* "hitmon" Hitmonlee or Hitmonchan gifted in the Fighting Dojo in Saffron City
* "casino1" Abra bought in the casino
* "casino2" Clefairy bought in the casino
* "casino3" Dratini bought in the casino
* "casino4" Scyther bought in the casino
* "casino5" Porygon bought in the casino

# Instructions

The tool works by soft-resetting the game repeatedly and performing the button presses to receive the pokémon again.
After that, it will compare the color of a pixels on the screen you selected for changes.

Instructions to use the tool:
* After starting the tool, you get a welcome screen for 5 seconds.
* After that, you have 5 seconds to have the emulator window active and focused and have the cursor at a pixel that is 
going to have a different color when a shiny pokémon appears. The easiest way to do this, is to first have the terminal 
window focused and then click on the emulator window at the pixel you selected to focus it, without moving the mouse anymore.
* After that, you get 5 seconds to select a pixel that will stay the same for every encounter. This is used to detect if 
the game got to the correct point again. The easiest way is again to first defocuse the emulator window by clicking on 
the terminal window and then clicking on the emulator window at the pixel you selected to focus it, without moving the 
mouse anymore. The best choice are pixels that are part of the pokémon's name.
* You repeat the last step again twice. So in total you select 4 pixels: one that changes when a shiny pokémon appears 
and three that stay the same for every encounter.

Common problems and solutions:
* The terminal window must be allowed to monitor the display pixels and must be allowed to control the keyboard.
* Fast-Forward speeds above 6x become very unstable and can lead to false positives. It's recommended to use speeds 
between 1x and 6x.