#include <string>
#include <iostream>
#include <fstream>
#include <utility>
#define NUMERIC_VALUE_CEILING 10
#define LITERAL_ONE_OR_TWO 12
#define LITERAL_TWO_OR_THREE_OR_EIGHT 238
#define LITERAL_FOUR_OR_FIVE 45
#define LITERAL_SIX_OR_SEVEN 67
#define LITERAL_ONE_OR_THREE_OR_FIVE_OR_EIGHT_OR_NINE 13589
#define LITERAL_SEVEN_OR_NINE 79
#define LITERAL_FOUR 44
#define LITERAL_SIX 66

class CalibrationInterpreter
{
    public:
        CalibrationInterpreter(const std::string& fileName) : m_fileName(fileName) 
        {
            std::cout << "Input file name: " << m_fileName << std::endl;
        };
        int computeSumOfCalibrationValues();
    private:
        /* Private Methods */
        // Helper method to find "one" efficiently
        int findOneInLine(const std::string& line, int position, bool forwardDirection)
        {
            if (forwardDirection)
            {
                // In forward direction, position points to "o", and we want to find "ne"
                auto nPosition = position + 1;
                auto ePosition = position + 2;
                if (ePosition < line.size() && line[position] == 'o' && line[nPosition] == 'n' && line[ePosition] == 'e')
                {
                    return 1;
                }
            }
            else
            {
                // In backwards direction, position points to "e", and we want to find "on"
                auto nPosition = position - 1;
                auto oPosition = position - 2;
                if (oPosition >= 0 && line[position] == 'e' && line[nPosition] == 'n' && line[oPosition] == 'o')
                {
                    return 1;
                }
            }
            // Return -1 if not found
            return -1;
        }
        // Helper method to find "two" efficiently
        int findTwoInLine(const std::string& line, int position, bool forwardDirection)
        {
            if (forwardDirection)
            {
                // In forward direction, position points to "t", and we want to find "wo"
                auto wPosition = position + 1;
                auto oPosition = position + 2;
                if (oPosition < line.size() && line[position] == 't' && line[wPosition] == 'w' && line[oPosition] == 'o')
                {
                    return 2;
                }
            }
            else
            {
                // In backwards direction, position points to "o", and we want to find "tw"
                auto wPosition = position - 1;
                auto tPosition = position - 2;
                if (tPosition >= 0 && line[position] == 'o' && line[wPosition] == 'w' && line[tPosition] == 't')
                {
                    return 2;
                }
            }
            // Return -1 if not found
            return -1;
        }
        // Helper method to find "three" efficiently
        int findThreeInLine(const std::string& line, int position, bool forwardDirection)
        {
            if (forwardDirection)
            {
                // In forward direction, position points to "t", and we want to find "hree"
                auto hPosition = position + 1;
                auto rPosition = position + 2;
                auto firstEPosition = position + 3;
                auto secondEPosition = position + 4;
                if (secondEPosition < line.size() && line[position] == 't' && line[hPosition] == 'h' && line[rPosition] == 'r' && line[firstEPosition] == 'e' && line[secondEPosition] == 'e')
                {
                    return 3;
                }
            }
            else
            {
                // In backwards direction, position points to "e", and we want to find "thre"
                auto ePosition = position - 1;
                auto rPosition = position - 2;
                auto hPosition = position - 3;
                auto tPosition = position - 4;
                if (tPosition >= 0 && line[position] == 'e' && line[ePosition] == 'e' && line[rPosition] == 'r' && line[hPosition] == 'h' && line[tPosition] == 't')
                {
                    return 3;
                }
            }
            // Return -1 if not found
            return -1;
        }
        // Helper method to find "four" efficiently
        int findFourInLine(const std::string& line, int position, bool forwardDirection)
        {
            if (forwardDirection)
            {
                // In forward direction, position points to "f", and we want to find "our"
                auto oPosition = position + 1;
                auto uPosition = position + 2;
                auto rPosition = position + 3;
                if (rPosition < line.size() && line[position] == 'f' && line[oPosition] == 'o' && line[uPosition] == 'u' && line[rPosition] == 'r')
                {
                    return 4;
                }
            }
            else
            {
                // In backwards direction, position points to "r", and we want to find "fou"
                auto uPosition = position - 1;
                auto oPosition = position - 2;
                auto fPosition = position - 3;
                if (fPosition >= 0 && line[position] == 'r' && line[uPosition] == 'u' && line[oPosition] == 'o' && line[fPosition] == 'f')
                {
                    return 4;
                }
            }
            // Return -1 if not found
            return -1;
        }
        // Helper method to find "five" efficiently
        int findFiveInLine(const std::string& line, int position, bool forwardDirection)
        {
            if (forwardDirection)
            {
                // In forward direction, position points to "f", and we want to find "ive"
                auto iPosition = position + 1;
                auto vPosition = position + 2;
                auto ePosition = position + 3;
                if (ePosition < line.size() && line[position] == 'f' && line[iPosition] == 'i' && line[vPosition] == 'v' && line[ePosition] == 'e')
                {
                    return 5;
                }
            }
            else
            {
                // In backwards direction, position points to "e", and we want to find "fiv"
                auto vPosition = position - 1;
                auto iPosition = position - 2;
                auto fPosition = position - 3;
                if (fPosition >= 0 && line[position] == 'e' && line[vPosition] == 'v' && line[iPosition] == 'i' && line[fPosition] == 'f')
                {
                    return 5;
                }
            }
            // Return -1 if not found
            return -1;
        }
        // Helper method to find "six" efficiently
        int findSixInLine(const std::string& line, int position, bool forwardDirection)
        {
            if (forwardDirection)
            {
                // In forward direction, position points to "s", and we want to find "ix"
                auto iPosition = position + 1;
                auto xPosition = position + 2;
                if (xPosition < line.size() && line[position] == 's' && line[iPosition] == 'i' && line[xPosition] == 'x')
                {
                    return 6;
                }
            }
            else
            {
                // In backwards direction, position points to "x", and we want to find "si"
                auto iPosition = position - 1;
                auto sPosition = position - 2;
                if (sPosition >= 0 && line[position] == 'x' && line[iPosition] == 'i' && line[sPosition] == 's')
                {
                    return 6;
                }
            }
            // Return -1 if not found
            return -1;
        }
        // Helper method to find "seven" efficiently
        int findSevenInLine(const std::string& line, int position, bool forwardDirection)
        {
            if (forwardDirection)
            {
                // In forward direction, position points to "s", and we want to find "even"
                auto firstEPosition = position + 1;
                auto vPosition = position + 2;
                auto secondEPosition = position + 3;
                auto nPosition = position + 4;
                if (nPosition < line.size() && line[position] == 's' && line[firstEPosition] == 'e' && line[vPosition] == 'v' && line[secondEPosition] == 'e' && line[nPosition] == 'n')
                {
                    return 7;
                }
            }
            else
            {
                // In backwards direction, position points to "n", and we want to find "seve"
                auto firstEPosition = position - 1;
                auto vPosition = position - 2;
                auto secondEPosition = position - 3;
                auto sPosition = position - 4;
                if (sPosition >= 0 && line[position] == 'n' && line[firstEPosition] == 'e' && line[vPosition] == 'v' && line[secondEPosition] == 'e' && line[sPosition] == 's')
                {
                    return 7;
                }
            }
            // Return -1 if not found
            return -1;
        }
        // Helper method to find "eight" efficiently
        int findEightInLine(const std::string& line, int position, bool forwardDirection)
        {
            if (forwardDirection)
            {
                // In forward direction, position points to "e", and we want to find "ight"
                auto iPosition = position + 1;
                auto gPosition = position + 2;
                auto hPosition = position + 3;
                auto tPosition = position + 4;
                if (iPosition < line.size() && line[position] == 'e' && line[iPosition] == 'i' && line[gPosition] == 'g' && line[hPosition] == 'h' && line[tPosition] == 't')
                {
                    return 8;
                }
            }
            else
            {
                // In backwards direction, position points to "t", and we want to find "eigh"
                auto hPosition = position - 1;
                auto gPosition = position - 2;
                auto iPosition = position - 3;
                auto ePosition = position - 4;
                if (ePosition >= 0 && line[position] == 't' && line[hPosition] == 'h' && line[gPosition] == 'g' && line[iPosition] == 'i' && line[ePosition] == 'e')
                {
                    return 8;
                }
            }
            // Return -1 if not found
            return -1;
        }
        // Helper method to find "nine" efficiently
        int findNineInLine(const std::string& line, int position, bool forwardDirection)
        {
            if (forwardDirection)
            {
                // In forward direction, position points to "n", and we want to find "ine"
                auto iPosition = position + 1;
                auto nPosition = position + 2;
                auto ePosition = position + 3;
                if (ePosition < line.size() && line[position] == 'n' && line[iPosition] == 'i' && line[nPosition] == 'n' && line[ePosition] == 'e')
                {
                    return 9;
                }
            }
            else
            {
                // In backwards direction, position points to "e", and we want to find "nin"
                auto firstNPosition = position - 1;
                auto iPosition = position - 2;
                auto secondNPosition = position - 3;
                if (secondNPosition >= 0 && line[position] == 'e' && line[firstNPosition] == 'n' && line[iPosition] == 'i' && line[secondNPosition] == 'n')
                {
                    return 9;
                }
            }
            // Return -1 if not found
            return -1;
        }
        // Compute the pair<int,int> as the result of the elf's line
        std::pair<int, int> computeFirstLastDigitForLine(const std::string& line);
        // Try to find the numeric value in the line, looking at m_lookUpCharToIntMap
        int findNumericValue(const std::string& line, int position, bool forwardDirection);

