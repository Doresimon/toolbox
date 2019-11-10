const XLSX = require('xlsx');


main()
function main() {
    const file = "./tmp/demo.xlsx"
    var workbook = XLSX.readFile(file);
    /* DO SOMETHING WITH workbook HERE */
    var first_sheet_name = workbook.SheetNames[0];
    var address_of_cell = 'A' + 4;
    // var address_of_cell = 'A1'

    /* Get worksheet */
    var worksheet = workbook.Sheets[first_sheet_name];

    /* Find desired cell */
    var desired_cell = worksheet[address_of_cell];

    /* Get the value */
    var desired_value = (desired_cell ? desired_cell.v : undefined);

    console.log("first_sheet_name", first_sheet_name)
    console.log("desired_value", desired_value)
}
