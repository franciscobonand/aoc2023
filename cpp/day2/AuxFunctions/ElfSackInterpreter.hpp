#include <unordered_map>
#include <string>
#include <regex>
#include <fstream>
#include <iostream>
#include <tuple>

enum class CubeColors
{
    Red,
    Green,
    Blue
};

using GameId = int;
using ColorMax = int;
using CubeColorsMax = std::unordered_map<CubeColors, ColorMax>;
// Game question - is this game possible? RGB order
using GameQuestionRGB = std::tuple<ColorMax, ColorMax, ColorMax>;

class ElfSackInterpreter
{
    public:
        // Constructor will populate my gameId -> colors count map
        ElfSackInterpreter(const std::string& fileName);
        // API to get the sum of possible games ids
        int sumOfPossibleGameIds(GameQuestionRGB bagConfiguration);
        // API to get the sum of the power of the cubes per game
        int sumPowerOfCubes();
    private:
        // Auxiliary API to get the max value of the color that appeared in the game
        ColorMax getMaxGameValueForColor(GameId id, CubeColors color)
        {
            if (auto gameItr = m_gameIdToColorsMax.find(id); gameItr != m_gameIdToColorsMax.end())
            {
                if (auto colorItr = gameItr->second.find(color); colorItr != gameItr->second.end())
                {
                    return colorItr->second;
                }
                // Color not found
                return 0;
            }
            // Game not found
            return 0;
        }
        // Data structure to keep tabs of game values
        // The first DS (that holds the game id values) did not need to be an unordered_map since input
        // is sorted by id, however in case this changes it will work
        std::unordered_map<GameId, CubeColorsMax> m_gameIdToColorsMax;
};