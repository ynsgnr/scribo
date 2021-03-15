class DeviceElement extends HTMLElement {

    static get observedAttributes() { return ['name', 'type']; }

    constructor() {
        super();
        let template = document.createElement("template")
        template.innerHTML = `
        <div style="display:flex;">
                <div style = "flex:1; margin: auto; text-align: center;">
                    <img id="device-type-logo"/>
                </div>
                <div style = "flex:3;">
                    <p id="device-name"></p>
                </div>
        </div>
        `
        let root = template.content
        this.deviceTypeLogo = root.getElementById("device-type-logo")
        this.deviceName = root.getElementById("device-name")
        let shadowRoot = this.attachShadow({ mode: "open" });
        shadowRoot.appendChild(root);

        this.deviceLogoMap = {
            "kindle":"assets/kindle.svg"
        }
    }

    updateContext(){
        this.deviceName.innerHTML = this.getAttribute("name")
        var logo = this.deviceLogoMap[this.getAttribute("type")]
        if (logo){
            this.deviceTypeLogo.setAttribute("src",logo)
        }
    }

    attributeChangedCallback(name, oldValue, newValue) {
        this.updateContext();
    }
    connectedCallback(){
        this.updateContext()
    }
    adoptedCallback() {
        this.updateContext()
    }
}
window.customElements.define("device-element", DeviceElement);