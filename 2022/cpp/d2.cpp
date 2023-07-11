#include <iostream>
#include <fstream>
#include <vector>
#include <map>

std::vector<std::vector<std::string>> readFile() {
	std::ifstream input;
    input.open("../inputs/input-d2.txt");
    std::string p1, p2;
	std::vector<std::vector<std::string>> lines;
    std::vector<std::string> j;
    if (input.is_open()) {
        while (input >> p1 >> p2) {
			j = {p1, p2};
			lines.push_back(j);
        }
        input.close();
    }

    return lines;	
} 

int getValue(std::string letra) {
    if (letra == "A" or letra == "X") {
        return 1;
    } else if (letra == "B" or letra == "Y") {
        return 2;
    } else if (letra == "C" or letra == "Z") {
        return 3;
    }
    return 0;
};

int getExpected(int p1, std::string p2) {
    if (p2 == "X") {
        return p1 - 1 < 1 ? p1 + 2 : p1 - 1;
    } else if (p2 == "Z") {
        return p1 + 1 > 3 ? p1 - 2 : p1 + 1;
    } else {
        return p1;
    }
}

int pontosNaRonda(int p1, int p2) {
    if (p1 - p2 == -1 or p1 - p2 == 2) {
        return 6 + p2;
    } else if (p1 - p2 == 1 or p1 - p2 == -2) {
        return p2;
    } else {
        return 3 + p2;
    }
}

int partOne(std::vector<std::vector<std::string>> lines) {
	int pontos = 0;

	for (int i = 0; i < lines.size(); ++i) {
		int j1 = getValue(lines.at(i).at(0)), j2 = getValue(lines.at(i).at(1));
		pontos += pontosNaRonda(j1, j2);
	}

	return pontos;
}

int partTwo(std::vector<std::vector<std::string>> lines) {
	int pontos = 0;

	for (int i = 0; i < lines.size(); ++i) {
		int j1 = getValue(lines.at(i).at(0)), j2 = getExpected(j1, lines.at(i).at(1));
		pontos += pontosNaRonda(j1, j2);
	}

	return pontos;
}

int main() {
	std::vector<std::vector<std::string>> lines = readFile();

    std::cout << "Part one: " << partOne(lines) << "\n";
    std::cout << "Part two: " << partTwo(lines) << "\n";
    
    return 0;   
}