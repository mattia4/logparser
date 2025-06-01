
export function showHtmlElement(htmlEl) {
    htmlEl.classList.add('show-element');
    htmlEl.classList.remove('hide-element');
}

export function hideHtmlElement(htmlEl) {
    htmlEl.classList.add('hide-element');
    htmlEl.classList.remove('show-element');
}

export function resetVisibilityHtmlElements() {
    document.querySelectorAll(".hide-element").forEach(htmlEl => {
        htmlEl.classList.remove("hide-element");
    })
}