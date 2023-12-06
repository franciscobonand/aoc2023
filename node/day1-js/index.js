import fs from 'fs';

function extractCalibrationValues(input) {
    const lines = input.split('\n');
    let sum = 0;

    const numMap = {
        one: '1',
        two: '2',
        three: '3',
        four: '4',
        five: '5',
        six: '6',
        seven: '7',
        eight: '8',
        nine: '9',
        zero: '0'
    };

    const leftRegex = new RegExp(`(${Object.keys(numMap).join("|")})$`);
  const rightRegex = new RegExp(`^(${Object.keys(numMap).join("|")})`);
  const digitRegex = new RegExp(`^([0-9])`);

  lines.forEach((line) => {
    let strLeft = "";
    let strRight = "";

    let leftDigit = "";
    let rightDigit = "";

    let i = 0;
    while (!leftDigit || !rightDigit) {
      if (!leftDigit) {
        strLeft += line[i];
        if (line[i].match(digitRegex)) {
          leftDigit = line[i];
        } else if (strLeft.match(leftRegex)) {
          leftDigit = numMap[strLeft.match(leftRegex)[0]];
        }
      }

      if (!rightDigit) {
        strRight = line[line.length - i - 1] + strRight;

        if (line[line.length - i - 1].match(digitRegex)) {
          rightDigit = line[line.length - i - 1];
        } else if (strRight.match(rightRegex)) {
          rightDigit = numMap[strRight.match(rightRegex)[0]];
        }
      }

      i += 1;
    }

    const numStr = leftDigit + rightDigit;
    sum += Number.parseInt(numStr);
  });

  return sum;
}

function readFile(filePath) {
    try {
        const fileContent = fs.readFileSync(filePath, 'utf8');
        return fileContent;
    } catch (error) {
        console.error('Error reading the file:', error.message);
        process.exit(1);
    }
}

const filePath = './input.txt';
const fileContent = readFile(filePath);
const totalCalibration = extractCalibrationValues(fileContent.trim());
console.log('Total Calibration:', totalCalibration);
