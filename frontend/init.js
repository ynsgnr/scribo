'use strict';

const WASM_URL = 'build/app.wasm';

var wasm;

function main(){
  var tbody = document.querySelector("tbody");
}

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

if ('content' in document.createElement('template')) {
  if (WebAssembly) {
      document.addEventListener("DOMContentLoaded", function(){document.getElementById("notSupportedMsg").remove();main()});
      initWasm();
  } else {
      console.log("web assembly is not supported in your browser")
  }
}else{
  console.log("html templates is not supported in your browser")
}

