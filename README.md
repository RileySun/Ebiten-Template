# Ebiten Template
This is my ebiten engine template for making games with the [Ebiten Engine](https://github.com/hajimehoshi/ebiten).  
This is the boilerplate code that I end up always using every time I make a game.  

### Config Info
Inside [utils.go](https://github.com/RileySun/Ebiten-Template/blob/main/utils.go#L19) there are the 3 lines near the top of the file that are specific to each game:
	`GAMENAME = "Ebiten Template"` -> Game Name (Appears on game window)
	`GAMEID = "com.organization.project"` -> Game Identifier (Used for app data storage)
	`GAMEWIDTH = 640` -> Game Width (Non-Fullscreen)
	`GAMEHEIGHT = 520` -> Game Height (Non-Fullscreen)
