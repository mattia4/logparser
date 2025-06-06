
export function filterArrayBySearchTerm(logs, searchInput) {
    return logs.filter(log => {
        return JSON.stringify(log).toLowerCase().includes(searchInput.toLowerCase());
    });
}

export function highlightText(text, filter) {
    if (!filter || filter.trim() === '') return text;

    const escapedFilter = filter.replace('/[.*+?^${}()|[\]\\]/g', '\\$&');
    const regex = new RegExp(escapedFilter, 'gi');
    return text.replace(regex, '<span class="highlight">$&</span>');
}

export function filterByColumns(columnName, logs, searchInput) {
    if (searchInput.length == 0) return logs;

    return logs.filter(log => {
        return log[columnName] && String(cleanString(log[columnName].toLowerCase())) == String(cleanString(searchInput.toLowerCase()));
    })

}

function cleanString(str) {
    return str.replace(/\s/g, '').replace(/\u00A0/g, '');
}

export function orderByCol(logs, columnName, isAscending = true) {
    logs.sort((a, b) => {
        const valA = a[columnName];
        const valB = b[columnName];

        let comparison = 0;

        comparison = String(valA).localeCompare(String(valB));

        return isAscending ? comparison : -comparison;
    });

    return logs;
}