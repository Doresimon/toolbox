const xlsx = require('xlsx');
const fs = require('fs');

JsonToSheet();

function JsonToSheet() {
    const inFile = "./tmp/input.json";
    const outFile = "./tmp/output.xlsx";
    const sheetName = "sheet0"

    const jsonString = fs.readFileSync(inFile);
    const jsonArray = JSON.parse(jsonString);

    // flatten data to depth=1

    // options
    const options = {
        header: ["_id"],
        // skipHeader: ["h2"],
    }

    // transform to work sheet
    const worksheet = xlsx.utils.json_to_sheet(jsonArray, options);

    // make work book
    const workbook = xlsx.utils.book_new();
    workbook.SheetNames.push(sheetName);
    workbook.Sheets[sheetName] = worksheet;

    // save to file
    xlsx.writeFile(workbook, outFile);

    console.log("done:", outFile);
}
