#include "CalibrationInterpreter.hpp"

int CalibrationInterpreter::findNumericValue(const std::string& line, int position, bool forwardDirection)
{
    if (auto intItr = m_lookUpCharToIntMap.find(line[position]); intItr != m_lookUpCharToIntMap.end())
    {
        // Code for first part
        if (intItr->second < NUMERIC_VALUE_CEILING)
            return intItr->second;
        // Code for second part
       else if (intItr->second == LITERAL_ONE_OR_TWO)
       {
           // Try to find one in forward direction or two in backwards direction
           return forwardDirection ? findOneInLine(line, position, true) : findTwoInLine(line, position, false);
       }
       else if (intItr->second == LITERAL_TWO_OR_THREE_OR_EIGHT)
       {
            if (forwardDirection)
            {
                // Try to find two or three forward direction
                auto temp = findTwoInLine(line, position, true);
                return temp != -1 ? temp : findThreeInLine(line, position, true);
            }
            // Try to find eight in backwards direction
            return findEightInLine(line, position, false);
       }
       else if (forwardDirection && intItr->second == LITERAL_FOUR_OR_FIVE)
       {
            // Try to find four or five in forward direction
            auto temp = findFourInLine(line, position, true);
            return temp != -1 ? temp : findFiveInLine(line, position, true);
       }
       else if (forwardDirection && intItr->second == LITERAL_SIX_OR_SEVEN)
       {
            // Try to find six or seven in forward direction
            auto temp = findSixInLine(line, position, true);
            return temp != -1 ? temp : findSevenInLine(line, position, true);
       }
       else if (intItr->second == LITERAL_ONE_OR_THREE_OR_FIVE_OR_EIGHT_OR_NINE)
       {
            if (!forwardDirection)
            {
                // Try to find one in backwards direction
                auto temp = findOneInLine(line, position, false);
                // Try to find three in backwards direction
                temp = temp != -1 ? temp : findThreeInLine(line, position, false);
                // Try to find five in backwards direction
                temp = temp != -1 ? temp : findFiveInLine(line, position, false);
                // Try to find nine in backwards direction
                return temp != -1 ? temp : findNineInLine(line, position, false);
            }
            // Try to find eight in forwards direction
            return findEightInLine(line, position, true);
       }
       else if (intItr->second == LITERAL_SEVEN_OR_NINE)
       {
            // Try to find nine in forward direction or seven in backwards direction
            return forwardDirection ? findNineInLine(line, position, true) : findSevenInLine(line, position, false);
       }
       else if (!forwardDirection && intItr->second == LITERAL_FOUR)
       {
            // Try to find four in backwards direction
            return findFourInLine(line, position, false);
       }
       else if (!forwardDirection && intItr->second == LITERAL_SIX)
       {
            return findSixInLine(line, position, false);
       }
    }
    // In case not found, return -1
    return -1;
}

std::pair<int, int> CalibrationInterpreter::computeFirstLastDigitForLine(const std::string& line)
{
    // Pair to store the digit
    std::pair<int, int> firstAndLastDigit{-1, -1};
    // Pointers to the beginning and end of the line    
    unsigned int beginPtr = 0;
    unsigned int endPtr = line.size() > 0 ? line.size() - 1 : 0;

    while (beginPtr <= endPtr)
    {
        // Base case, both digits were found
        if (firstAndLastDigit.first != -1 && firstAndLastDigit.second != -1)
            break;
        // Finding the first digit
        // If value was found, avoid one look-up
        firstAndLastDigit.first = firstAndLastDigit.first == -1 ? findNumericValue(line, beginPtr, true) : firstAndLastDigit.first;
        // Value not found, increment the begin pointer
        if (firstAndLastDigit.first == -1)
            ++beginPtr;
        // Finding the second digit
        firstAndLastDigit.second = firstAndLastDigit.second == -1 ? findNumericValue(line, endPtr, false) : firstAndLastDigit.second;
        // Value not found, increment the end pointer
        if (firstAndLastDigit.second == -1)
            --endPtr;
    }
    // Maybe change this assertion to another early return?
    assert(firstAndLastDigit.first != -1 || firstAndLastDigit.second != -1);
    // Case in which there is only one digit
    if (firstAndLastDigit.first == -1)
        firstAndLastDigit.first = firstAndLastDigit.second;
    if (firstAndLastDigit.second == -1)
        firstAndLastDigit.second = firstAndLastDigit.first;
    // Return the pair
    return firstAndLastDigit;
}

int CalibrationInterpreter::computeSumOfCalibrationValues()
{
    // Maybe do memory mapping? Since we are only reading this could save time
    std::ifstream fileContentStream;
    fileContentStream.open(m_fileName);
    // Maybe change this from assertion to early return?
    assert(fileContentStream.is_open());
    // Read line by line of the file
    std::string line;
    // Accumulate value
    int accum = 0;
    while (getline(fileContentStream, line))
    {
        std::cout << "line : " << line << std::endl;
        auto firstAndLastDigit = computeFirstLastDigitForLine(line);
        std::cout << "first : " << firstAndLastDigit.first << std::endl;
        std::cout << "second : " << firstAndLastDigit.second << std::endl;
        accum += (10 * firstAndLastDigit.first) + firstAndLastDigit.second;
    }
    std::cout << "final value is : " << accum << std::endl;
    return 0;
}