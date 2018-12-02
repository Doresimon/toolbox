const PDFDocument = require('pdfkit')
const fs = require('fs')
 
// Create a document 
// doc = new PDFDocument

// doc.pipe(fs.createWriteStream('output.pdf'))
// doc.image('./img-folder/001/73004-02-01-001.jpg', 0, 0, {width: 612 })
// doc.addPage()
// doc.image('./img-folder/001/85cae9868086f431ef3080d0f98a6f9.png', 0, 0, {width: 612 })
// doc.addPage()
// doc.end()


let sql = {
    str:"",
    categoryID:103,
    AddCat(){
        this.str += `INSERT INTO "main"."category"("id", "pre_id", "category_name") VALUES (${id}, ${pre_id}, ${category_name});\n`
    },
}
// INSERT INTO "main"."category"("id", "pre_id", "category_name") VALUES (103, 6, '啊啊啊');
// let categoryID = 103;

// function 
 
function readBigBox(path){
    let boxs = fs.readdirSync(path)

    console.log(boxs)

    return

    // get No

    let box = {
        id: catID,
        name: `盒-${No}`,
        pre_id:0,
    }

    return list
}

function readSmallBox(path){

    return list
}


readBigBox("./box")