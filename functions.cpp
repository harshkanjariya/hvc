#include <filesystem>
#include <algorithm>

namespace fs = std::filesystem;

void printMenu() {
	cout<<"Available commands:"<<endl;
	cout<<"\tadd <file/folder>\t\t: Add file/folder to commit"<<endl;
	cout<<"\tcommit <message>\t\t: Commit the changes to repo"<<endl;
}

void addFileToObjects(string path) {
	cout<<"adding="<<path<<endl;
	string hash = sha256(path);
	cout<<"hash="<<hash<<endl;
}

void add(string pathStr) {
	const fs::path p(pathStr);
	error_code ec;
	string finalPath;


	if (fs::is_directory(p, ec)) {
		for (fs::directory_iterator i(pathStr), end; i != end; ++i) {
			string subPath(i->path());
			if (subPath == "./.hvc" || subPath == ".hvc") {
				continue;
			}
			if (is_directory(i->path())) {
				for (fs::recursive_directory_iterator i(subPath), end; i != end; ++i) {
					if (!is_directory(i->path())) {
						addFileToObjects(i->path());
					}
				}
			} else {
				addFileToObjects(subPath);
			}
		}
	}
	if (ec)	{
		cout<<"Error:"<<ec.message()<<endl;
	}
	if (fs::is_regular_file(p, ec))	{
		addFileToObjects(pathStr);
	}
	if (ec)	{
		cout<<"Error:"<<ec.message()<<endl;
	}
}

void commit(string message) {
	cout<<"commiting:"<<message;
}