        /* Private member variables */
        // Map for fast look-up
        const std::unordered_map<char, int> m_lookUpCharToIntMap{
            {'0', 0},
            {'1', 1},
            {'2', 2},
            {'3', 3},
            {'4', 4},
            {'5', 5},
            {'6', 6},
            {'7', 7},
            {'8', 8},
            {'9', 9},
            // For part two I must consider letters 'o' for 'one', 't' for 'two' and 'three', 'f' for 'four' and 'five',
            // 's' for 'six' and 'seven', 'e' for 'eight', and finally 'n' for 'nine' - beginPtr
            //
            // I must also consider the letters 'e' for 'one', 'three', 'five', and 'nine'
            // 'o' for 'two', 'r' for 'four', 'x' for 'six', 't' for 'eight', 'n' for 'seven'
            {'o', LITERAL_ONE_OR_TWO},
            {'t', LITERAL_TWO_OR_THREE_OR_EIGHT},
            {'f', LITERAL_FOUR_OR_FIVE},
            {'s', LITERAL_SIX_OR_SEVEN},
            {'e', LITERAL_ONE_OR_THREE_OR_FIVE_OR_EIGHT_OR_NINE},
            {'n', LITERAL_SEVEN_OR_NINE},
            {'r', LITERAL_FOUR},
            {'x', LITERAL_SIX}
        };
        // File name with the inputs
        std::string m_fileName;
};