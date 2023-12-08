#include <regex>

// Functionality below should not depend on inline or static to make the linking stage work
// I should have created Helpers.cpp -> but due to laziness this could be a TO-DO or something.

namespace BasicMath
{
    // Careful, if not inline we will have a linking error
    inline int max(int a, int b)
    {
        return (a > b) ? a : b;
    }
};

namespace RegexHelper
{
    // Careful, if not static we will have a linking error
    static int returnMaxUsingPattern(const std::regex& pattern, const std::string& line)
    {
        // To accumulate the count (if it is game id, it should only add once which is fine)
        int accum = 0;
        // Doing the iterator way
        auto countBeginItr = std::sregex_iterator(line.begin(), line.end(), pattern);
        auto countEndItr = std::sregex_iterator();
        for (auto itr = countBeginItr; itr != countEndItr; ++itr)
        {
            auto match = (*itr).str();
            auto matchInt = std::atoi(match.c_str());
            accum = BasicMath::max(matchInt, accum);
        }
        return accum;
    }
};