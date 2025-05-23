
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

    visibleRows.forEach(row => {
        const rowElement = document.createElement('tr');

        rowElement.innerHTML = `
                <td class="date">${row.Date}</td>
                <td class="time">${row.Time}</td>
                <td class="level">${row.Level}</td>
                <td class="tag">${row.Tag}</td>
                <td class="pid">${row.PID}</td>
                <td class="message">${row.Message}</td>
            `;
        tableBody.appendChild(rowElement);
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
        return
        row.Date.toLowerCase().includes(filter) ||
            row.Level.toLowerCase().includes(filter) ||
            row.Message.toLowerCase().includes(filter) ||
            row.PID.toLowerCase().includes(filter) ||
            row.Tag.toString().includes(filter) ||
            row.Time.toLowerCase().includes(filter)
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
    doc.text("Date", 14, y);
    doc.text("Time", 34, y);
    doc.text("Process", 64, y);
    doc.text("PID", 94, y);
    doc.text("Log Level", 114, y);
    doc.text("Message", 134, y);
    doc.setFont(undefined, 'normal');
    y += lineHeight + 2;

    const maxY = 280;

    const colPositions = {
        date: { x: 14, width: 20 },
        time: { x: 34, width: 30 },
        tag: { x: 64, width: 28 },
        pid: { x: 94, width: 18 },
        level: { x: 114, width: 18 },
        message: { x: 134, width: 65 }
    };

    filteredRows.forEach((row) => {
        let currentMaxRowHeight = lineHeight;
        const date = String(row.Date || "");
        const time = String(row.Time || "");
        const tag = String(row.Tag || "");
        const pid = String(row.PID || "");
        const level = String(row.Level || "");
        const message = String(row.Message || "");

        const splitMessage = doc.splitTextToSize(message, colPositions.message.width);
        const splitTag = doc.splitTextToSize(tag, colPositions.tag.width);
        const messageHeight = splitMessage.length * lineHeight;
        const processHeight = splitTag.length * lineHeight;
        currentMaxRowHeight = Math.max(currentMaxRowHeight, messageHeight, processHeight);

        if (y + currentMaxRowHeight > maxY) {
            doc.addPage();
            y = startY;

            doc.setFont(undefined, 'bold');
            doc.text("Date", colPositions.date.x, y);
            doc.text("Time", colPositions.time.x, y);
            doc.text("Tag", colPositions.tag.x, y);
            doc.text("PID", colPositions.pid.x, y);
            doc.text("Log level", colPositions.level.x, y);
            doc.text("Message", colPositions.message.x, y);
            doc.setFont(undefined, 'normal');
            y += lineHeight + 2;
        }

        doc.text(date, colPositions.date.x, y);
        doc.text(time, colPositions.time.x, y);
        doc.text(splitTag, colPositions.tag.x, y);
        doc.text(pid, colPositions.pid.x, y);
        doc.text(level, colPositions.level.x, y);
        doc.text(splitMessage, colPositions.message.x, y);

        y += currentMaxRowHeight;
    });

    await doc.save("log_filtrati.pdf");
}