#include <iostream>
#include "crypto_functions.cpp"
#include "functions.cpp"


void parseCommands(string command, int argc, string args[]) {
	if (command == "add") {
		if (argc < 1) {
			printMenu();
		} else {
			add(args[0]);
		}
	} else if (command == "commit") {
		if (argc < 1) {
			printMenu();
		} else {
			commit(args[0]);
		}
	}
}

int main(int argc, char* argv[]) {
	if (argc < 2) {
		printMenu();
	} else {
		string args[argc - 2];
		for (int i = 0; i < argc - 2; ++i) {
			args[i] = argv[i + 2];
		}
		parseCommands(argv[1], argc - 2, args);
	}
	return 0;
}