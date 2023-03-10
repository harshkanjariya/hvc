#include<iostream>
#include "llvm/Support/FileSystem.h"

using namespace std;

void printMenu() {
	cout<<"Available commands:"<<endl;
	cout<<"\tadd <file/folder>\t\t: Add file/folder to commit"<<endl;
	cout<<"\tcommit <message>\t\t: Commit the changes to repo"<<endl;
}

void add(string path) {
	cout<<"adding:"<<path<<endl;
	for (directory_iterator itr(path); itr != end_itr; ++itr) {
		cout<<itr->path()<<endl;
	}
}

void commit(string message) {
	cout<<"commiting:"<<message;
}

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