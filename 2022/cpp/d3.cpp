#include <iostream>
#include <fstream>
#include <vector>
#include <cstdio>

std::vector<std::string> readFile() {
	std::fstream input;
	std::string line;
	input.open("/home/jucapoid/aoc/2022/inputs/input-d3.txt");
    std::vector<std::string> rucksacks;

	if (input.is_open()) {
		while ( getline(input, line)) {
			rucksacks.push_back(line);
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

char letter(int priority) {
	char letra;
	if (priority <= 26) {
		sprintf(&letra, "%c", priority + 96);
	} else {
		sprintf(&letra, "%c", priority + 38);
	}
	
	return letra;
}

char findCommon(std::string rucksack) {
	std::string firstHalf = rucksack.substr(0, rucksack.size() / 2);
	std::string secondHalf = rucksack.substr(rucksack.size() / 2);
	bool map1[52] = {false}, map2[52] = {false};
	for (int i = 0; i < firstHalf.size(); ++i) {
		map1[priority(firstHalf[i]) - 1] = true;
	}
	for (int i = 0; i < secondHalf.size(); ++i) {
		map2[priority(secondHalf[i]) - 1] = true;
	}
	for (int i = 0; i < 52; ++i) {
		if (map1[i] && map2[i]) {
			return i + 1;
		}
	}
	return 0;
}

char findCommonPartTwo(std::string rucksack1, std::string rucksack2, std::string rucksack3) {
	bool map1[52] = {false}, map2[52] = {false}, map3[52] = {false};
	for (int i = 0; i < rucksack1.size(); ++i) {
		map1[priority(rucksack1[i]) - 1] = true;
	}
	for (int i = 0; i < rucksack2.size(); ++i) {
		map2[priority(rucksack2[i]) - 1] = true;
	}
	for (int i = 0; i < rucksack3.size(); ++i) {
		map3[priority(rucksack3[i]) - 1] = true;
	}

	for (int i = 0; i < 52; ++i) {
		if (map1[i] && map2[i] && map3[i]) {
			std::cout << rucksack1 << " " << rucksack2 << " " << rucksack3 << " common " << letter(i + 1) << "\n";
			return i + 1;
		}
	}
	return 0;
}

int partOne(std::vector<std::string> rucksacks) {
	int sum = 0;
	for (int i = 0; i < rucksacks.size(); ++i) {
		sum += findCommon(rucksacks.at(i));
	}
	return sum;
}

int partTwo(std::vector<std::string> rucksacks) {
	int sum = 0;
	for (int i = 0; i < rucksacks.size() - 1; i+=3) {
		sum += findCommonPartTwo(rucksacks.at(i), rucksacks.at(i+1), rucksacks.at(i+2));
	}
	return sum;
	
}

int main() {
	std::vector<std::string> rucksacks = readFile();

	std::cout << "Part one: " << partOne(rucksacks) << "\n";

	std::cout << "Part two: " << partTwo(rucksacks) << "\n";

	return 0;
}