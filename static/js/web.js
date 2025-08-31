import { initWASMSimulator, MIPSSimulatorUI } from './mips.js';

const webEditorElement = document.getElementById('editor');
const runButton = document.getElementById('run-btn');

let simulator;

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
        this.updateLineNumbers();
    }

    // We need both elements to scroll together
    syncScroll() {
        const codeElement = document.getElementById('hl');
        const lineNumbers = document.querySelector('.lns');

        codeElement.scrollTop = this.editor.scrollTop;
        codeElement.scrollLeft = this.editor.scrollLeft;
        lineNumbers.scrollTop = this.editor.scrollTop;
    }

    updateLineNumbers() {
        const editor = document.getElementById('editor');
        const lineNumbers = document.querySelector('.lns');

        const lines = editor.value.split('\n').length;

        lineNumbers.innerHTML = Array(lines)
            .fill(0)
            .map((_, i) => `<span class="ln">${i + 1}</span>`)
            .join('');
    }
}

new SyntaxHighlighter(webEditorElement, {
    language: 'mips',
});

runButton.addEventListener('click', () => {
    const assemblyCode = webEditorElement.value.split('\n');
    const filteredAssembly = assemblyCode.filter(
        (line) => !line.startsWith('#') && line.trim().length != 0
    );

    simulator.resetState();
    simulator.load(filteredAssembly);
    console.log("Simulator state:", inspectSimulator());

    for (let i = 0; i < filteredAssembly.length; i++) {
        simulator.step();
    }

    // console.log("sim:", simulator.instructions)
});

document.addEventListener('DOMContentLoaded', async () => {
    await initWASMSimulator();
    simulator = new MIPSSimulatorUI();

    // TESTING
    // const program = [
    //     'add $t0, $t1, $t2',
    //     'add $t0, $t2, $t3',
    //     'add $t0, $t4, $t7',
    // ];

    // Load the program into the simulator
    let success = simulator.load([]);

    if (success) {
        console.log('Successfully loaded program.');
        // Inspect the simulator state
        const state = inspectSimulator();
        console.log('Simulator State:', state);
    } else {
        console.error('Failed to load program');
    }
});
