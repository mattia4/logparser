import { exportToPDF } from './pdf-utils.js';

const rowsPerPage = 100;
let currentPage = 1;
let filteredRows = [];
let allRows = [];
let searchInput;
let output;
let logSection;
let dropFileSection;
let btnExportPdf;
let loadingSpinner;
const activeColumnFilters = {};

document.addEventListener("DOMContentLoaded", () => {
    const fileInput = document.getElementById('jsonFileInput');
    loadingSpinner = document.getElementById('loadingSpinner');
    output = document.getElementById('output');
    searchInput = document.getElementById('searchInput');
    btnExportPdf = document.getElementById('btn-export-pdf');

    logSection = document.getElementById("log-section")
    dropFileSection = document.getElementById("drop-file-section")

    logSection.style.display = 'none';
    btnExportPdf.style.display = 'none';
    dropFileSection.style.display = 'block';

    fileInput.addEventListener('change', (event) => {
        handleFiles(event.target.files);
    });

    dropZone.addEventListener('dragover', (event) => {
        event.preventDefault();
        dropZone.classList.add('dragover');
    });

    dropZone.addEventListener('dragleave', () => {
        dropZone.classList.remove('dragover');
    });

    dropZone.addEventListener('drop', (event) => {
        event.preventDefault();
        dropZone.classList.remove('dragover');
        handleFiles(event.dataTransfer.files);
    });



    const generatePdfButton = document.getElementById('generate-pdf-button');

    if (generatePdfButton) {
        generatePdfButton.addEventListener('click', () => {
            const dataToInclude = {
                activeFilters: activeColumnFilters,
            };
            exportToPDF(dataToInclude, 'Log_Report_2025');
        });
    }
});

function handleFiles(files) {
    showSpinner()

    if (files.length === 0) {
        fileMessage.textContent = "No file selected";
        fileMessage.className = "message";
        return;
    }

    const selectedFile = files[0];

    const reader = new FileReader();
    reader.onload = (e) => {
        try {
            logSection.style.display = 'block';
            btnExportPdf.style.display = 'block';
            dropFileSection.style.display = 'none';
            loadingSpinner.style.display = "none";

            const logsData = JSON.parse(e.target.result);

            filteredRows = [...logsData.LogResult];
            allRows = filteredRows;
            allColumns = logsData.Cols;

            searchInput.addEventListener("input", filterArrayBySearchTerm);
            generateTableColumns(allColumns)

            renderTable();
            hideSpinner()

        } catch (err) {
            output.textContent = "Parse error: " + err.message;
        }
    };

    reader.onerror = (e) => {
        fileMessage.textContent = `Error reading: ${e.target.error}`;
        fileMessage.className = "message error";
        console.error("Errore FileReader:", e.target.error);
    };

    reader.readAsText(selectedFile);
}

function changePage(direction) {
    const totalPages = Math.ceil(allRows.length / rowsPerPage);
    currentPage = Math.min(Math.max(currentPage + direction, 1), totalPages);
    renderTable();
}

function generateTableColumns(allColumns) {
    var colsSpawn = document.getElementById("trSpawnColums");

    for (const key in allColumns) {
        const div = document.createElement('div')
        div.className = "div-filter-col"
        const thElement = document.createElement('th');

        const input = document.createElement('input');
        input.type = 'text';
        input.name = 'filter-' + allColumns[key].Name.toLowerCase();
        input.placeholder = allColumns[key].Name.toLowerCase();
        input.id = 'filter-' + allColumns[key].Name.toLowerCase();

        input.addEventListener('input', function (event) {
            const inputElement = event.target;
            const columnName = inputElement.value;

            updateColumnFilter(allColumns[key].Name, columnName)
        });

        thElement.textContent = allColumns[key].Value;
        thElement.className = allColumns[key].Value.toLowerCase();

        colsSpawn.appendChild(thElement);
        thElement.appendChild(div);
        div.appendChild(input);

    }
}
function updateColumnFilter(columnName, filterValue) {
    if (filterValue) {
        activeColumnFilters[columnName] = filterValue;
    } else {
        delete activeColumnFilters[columnName];
    }

}

