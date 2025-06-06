export class DialogComponent {

    constructor(options) {
        this.options = {
            title: '',
            content: '',
            onClose: null,
            closeButtonText: 'Close',
            ...options
        };

        this.dialogElement = null;
        this._createDialogElement();
    }

    _createDialogElement() {
        this.dialogElement = document.createElement('div');
        this.dialogElement.className = 'dialog-overlay hidden';

        const contentDiv = document.createElement('div');
        contentDiv.className = 'dialog-content';

        const titleElement = document.createElement('h2');
        titleElement.textContent = this.options.title;
        contentDiv.appendChild(titleElement);

        const contentParagraph = document.createElement('p');
        contentParagraph.innerHTML = this.options.content;
        contentDiv.appendChild(contentParagraph);

        const closeButton = document.createElement('button');
        closeButton.className = 'btn btn-primary';
        closeButton.textContent = this.options.closeButtonText;
        closeButton.addEventListener('click', () => this.close());
        contentDiv.appendChild(closeButton);

        this.dialogElement.appendChild(contentDiv);

        document.body.appendChild(this.dialogElement);

        this.dialogElement.addEventListener('click', (event) => {
            if (event.target === this.dialogElement) {
                this.close();
            }
        });
    }

    open() {
        if (!this.dialogElement) return;

        this.dialogElement.classList.remove('hidden');
        this.dialogElement.offsetWidth;
        this.dialogElement.classList.add('show');
    }

    close() {
        if (!this.dialogElement) return;

        this.dialogElement.classList.remove('show');
        this.dialogElement.addEventListener('transitionend', () => {
            this.dialogElement.classList.add('hidden');
            this.dialogElement.removeEventListener('transitionend', arguments.callee);
        }, { once: true });

        if (typeof this.options.onClose === 'function') {
            this.options.onClose();
        }
    }

    update(newOptions) {
        this.options = { ...this.options, ...newOptions };

        if (this.dialogElement) {
            this.dialogElement.querySelector('h2').textContent = this.options.title;
            this.dialogElement.querySelector('p').innerHTML = this.options.content;
            const closeBtn = this.dialogElement.querySelector('.btn');
            if (closeBtn) closeBtn.textContent = this.options.closeButtonText;
        }
    }

}