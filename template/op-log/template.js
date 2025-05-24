
const rowsPerPage = 100;
let currentPage = 1;
let filteredRows = [];
let allRows = [];

function init(logsData) {
    allRows = logsData
    filteredRows = [...logsData];
    renderTable();
}

function changePage(direction) {
    const totalPages = Math.ceil(allRows.length / rowsPerPage);
    currentPage = Math.min(Math.max(currentPage + direction, 1), totalPages);
    renderTable();
}

function renderTable() {
    const tableBody = document.getElementById('logTableRows');
    tableBody.innerHTML = '';

    const startIdx = (currentPage - 1) * rowsPerPage;
    const endIdx = Math.min(currentPage * rowsPerPage, filteredRows.length);
    const visibleRows = filteredRows.slice(startIdx, endIdx);

    visibleRows.forEach(r => {

        var row = r.ParsedData

        if (row != null) {
            const rowElement = document.createElement('tr');
            rowElement.innerHTML = `
                <td class="site">${row.Site}</td>
                <td class="iPAddress">${row.IPAddress}</td>
                <td class="date">${row.Date}</td>
                <td class="time">${row.Time}</td>
                <td class="message">${row.Message}</td>
                <td class="statusCode">${row.StatusCode}</td>
            `;
            tableBody.appendChild(rowElement);
        }

    });

    renderPagination()
}

function filterTable() {
    const input = document.getElementById("searchInput");
    const filter = input.value.toLowerCase();

    if (filter == "") {
        filteredRows = allRows
        currentPage = 1;
        renderTable();
        return;
    }

    filteredRows = allRows.filter(row => {
        return row.Site.toLowerCase().includes(filter) ||
            row.IPAddress.toLowerCase().includes(filter) ||
            row.Timestamp.toLowerCase().includes(filter) ||
            row.Message.toLowerCase().includes(filter) ||
            row.StatusCode.toString().includes(filter)
    });

    currentPage = 1;
    renderTable();

}

function renderPagination() {
    const pagination = document.getElementById('pagination_number');
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
        if (i === currentPage) btn.style.fontWeight = 'bold';
        pagination.appendChild(btn);
    }
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
                const value = item[key];

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

renderTable();


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
    doc.text("Site", 14, y);
    doc.text("IP address", 34, y);
    doc.text("Timestamp", 64, y);
    doc.text("Message", 94, y);
    doc.text("Status code", 114, y);
    doc.setFont(undefined, 'normal');
    y += lineHeight + 2;

    const maxY = 280;

    const colPositions = {
        site: { x: 14, width: 20 },
        iPAddress: { x: 34, width: 30 },
        timestamp: { x: 64, width: 28 },
        message: { x: 94, width: 18 },
        statusCode: { x: 114, width: 18 }
    };



    filteredRows.forEach((r) => {
        var row = r.ParsedData

        let currentMaxRowHeight = lineHeight;
        const site = String(row.Site || "");
        const iPAddress = String(row.IPAddress || "");
        const timestamp = String(row.Timestamp || "");
        const message = String(row.Message || "");
        const statusCode = String(row.StatusCode || "");

        const splitMessage = doc.splitTextToSize(message, colPositions.message.width);
        const messageHeight = splitMessage.length * lineHeight;
        currentMaxRowHeight = Math.max(currentMaxRowHeight, messageHeight);

        if (y + currentMaxRowHeight > maxY) {
            doc.addPage();
            y = startY;

            doc.setFont(undefined, 'bold');
            doc.text("Site", colPositions.site.x, y);
            doc.text("IP address", colPositions.iPAddress.x, y);
            doc.text("Timestamp", colPositions.timestamp.x, y);
            doc.text("Message", colPositions.message.x, y);
            doc.text("StatusCode", colPositions.statusCode.x, y);
            doc.setFont(undefined, 'normal');
            y += lineHeight + 2;
        }

        doc.text(site, colPositions.site.x, y);
        doc.text(iPAddress, colPositions.iPAddress.x, y);
        doc.text(timestamp, colPositions.timestamp.x, y);
        doc.text(message, colPositions.message.x, y);
        doc.text(statusCode, colPositions.statusCode.x, y);

        y += currentMaxRowHeight;
    });

    await doc.save("log_filtrati.pdf");
}


