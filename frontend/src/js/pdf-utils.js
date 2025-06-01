export function exportTableToPdf(cols, logs) {
    if (typeof jspdf === 'undefined' || !jspdf.jsPDF) {
        console.error("jsPDF Not loaded or not correctly initialized. Cannot export to PDF.");
        return;
    }

    const { jsPDF } = jspdf;
    const doc = new jsPDF('landscape');

    const data = logs.map(log => {
        return cols.map(col => (log[col.Name] !== undefined ? String(log[col.Name]) : ''));
    });

    const headers = cols.map(col => col.DisplayName);

    const fontSize = 10;
    const lineHeight = doc.internal.getFontSize() * doc.internal.getLineHeightFactor();
    doc.setFontSize(fontSize);

    const marginX = 10;
    let currentY = 10;
    const pageHeight = doc.internal.pageSize.height;
    const pageWidth = doc.internal.pageSize.width;

    let maxColWidth = pageWidth / cols.length;

    const columnWidths = cols.map(() => maxColWidth);

    const addPageAndHeader = () => {
        doc.addPage();
        currentY = marginX;
        doc.setFont(undefined, 'bold');
        headers.forEach((header, index) => {
            doc.text(header, marginX + (index * columnWidths[index]), currentY);
        });
        doc.setFont(undefined, 'normal');
        currentY += lineHeight + 5;
        doc.line(marginX, currentY, pageWidth - marginX, currentY);
        currentY += 5;
    };

    doc.text("Log Report", marginX, currentY);
    currentY += lineHeight * 2;

    doc.setFont(undefined, 'bold');
    headers.forEach((header, index) => {
        doc.text(header, marginX + (index * columnWidths[index]), currentY);
    });
    doc.setFont(undefined, 'normal');

    currentY += lineHeight + 5;
    doc.line(marginX, currentY, pageWidth - marginX, currentY);
    currentY += 5;

    data.forEach(rowData => {
        let currentRowHeight = lineHeight;
        let splittedCells = {};

        rowData.forEach((cellValue, colIndex) => {

            if (doc.getTextWidth(cellValue) > columnWidths[colIndex] - 2) {
                const columnWidth = columnWidths[colIndex];

                const splitContent = doc.splitTextToSize(cellValue, columnWidth - 2);
                splittedCells[colIndex] = splitContent;

                const contentHeight = splitContent.length * lineHeight;

                currentRowHeight = Math.max(currentRowHeight, contentHeight);
            }
        });

        if (currentY + currentRowHeight > pageHeight) {
            addPageAndHeader();
        }

        rowData.forEach((cellValue, colIndex) => {
            const columnX = marginX + (colIndex * columnWidths[colIndex]);

            if (splittedCells[colIndex]) {
                doc.text(splittedCells[colIndex], columnX, currentY);
            } else {
                doc.text(cellValue, columnX, currentY);
            }
        });

        currentY += currentRowHeight;
        doc.line(marginX, currentY, pageWidth - marginX, currentY);
        currentY += 5;
    });

    doc.save('log_report.pdf');
}
