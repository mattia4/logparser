
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

export function cleanString(str) {
    return str.replace(/\s/g, '').replace(/\u00A0/g, '');
}

export function orderByCol(logs, columnName, isAscending = true) {
    logs.sort((a, b) => {

        let comparison = 0;
        const valueA = a[columnName];
        const valueB = b[columnName];

        const numA = Number(valueA);
        const numB = Number(valueB);

        const isANumber = !isNaN(numA);
        const isBNumber = !isNaN(numB);

        if (isANumber && isBNumber) {
            comparison = numA - numB;
        }
        else if (isANumber && !isBNumber) {
            comparison = -1;
        }
        else if (!isANumber && isBNumber) {
            comparison = 1;
        }
        else {
            const strA = String(valueA || '').trim().toLowerCase();
            const strB = String(valueB || '').trim().toLowerCase();

            comparison = strA.localeCompare(strB);
        }

        return isAscending ? comparison : -comparison;
    });

    return logs;
}