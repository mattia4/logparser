import { highlightText, filterArrayBySearchTerm, filterByColumns, orderByCol, cleanString } from './table-functions.js';
import { hideHtmlElement, resetVisibilityHtmlElements, DivHelper } from './helpers.js';
import { exportTableToPdf } from './pdf-utils.js';
import { DialogComponent } from './dialog.js'

let allLogs = [];
let allCols = [];
let filteredLogs = [];
let currentPage = 1;
const logsPerPage = 50;

const loadingSpinner = document.getElementById('loadingSpinner');
const logTableRows = document.getElementById('logTableRows');
const trSpawnColums = document.getElementById('trSpawnColums');
const searchInput = document.getElementById('searchInput');
const paginationNumbers = document.getElementById('pagination-numbers');
const pagination = document.getElementById('pagination');
const prevPageBtn = document.getElementById('prevPageBtn');
const nextPageBtn = document.getElementById('nextPageBtn');
const btnExportPdf = document.getElementById('btn-export-pdf');

hideHtmlElement(paginationNumbers)
hideHtmlElement(pagination)

function renderColumns(cols) {
    trSpawnColums.innerHTML = '';
    cols.forEach((col, index) => {
        const th = document.createElement('th');

        let r = DivHelper.makeRow(null, {});
        let c1 = DivHelper.makeCol(col.DisplayName, { class: 'col-auto' });
        let c2 = DivHelper.makeCol("▼", { class: ['col-1', "order_" + index, 'order_col_style'] });

        c2.addEventListener('click', () => {
            let d = document.getElementsByClassName("order_" + index);
            if (d[0].classList.contains("order_" + index) && !d[0].classList.contains("rot180deg")) {
                d[0].classList.add('rot180deg');
                orderBy(col, true);
            } else if (d[0].classList.contains("order_" + index) && d[0].classList.contains("rot180deg")) {
                d[0].classList.remove('rot180deg');
                orderBy(col, false);
            }
        });

        r.appendChild(c1)
        r.appendChild(c2)

        const dr2 = document.createElement('div')
        dr2.classList.add('row')

        const input = document.createElement('input')
        input.type = "search"
        input.placeholder = "Search...";
        input.classList.add('search-input')
        dr2.appendChild(input)

        input.addEventListener('keyup', () => {
            if (input == null) return;

            searchFor(input.value.toLowerCase(), col.Name);
        });

        input.addEventListener('input', (event) => {
            if (event.target.value === '') clearData(allLogs)
        });

        th.appendChild(r)
        th.appendChild(dr2)
        trSpawnColums.appendChild(th);
    });
}

function searchFor(searchedInput, colName) {
    filteredLogs = filterByColumns(colName, allLogs, searchedInput);
    currentPage = 1;
    renderTable(filteredLogs);
}

function clearData(logs) {
    filteredLogs = logs;
    currentPage = 1;
    renderTable(filteredLogs)
}

function renderTable(logsToRender) {
    const start = (currentPage - 1) * logsPerPage;
    const end = start + logsPerPage;
    const paginatedLogs = logsToRender.slice(start, end);

    logTableRows.innerHTML = '';

    paginatedLogs.forEach(logEntry => {
        const tr = document.createElement('tr');

        allCols.forEach(colDef => {
            const td = document.createElement('td');
            const p = document.createElement('p');

            p.classList.add('col-p-style')

            const rowVal = DivHelper.makeRow("", {});
            const rowActions = DivHelper.makeRow("", {});
            const colActionSeeMore = DivHelper.makeCol("see more...", { class: 'see-more' });

            if (searchInput != null && searchInput.value.length != 0) {
                p.innerHTML = logEntry[colDef.Name] !== undefined ? highlightText(logEntry[colDef.Name], searchInput.value) : '';
            } else {
                p.innerText = logEntry[colDef.Name] !== undefined ? logEntry[colDef.Name] : '';
            }

            if (cleanString(p.innerText).length != 0) {
                colActionSeeMore.addEventListener('click', () => {
                    const resultDialog = new DialogComponent({
                        title: colDef.DisplayName,
                        content: p.innerHTML,
                        onClose: () => { }
                    });
                    resultDialog.open();
                });
            }

            rowVal.appendChild(p)
            td.appendChild(rowVal);
            if (cleanString(p.innerText).length != 0) {
                rowActions.appendChild(colActionSeeMore);
            }
            td.appendChild(rowActions);
            tr.appendChild(td);
        });
        logTableRows.appendChild(tr);
    });

    renderPaginationControls(logsToRender.length);
}

