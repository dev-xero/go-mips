class SyntaxHighlighter {
    constructor(editorElement, options = {}) {
        this.editor = editorElement;
        this.options = {
            language: options.language || 'mips',
            themes: {
                operator: 'operator',
                register: 'register',
                immediate: 'immediate',
                comment: 'comment',
            },
        };

        this.setupEventListeners();
    }

    setupEventListeners() {
        this.editor.addEventListener('input', this.handleInput.bind(this));
        this.editor.addEventListener('scroll', this.syncScroll.bind(this));
    }

    highlightCode(code) {
        const escapedCode = code
            .replace(
                /[&<>]/g,
                (char) => ({ '&': '&amp;', '<': '&lt;', '>': '&gt;' }[char])
            )
            .replace(/\n/g, '<br>');

        const lines = escapedCode.split('<br>');

        console.log("lines:", lines);

        const highlightedLines = lines.map((line) => {
            if (/^\s*#/.test(line)) {
                return `<span class="${this.options.themes.comment}">${line}</span>`;
            }

            // Process operators, registers, immediate values
            return line
                .replace(
                    /\b(add|addi|and|or|sub|mul|div|lw|sw|beq|bne|j|jr)\b/g,
                    `<span class="${this.options.themes.operator}">$1</span>`
                )
                .replace(
                    /\$[a-z0-9]+/g,
                    `<span class="${this.options.themes.register}">$&</span>`
                )
                .replace(
                    /-?\b\d+\b|0x[0-9a-fA-F]+/g,
                    `<span class="${this.options.themes.immediate}">$&</span>`
                );
        });

        return highlightedLines.join('<br />');
    }

    handleInput(e) {
        const codeElement = document.getElementById('hl-content');
        codeElement.innerHTML = this.highlightCode(e.target.value);
    }

    // We need both elements to scroll together
    syncScroll() {
        const codeElement = document.getElementById('hl');
        codeElement.scrollTop = this.editor.scrollTop;
        codeElement.scrollLeft = this.editor.scrollLeft;
    }
}

const webEditorElement = document.getElementById('editor');
const highlighter = new SyntaxHighlighter(webEditorElement, {
    language: 'mips',
});
