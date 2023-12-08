#include <iostream>
#include <fstream>
#include "./AuxFunctions/ElfSackInterpreter.hpp"
#include "./AuxFunctions/Helpers.hpp"

// To run my answer do:
    // make clear && make day2 && ./day2 ./Input/input.txt 12r13g14b
int main(int argc, char* argv[])
{
    // 2nd argument should be file name with path
    // 3rd argument should be the elf's sack configuration in the format:
        // 12r13g14b (number + r for red, g for green, b for blue)
    if (argc < 3)
        return 1;
    ElfSackInterpreter elfSack(argv[1]);

    // Defining regex's to parse the bag configuration in the input
    std::string configuration = argv[2];
    std::regex redConfigurationRegex("[\\d+\\s+]+(?=r)");
    std::regex greenConfigurationRegex("[\\d+\\s+]+(?=g)");
    std::regex blueConfigurationRegex("[\\d+\\s+]+(?=b)");
    // Construct the request tuple
    auto redMax = RegexHelper::returnMaxUsingPattern(redConfigurationRegex, argv[2]);
    auto greenMax = RegexHelper::returnMaxUsingPattern(greenConfigurationRegex, argv[2]);
    auto blueMax = RegexHelper::returnMaxUsingPattern(blueConfigurationRegex, argv[2]);
    std::cout << "redMax : " << redMax << std::endl;
    std::cout << "greenMax : " << greenMax << std::endl;
    std::cout << "blueMax : " << blueMax << std::endl;
    GameQuestionRGB request{redMax, greenMax, blueMax};
    std::cout << "Answer : " << elfSack.sumOfPossibleGameIds(request) << std::endl;
    return 0;
}