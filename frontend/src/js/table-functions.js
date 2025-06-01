
export function filterArrayBySearchTerm(logs, searchInput) {
    return logs.filter(log => {
        return JSON.stringify(log).toLowerCase().includes(searchInput);
    });
}

export function highlightText(text, filter) {
    if (!filter || filter.trim() === '') return text;

    const escapedFilter = filter.replace('/[.*+?^${}()|[\]\\]/g', '\\$&');
    const regex = new RegExp(escapedFilter, 'gi');
    return text.replace(regex, '<span class="highlight">$&</span>');
}

export function filterByColumns(columnName, logs, searchInput) {
    return logs.filter(log => {
        if(log.Name == columnName && log.p) {

        }
    })
}