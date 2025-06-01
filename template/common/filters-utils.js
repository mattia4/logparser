
function filterArrayBySearchTerm(searchedText, datas) {

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