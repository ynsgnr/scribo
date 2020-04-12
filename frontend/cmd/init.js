'use strict';

const WASM_URL = 'app.wasm';

var wasm;

function initWasm() {
  const go = new Go();
  if ('instantiateStreaming' in WebAssembly) {
    WebAssembly.instantiateStreaming(fetch(WASM_URL), go.importObject).then(function (obj) {
      wasm = obj.instance;
      go.run(wasm);
    })
  } else {
    fetch(WASM_URL).then(resp =>
      resp.arrayBuffer()
    ).then(bytes =>
      WebAssembly.instantiate(bytes, go.importObject).then(function (obj) {
        wasm = obj.instance;
        go.run(wasm);
      })
    )
  }
}

if (WebAssembly) {
    document.addEventListener("DOMContentLoaded", function(){document.getElementById("notSupportedMsg").remove()});
    initWasm();
} else {
    console.log("web assembly is not supported in your browser")
}

