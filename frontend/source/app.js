class App extends HTMLElement {
        constructor() {
            super();
            let template = document.createElement("template")
            template.innerHTML = `
            <div>
                <p>Custom element shadow tree content...</p>
            </div>
            `
            let shadowRoot = this.attachShadow({ mode: "open" });
            shadowRoot.appendChild(template.content.cloneNode(true));

        }
    }
window.customElements.define("app-main", App);