function renderTable() {
    const tableBody = document.getElementById('logTableRows');
    tableBody.innerHTML = '';

    const startIdx = (currentPage - 1) * rowsPerPage;
    const endIdx = Math.min(currentPage * rowsPerPage, filteredRows.length);
    const visibleRows = filteredRows.slice(startIdx, endIdx);

    const searchInput = document.getElementById("searchInput");
    const searchedText = searchInput.value.toLowerCase();

    visibleRows.forEach(r => {
        const rowElement = document.createElement('tr');
        var row = r.ParsedData
        if (row != null && r.Cols != null) {
            r.Cols.forEach(col => {
                var value = row[col.Name]
                if (searchedText != "") {
                    rowElement.innerHTML += `<td class="${value}">${highlightText(value, searchedText)}</td>`;
                } else {
                    rowElement.innerHTML += `<td class="${value}">${value}</td>`;
                }
                tableBody.appendChild(rowElement);
            })
        }
    });
    renderPagination()
}

function renderPagination() {
    const pagination = document.getElementById('pagination-numbers');
    pagination.innerHTML = '';
    const totalPages = Math.ceil(filteredRows.length / rowsPerPage);

    for (let i = 1; i <= totalPages; i++) {
        const btn = document.createElement('button');
        btn.className = 'page-btn';
        btn.innerText = i;
        btn.onclick = () => {
            currentPage = i;
            renderTable();
        };
        if (i === currentPage) {
            btn.style.fontWeight = 'bold';
            btn.style.backgroundColor = "#e0f2fe"
        }
        pagination.appendChild(btn);
    }
}

function filterCols() {
    if (Object.keys(activeColumnFilters).length > 0) {
        visibleRows.filter(logEntry => {
            for (const colName in activeColumnFilters) {
                const filterValue = activeColumnFilters[colName].toLowerCase();
                const logEntryValue = String(logEntry[colName] || '').toLowerCase();
                if (logEntryValue == filterValue) {
                    return false;
                }
            }
            return true;
        });
        renderPagination()
    }

    currentPage = 1;
    renderTable();
}

function filterArrayBySearchTerm() {

    const input = document.getElementById("searchInput");
    const filter = input.value.toLowerCase();

    if (filter == "") {
        filteredRows = allRows
        currentPage = 1;
        renderTable();
        return;
    }

    filteredRows = allRows.filter(item => {
        for (const key in item) {
            if (Object.prototype.hasOwnProperty.call(item, key)) {
                const value = item.RawLine;
                if (typeof value === 'string') {
                    if (value.toLowerCase().includes(filter)) {
                        return true;
                    }
                }
            }
        }
    });

    currentPage = 1;
    renderTable();
}

function highlightText(text, filter) {
    if (!filter || filter.trim() === '') {
        return text;
    }

    const escapedFilter = filter.replace('/[.*+?^${}()|[\]\\]/g', '\\$&');
    const regex = new RegExp(escapedFilter, 'gi');
    return text.replace(regex, '<span class="highlight">$&</span>');
}

/*
async function exportToPDF() {
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

    if (!filteredRows || filteredRows.length === 0 || !filteredRows[0].Cols) {
        alert("No data filtered to export");
        return;
    }
    const colsToExport = filteredRows[0].Cols;

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

    filteredRows.forEach((r) => {
        var rowData = r.ParsedData

        let currentMaxRowHeight = lineHeight;
        const splitTextColumns = {};

        pdfColumnLayout.forEach(col => {
            const cellValue = String(rowData[col.name] || "");
            // Controlla se questa colonna è una di quelle che possono avere testo lungo
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
*/
function showSpinner() { loadingSpinner.classList.add('active'); }

function hideSpinner() { loadingSpinner.classList.remove('active'); }

function printHeaders(doc, layout, currentY, lineHeight) {
    doc.setFont(undefined, 'bold');
    layout.forEach(col => {
        console.log(col)
        doc.text(col.name, col.x, currentY);
    });
    doc.setFont(undefined, 'normal');
    return currentY + lineHeight + 2;
}


