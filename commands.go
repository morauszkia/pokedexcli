package main

type cliCommand struct {
	name		string
	description	string
	callback	func(*Config, ...string) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:			"help",
			description: 	"Lists available commands",
			callback: 		commandHelp,
		},
		"exit": {
			name: 			"exit",
			description:	"Exits the Pokedex",
			callback:		commandExit,
		},
		"map": {
			name: 			"map",
			description: 	"Lists next 20 location areas",
			callback: 		commandMap,
		},
		"mapb": {
			name: 			"mapb",
			description:	"Lists previous 20 location areas",
			callback: 		commandMapb,
		},
		"explore": {
			name:			"explore",
			description:	"Look for pokemons in a location area",
			callback:		commandExplore,
		},
		"catch": {
			name:           "catch",
			description:    "Attempt to catch a pokemon",
			callback:       commandCatch,
		},
		"inspect": {
			name:			"inspect",
			description: 	"Show detailed information about caught pokemon",
			callback: 		commandInspect,
		},
		"pokedex": {
			name: 			"pokedex",
			description: 	"List the pokemon you've caught",
			callback: 		commandPokedex,
		},
	}
}