function orderBy(col, ordDirection) {
    filteredLogs = orderByCol(allLogs, col.Name, ordDirection)
    currentPage = 1;
    renderTable(filteredLogs);
}

function renderPaginationControls(totalLogsCount) {
    const totalPages = Math.ceil(totalLogsCount / logsPerPage);
    paginationNumbers.innerHTML = '';

    for (let i = 1; i <= totalPages; i++) {
        const pageSpan = document.createElement('span');
        pageSpan.innerText = i;
        pageSpan.classList.add('page-number');

        if (i === currentPage) {
            pageSpan.classList.add('active', 'active-page');
        }

        pageSpan.addEventListener('click', () => {
            currentPage = i;
            renderTable(filteredLogs);
        });
        paginationNumbers.appendChild(pageSpan);
    }

    prevPageBtn.disabled = currentPage === 1;
    nextPageBtn.disabled = currentPage === totalPages;
}

function changePage(direction) {
    currentPage += direction;
    renderTable(filteredLogs);
}

function filterByUserSearch() {
    filteredLogs = filterArrayBySearchTerm(allLogs, searchInput.value.toLowerCase())
    currentPage = 1;
    renderTable(filteredLogs);
}

document.addEventListener('DOMContentLoaded', () => {
    showSpinner();

    getLogs();

    searchInput.addEventListener('keyup', filterByUserSearch);
    prevPageBtn.addEventListener('click', () => changePage(-1));
    nextPageBtn.addEventListener('click', () => changePage(1));
    btnExportPdf.addEventListener('click', () => exportTableToPdf(allCols, filteredLogs));
});

function showSpinner() { loadingSpinner.classList.add('active'); }
function hideSpinner() { loadingSpinner.classList.remove('active'); }

function getLogs() {

    fetch('/api/logs')
        .then(response => {
            if (!response.ok) {
                hideSpinner();
                resetVisibilityHtmlElements()
                throw new Error(`Error HTTP! state: ${response.status} - ${response.statusText}`);
            }
            return response.json();
        })
        .then(data => {


            hideSpinner();
            resetVisibilityHtmlElements()
            allLogs = data.logs || [];
            allCols = data.cols || [];

            if (allLogs.length === 0) {
                const resultDialog = new DialogComponent({
                    title: "Attention!",
                    content: "No log founds. file can be empty or not parsable.",
                    onClose: () => { }
                });
                resultDialog.open();
                logTableRows.innerHTML = '<tr><td colspan="' + allCols.length + '">No logs available.</td></tr>';
            } else {
                filteredLogs = [...allLogs];
                renderColumns(allCols);
                renderTable(filteredLogs);

                const resultDialog = new DialogComponent({
                    title: "Success!",
                    content: "Logs loaded",
                    onClose: () => { }
                });
                resultDialog.open();

            }
        })
        .catch(error => {
            hideSpinner();
            resetVisibilityHtmlElements()
            console.error('Error in loading Logs:', error);
            const resultDialog = new DialogComponent({
                title: "Error!",
                content: 'Errore nel caricamento dei log: ' + error.message,
                onClose: () => { }
            });
            resultDialog.open();
        })
        .finally(() => {
            hideSpinner();
            resetVisibilityHtmlElements()
        });
}