#include <iostream>
#include <fstream>
#include <vector>

std::vector<std::vector<std::string>> readFile() {
	std::fstream input;
	std::string line;
	input.open("../inputs/input-d3.txt");
    std::vector<std::vector<std::string>> rucksacks;

	if (input.is_open()) {
        std::vector<std::string> rucksack;
		while ( getline(input, line)) {
			rucksack.push_back(line.substr(0, line.size() / 2));
			rucksack.push_back(line.substr(line.size() / 2));
			rucksacks.push_back(rucksack);
			rucksack.clear();
		}
		input.close();
	}

    return rucksacks;
}

int priority(char letra) {
	if (int(letra) >= int('A') and int(letra) <= int('Z')) {
		return int(letra) - 38;
	}

	return int(letra) - 96;
}

char findCommon(std::vector<std::string> rucksack) {
	bool map1[52] = {false}, map2[52] = {false};
	for (int i = 0; i < rucksack[0].size(); ++i) {
		map1[priority(rucksack[0][i]) - 1] = true;
	}
	for (int i = 0; i < rucksack[1].size(); ++i) {
		map2[priority(rucksack[1][i]) - 1] = true;
	}
	for (int i = 0; i < 52; ++i) {
		if (map1[i] && map2[i]) {
			return i + 1;
		}
	}
	return 0;
}

int partOne(std::vector<std::vector<std::string>> rucksacks) {
	int sum = 0;
	for (int i = 0; i < rucksacks.size(); ++i) {
		sum += findCommon(rucksacks.at(i));
	}
	return sum;
}

int main() {
	std::vector<std::vector<std::string>> rucksacks = readFile();

	std::cout << "Part one: " << partOne(rucksacks) << "\n";

	return 0;
}