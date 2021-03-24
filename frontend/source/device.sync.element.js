class DeviceSyncElement extends HTMLElement {

    static get observedAttributes() { return ['filelocation', 'syncstate']; }

    constructor() {
        super();
        let template = document.createElement("template")
        template.innerHTML = `
        <div id="container" style="flex-direction: row; display: flex; margin-left:3%">
                <div style="margin: 1%; text-align: left;">
                    <p id="file-name"></p>
                </div>
                <div style="margin: 1%; text-align: left;">
                    <p id="status"></p>
                </div>
        </div>
        `
        let root = template.content
        this.status = root.getElementById("status")
        this.fileName = root.getElementById("file-name")
        let shadowRoot = this.attachShadow({ mode: "open" });
        shadowRoot.appendChild(root);
    }

    update(){
        var fileLink = this.getAttribute("filelocation")
        if (fileLink){
            this.fileName.innerHTML = fileLink.split("/").pop()
        }
        this.status.innerHTML = this.getAttribute("syncstate")
    }

    attributeChangedCallback(name, oldValue, newValue) {
        this.update();
    }
    connectedCallback(){
        this.update()
    }
    adoptedCallback() {
        this.update()
    }
}
window.customElements.define("device-sync-element", DeviceSyncElement);