export async function initWASMSimulator() {
    const go = new Go();
    const result = await WebAssembly.instantiateStreaming(
        fetch('/static/wasm/main.wasm'),
        go.importObject
    );
    go.run(result.instance);
}

export class MIPSSimulatorUI {
    constructor() {
        this.instructions = [];
        this.currentStep = 0;
    }

    loadProgram(instructions) {
        this.instructions = instructions;
        return loadProgram(this.instructions);
    }

    step() {
        const result = simulatorStep();
        // If result is not a boolean, update the UI
        this.updateUI(result);
    }

    updateUI(stepResult) {
        // Update registers
        Object.keys(stepResult.registers).forEach((reg) => {
            const targetReg = document.getElementById(reg.substring(1))
            if (targetReg) {
                const h3Tag = targetReg.querySelector('h3.value');
    
                if (h3Tag) {
                    h3Tag.innerText = stepResult.registers[reg]
                }
            }
        });

        // Update operations
        document.getElementById("op-count").innerText = stepResult.currentStep

        // Highlight current instruction
        // this.highlightInstruction(stepResult.step);
    }

    highlightInstruction(stepIndex) {
        const instructions = document.querySelectorAll('.instruction');
        instructions.forEach((el, index) => {
            el.classList.toggle('current', index === stepIndex);
        });
    }
}
