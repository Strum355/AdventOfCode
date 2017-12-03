To run a specific day and part, use the following syntax, assuming gradle and kotlin are installed:

X = Day
y = Part

FISH: 
```fish
gradle fatjar; and kotlin -cp build/libs/Kotlin.jar DayX_partYKt
```

BASH:
```bash
gradle fatjar && kotlin -cp build/libs/Kotlin.jar DayX_partYKt
```
