export async function exportToPDF(datas) {
    const { jsPDF } = window.jspdf;
    const doc = new jsPDF();
    const startY = 20;
    let y = startY;

    doc.setFontSize(10);
    const lineHeight = 8;

    doc.setFontSize(14);
    doc.text("Exported Log", 14, 10);
    doc.setFontSize(10);

    doc.setFont(undefined, 'bold');

    const pdfColumnWidthsMap = {
        "Source": 20,
        "IPAddress": 25,
        "Date": 20,
        "Time": 25,
        "PID": 15,
        "TID": 15,
        "Level": 15,
        "Tag": 30,
        "Message": 30,
        "StatusCode": 20,
        "EventID": 25,
        "UserID": 25,
        "Action": 40,
        "Timestamp": 35,
        "Component": 25,
        "Severity": 20,
        "Description": 80
    };

    if (!datas || datas.length === 0 || !datas[0].Cols) {
        alert("No data filtered to export");
        return;
    }
    const colsToExport = datas[0].Cols;

    const pdfColumnLayout = [];
    let startX = 14;

    colsToExport.forEach(colName => {
        const width = pdfColumnWidthsMap[colName.Name] || 25;
        pdfColumnLayout.push({
            name: colName.Name,
            x: startX,
            width: width
        });
        startX += width + 2;
    });

    y = printHeaders(doc, pdfColumnLayout, y, lineHeight);

    const maxY = 280;

    datas.forEach((r) => {
        var rowData = r.ParsedData

        let currentMaxRowHeight = lineHeight;
        const splitTextColumns = {};

        pdfColumnLayout.forEach(col => {
            const cellValue = String(rowData[col.name] || "");
            if (["Message", "Description", "Action"].includes(col.name)) {
                const splitContent = doc.splitTextToSize(cellValue, col.width);
                const contentHeight = splitContent.length * lineHeight;
                currentMaxRowHeight = Math.max(currentMaxRowHeight, contentHeight);
                splitTextColumns[col.name] = splitContent;
            }
        });

        if (y + currentMaxRowHeight > maxY) {
            doc.addPage();
            y = startY;
            y = printHeaders(doc, pdfColumnLayout, y, lineHeight);
        }

        pdfColumnLayout.forEach(col => {
            if (splitTextColumns[col.name]) {
                doc.text(splitTextColumns[col.name], col.x, y);
            } else {
                const cellValue = String(rowData[col.name] || "");
                doc.text(cellValue, col.x, y);
            }
        });

        y += currentMaxRowHeight;
    });

    await doc.save("filtered_logs.pdf");
}

function printHeaders(doc, layout, currentY, lineHeight) {
    doc.setFont(undefined, 'bold');
    layout.forEach(col => {
        console.log(col)
        doc.text(col.name, col.x, currentY);
    });
    doc.setFont(undefined, 'normal');
    return currentY + lineHeight + 2;
}
