#include "ElfSackInterpreter.hpp"
#include "Helpers.hpp"

ElfSackInterpreter::ElfSackInterpreter(const std::string& fileName)
{
    // Define a regular expression to get: GameId, red amount, green amount, blue amount
    // This is not going to be very efficient since std::regex is known to have performance issues
    // Will use std::regex on a per-line base to get all the numbers I need
    std::regex gameId("[\\d+\\s+]+(?=\\:)");
    std::regex redRegex("[\\d+\\s+]+(?=red)");
    std::regex greenRegex("[\\d+\\s+]+(?=green)");
    std::regex blueRegex("[\\d+\\s+]+(?=blue)");

    // Read the file
    std::ifstream fileContent;
    fileContent.open(fileName);

    // Replace for an early return
    assert(fileContent.is_open());

    // Reading file line by line
    std::string line;
    while(getline(fileContent, line))
    {
        auto id = RegexHelper::returnMaxUsingPattern(gameId, line);
        auto redMax = RegexHelper::returnMaxUsingPattern(redRegex, line);
        auto greenMax = RegexHelper::returnMaxUsingPattern(greenRegex, line);
        auto blueMax = RegexHelper::returnMaxUsingPattern(blueRegex, line);
        std::cout << id << " : " << redMax << " : " << greenMax << " : " << blueMax << std::endl;
        m_gameIdToColorsMax[id][CubeColors::Red] = redMax;
        m_gameIdToColorsMax[id][CubeColors::Green] = greenMax;
        m_gameIdToColorsMax[id][CubeColors::Blue] = blueMax;
    }

}

int ElfSackInterpreter::sumOfPossibleGameIds(GameQuestionRGB bagConfiguration)
{
    // Final result accumulator
    int accum = 0;
    // Getting the max values from the request tuple
    auto redMax = std::get<0>(bagConfiguration);
    auto greenMax = std::get<1>(bagConfiguration);
    auto blueMax = std::get<2>(bagConfiguration);
    // Verify the valid games ids and sum them up
    for (auto& [gameId, colorMap] : m_gameIdToColorsMax)
    {
        if (colorMap[CubeColors::Red] > redMax || colorMap[CubeColors::Green] > greenMax || colorMap[CubeColors::Blue] > blueMax)
        {
            continue;
        }
        accum += gameId;
    }
    return accum;
}