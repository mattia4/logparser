import { highlightText, filterArrayBySearchTerm, filterByColumns } from './table-functions.js';
import { hideHtmlElement, resetVisibilityHtmlElements } from './helpers.js';
import { exportTableToPdf } from './pdf-utils.js';

let allLogs = [];
let allData;
let allCols = [];
let filteredLogs = [];
let currentPage = 1;
const logsPerPage = 50;

const loadingSpinner = document.getElementById('loadingSpinner');
const logTableRows = document.getElementById('logTableRows');
const trSpawnColums = document.getElementById('trSpawnColums');
const searchInput = document.getElementById('searchInput');
const paginationNumbers = document.getElementById('pagination-numbers');
const prevPageBtn = document.getElementById('prevPageBtn');
const nextPageBtn = document.getElementById('nextPageBtn');
const btnExportPdf = document.getElementById('btn-export-pdf');
const outputDiv = document.getElementById('output');

hideHtmlElement(paginationNumbers)
hideHtmlElement(pagination)

function renderColumns(cols) {
    trSpawnColums.innerHTML = '';
    cols.forEach(col => {
        const th = document.createElement('th');
        //th.innerText = col.DisplayName;

        const dr = document.createElement('div')
        dr.classList.add('row')
        dr.innerText = col.DisplayName;

        const dr2 = document.createElement('div')
        dr2.classList.add('row')

        const input = document.createElement('input')
        input.type = "search"
        dr2.appendChild(input)
        input.addEventListener('keyup', () => {
            if (input == null) return;

            searchFor(input.value.toLowerCase(), col.Name);
        });

        th.appendChild(dr)
        th.appendChild(dr2)
        trSpawnColums.appendChild(th);
    });
}

function searchFor(searchedInput, colName) {
    console.log(searchedInput, colName, allData);
    filteredLogs = filterByColumns(colName, allLogs, searchedInput);

    currentPage = 1;
    renderTable(filteredLogs);
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
            if (searchInput != null && searchInput.value.length != 0) {
                td.innerHTML = logEntry[colDef.Name] !== undefined ? highlightText(logEntry[colDef.Name], searchInput.value) : '';
            } else {
                td.innerText = logEntry[colDef.Name] !== undefined ? logEntry[colDef.Name] : '';
            }
            tr.appendChild(td);
        });
        logTableRows.appendChild(tr);
    });

    renderPaginationControls(logsToRender.length);
}


function orderBy() {

}

function renderPaginationControls(totalLogsCount) {
    const totalPages = Math.ceil(totalLogsCount / logsPerPage);
    paginationNumbers.innerHTML = '';

    for (let i = 1; i <= totalPages; i++) {
        const pageSpan = document.createElement('span');
        pageSpan.innerText = i;
        pageSpan.classList.add('page-number');

        if (i === currentPage) {
            pageSpan.classList.add('active');
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
            allData = data;
            allLogs = data.logs || [];
            allCols = data.cols || [];
            console.log(data)

            if (allLogs.length === 0) {
                outputDiv.innerText = "No log founds. file can be empty or not parsable.";
                logTableRows.innerHTML = '<tr><td colspan="' + allCols.length + '">No logs available.</td></tr>';
            } else {
                filteredLogs = [...allLogs];
                renderColumns(allCols);
                renderTable(filteredLogs);
                outputDiv.innerText = "Logs loaded.";
            }
        })
        .catch(error => {
            hideSpinner();
            resetVisibilityHtmlElements()
            console.error('Error in loading Logs:', error);
            outputDiv.innerText = 'Errore nel caricamento dei log: ' + error.message;
        })
        .finally(() => {
            hideSpinner();
            resetVisibilityHtmlElements()
        });
}