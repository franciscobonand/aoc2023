import 'dart:io';

File _readFile(String path) {
  return new File(path);
}

void part1(List<String> lines) {
  RegExp exp = RegExp(r'\d');

  int totalSum = 0;

  for (var line in lines) {
    final match = exp.allMatches(line);
    final strAsNum = match.first[0]! + match.last[0]!;
    totalSum += int.parse(strAsNum);
  }

  print(totalSum);
}

void part2(List<String> lines) {
  final dict = {
    "one": 1,
    "eno": 1,
    "two": 2,
    "owt": 2,
    "three": 3,
    "eerht": 3,
    "four": 4,
    "ruof": 4,
    "five": 5,
    "evif": 5,
    "six": 6,
    "xis": 6,
    "seven": 7,
    "neves": 7,
    "eight": 8,
    "thgie": 8,
    "nine": 9,
    "enin": 9,
    "ten": 10,
    "net": 10
  };

  int totalSum = 0;

  RegExp exp = RegExp(
      r'one|two|three|four|five|six|seven|eight|nine|ten|1|2|3|4|5|6|7|8|9');
  RegExp exp2 = RegExp(
      r'9|8|7|6|5|4|3|2|1|net|enin|thgie|neves|xis|evif|ruof|eerht|owt|eno');

  for (String line in lines) {
    final match = exp.firstMatch(line);
    final match2 = exp2.firstMatch(line.split('').reversed.join());
    String _tempFirstNum = match![0]!;
    String _tempSecondNum = match2![0]!;

    if (_tempFirstNum.length > 1) {
      _tempFirstNum = dict[_tempFirstNum]!.toString();
    }
    if (_tempSecondNum.length > 1) {
      _tempSecondNum = dict[_tempSecondNum]!.toString();
    }
    totalSum += int.parse(_tempFirstNum + _tempSecondNum);
  }

  print(totalSum);
}

void main(List<String> args) {
  final _input = _readFile('input');
  List<String> lines = _input.readAsLinesSync();
  part1(lines);
  part2(lines);
}
