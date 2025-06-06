
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

export class TableHelper {
    static makeTableRow(innerText, style = {}) {
        let attributes = '';
        if (style.class && typeof style.class === 'string') {
            attributes += `${style.class}`;
            delete style.class;
        } else if (Array.isArray(style.class)) {
            attributes += `${style.class.join(';')}`;
            delete style.class;
        }

        for (const key in style) {
            if (style.hasOwnProperty(key)) {
                attributes += ` ${key}=${style[key]}`;
            }
        }

        attributes = attributes.replace("style=", "")
        const tr = document.createElement('tr');
        tr.style.cssText = attributes;
        tr.innerText = innerText;
        return tr;
    }

    static makeTableCol(innerText, style = {}) {
        let attributes = '';
        if (style.class && typeof style.class === 'string') {
            attributes += `${style.class}`;
            delete style.class;
        } else if (Array.isArray(style.class)) {
            attributes += `${style.class.join(';')}`;
            delete style.class;
        }

        for (const key in style) {
            if (style.hasOwnProperty(key)) {
                attributes += ` ${key}=${style[key]}`;
            }
        }

        attributes = attributes.replace("style=", "")
        const tr = document.createElement('td');
        tr.style.cssText = attributes;
        tr.innerText = innerText;
        return tr;
    }
}

export class DivHelper {

    static makeDiv(innerText, props = {}) {
        let attributes = '';
        if (props.class && typeof props.class === 'string') {
            attributes += `${props.class}`;
            delete props.class;
        } else if (Array.isArray(props.class)) {
            attributes += `${props.class.join(' ')}`;
            delete props.class;
        }

        for (const key in props) {
            if (props.hasOwnProperty(key)) {
                attributes += ` ${key}=${props[key]}`;
            }
        }

        const div = document.createElement('div');
        div.className = attributes;
        div.innerText = innerText;
        return div;
    }

    static makeCol(innerText, props = {}) {
        let classes = 'col';
        if (props.class) {
            if (typeof props.class === 'string') {
                classes += ` ${props.class}`;
            } else if (Array.isArray(props.class)) {
                classes += ` ${props.class.join(' ')}`;
            }
        }
        props.class = classes;
        return DivHelper.makeDiv(innerText, props);
    }

    static makeRow(innerText, props = {}) {
        let classes = 'row';
        if (props.class) {
            if (typeof props.class === 'string') {
                classes += ` ${props.class}`;
            } else if (Array.isArray(props.class)) {
                classes += ` ${props.class.join(' ')}`;
            }
        }
        props.class = classes;
        return DivHelper.makeDiv(innerText, props);
    }


}
