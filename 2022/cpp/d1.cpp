#include <iostream>
#include <fstream>
#include <vector>

std::vector<std::vector<int>> readFile() {
	std::fstream input;
	std::string line;
	input.open("../inputs/input-d1.txt");
    std::vector<std::vector<int>> elfs;

	if (input.is_open()) {
        std::vector<int> elf;
		while ( getline(input, line)) {
			if (! line.size()) {
                elfs.push_back(elf);
                elf.clear();
			} else {
				elf.push_back(stoi(line));
			}
		}
		input.close();
	}

    return elfs;
}

int partOne(std::vector<std::vector<int>> elfs) {
    int current = 0, max = 0;

    for (int i = 0; i < elfs.size(); ++i) {
        for (int j = 0; j < elfs.at(i).size(); ++j) {
            current += elfs.at(i).at(j);
        }
        if (current > max) {
            max = current;
        }
        current = 0;
    }

    return max;
}

int partTwo(std::vector<std::vector<int>> elfs) {
    int current = 0, max[3] = {0, 0, 0};

    for (int i = 0; i < elfs.size(); ++i) {
        for (int j = 0; j < elfs.at(i).size(); ++j) {
            current += elfs.at(i).at(j);
        }

        int m = 0;
        while (current != 0) {
            if (current > max[m]) {
                max[m] = current;
                current = 0;
            }
            ++m;
            if (m == sizeof(max)/sizeof(int)) {
                current = 0;
            }
        }
    }

    int sum = 0;
    for (int i = 0; i < sizeof(max)/sizeof(int); ++i) {
        sum += max[i];
    }
    return sum;
}

int main() {
    std::vector<std::vector<int>> elfs = readFile();

    std::cout << "Part one: " << partOne(elfs) << "\nPart two: " << partTwo(elfs) << "\n";
    
	return 0;
}
