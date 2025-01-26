let simulator = null;

async function initWASMSimulator() {
    const go = new Go();
    const result = await WebAssembly.instantiateStreaming(
        fetch('/static/wasm/main.wasm'),
        go.importObject
    );
    go.run(result.instance);
}

class MIPSSimulatorUI {
    constructor() {
        this.instructions = [];
        this.currentStep = 0;
    }

    loadProgram(instructions) {
        this.instructions = instructions;
        return loadProgram(this.instructions);
    }
}

document.addEventListener('DOMContentLoaded', async () => {
    await initWASMSimulator();
    const simulator = new MIPSSimulatorUI();

    // TESTING
    // const program = [
    //     'add $t0, $t1, $t2',
    //     'add $t0, $t2, $t3',
    //     'add $t0, $t4, $t7',
    // ];

    // Load the program into the simulator
    success = simulator.loadProgram([]);

    if (success) {
        console.log('Successfully loaded program.');
        // Inspect the simulator state

        const state = inspectSimulator();
        console.log('Simulator State:', state);
    } else {
        console.error('Failed to load program');
    }
});
