async function loadAndInstantiateWasm() {
    const go = new Go();
    const result = await WebAssembly.instantiateStreaming(fetch("./dist/test.wasm"), go.importObject).catch((err) => {
        console.error(err);
    });
    globalThis.goWasmModule = result.module;
    globalThis.goWasmInstance = result.instance;
    document.getElementById("runButton").disabled = false;


    // document.getElementById("reloadButton").disabled = false;
    document.getElementById("restartButton").disabled = false;
    document.getElementById("runButton").disabled = false;

    go.run(globalThis.goWasmInstance).then(
        (result) => {
            window.go = go;
            window.inst = result;
        }
    ).catch(
        (err) => {
            console.error("[WASM ERROR]", err);
        }
    );
    const version = await fetch("./dist/VERSION").then((res) => res.text());
    const wasmVersion = getVersion();
    if (version.trim() !== wasmVersion.trim()) {
        console.error(`Version mismatch: JS version is "${version}" and Wasm version is "${wasmVersion}"`);
        alert("Version mismatch! Please reload the page.");
    }
}

async function restartWasm() {
    await window.go.run(inst).catch(
        (err) => {
            console.error("[WASM ERROR]", err);
        }
    );
}

function ImgData() {
    const can = document.createElement("canvas");
    const ctx = can.getContext("2d");
    can.width = vrImg.width
    can.height = vrImg.height
    ctx.drawImage(vrImg, 0, 0)
    return ctx.getImageData(0, 0, vrImg.width, vrImg.height)
}

function drawImgData() {
    const out = goMI(document.getElementById("filter").value, ImgData());
    const can = document.createElement("canvas");
    const ctx = can.getContext("2d");
    // Size the canvas to the returned image dimensions
    can.width = out.width;
    can.height = out.height;

    const imageData = new ImageData(Uint8ClampedArray.from(out.data), out.width, out.height);
    ctx.putImageData(imageData, 0, 0);
    vrImg.src = can.toDataURL();
}
