package main

type cliCommand struct {
	name		string
	description	string
	callback	func(*Config) error
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
	}
}