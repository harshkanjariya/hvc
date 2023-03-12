#include <filesystem>
#include <algorithm>
#include <fstream>
#include <vector>
#include <zlib.h>

using namespace std;
namespace fs = std::filesystem;

void printMenu() {
	cout<<"Available commands:"<<endl;
	cout<<"\tadd <file/folder>\t\t: Add file/folder to commit"<<endl;
	cout<<"\tcommit <message>\t\t: Commit the changes to repo"<<endl;
}

bool isCompressable(string content) {
	const char* s = content.c_str();
	for (int i = 0; i < content.size(); ++i) {
		if (s[i] & 0x80) {
			return false;
		}
	}
	return true;
}

bool isVectorCompressable(vector<char> v) {
	for (int i = 0; i < v.size(); ++i) {
		if (v[i] & 0x80) {
			return false;
		}
	}
	return true;
}
vector<char> compressVector(vector<char> v) {
	vector<char> compressed;
	return compressed;
}

void addFileToObjects(string path) {
	ifstream stagedIn("./.hvc/staged", ios::binary);
	vector<unsigned char> buf(istreambuf_iterator<char>(stagedIn), {});
	stagedIn.close();

	string s(buf.begin(), buf.end());

	ifstream file(path, ios::binary);
	vector<unsigned char> fileBuf(istreambuf_iterator<char>(file), {});

	ofstream stagedOut("./.hvc/staged", ios::binary);
	if (s.size() == 0) {
		// stagedOut.write();
	}
	// string hash = sha256(path);
	// cout<<"hash="<<hash<<endl;
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

