class DeviceElement extends HTMLElement {

    static get observedAttributes() { return ['name', 'type', 'selected']; }

    constructor() {
        super();
        let template = document.createElement("template")
        template.innerHTML = `
        <div id="container" style="display:flex;">
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
        this.container = root.getElementById("container")
        let shadowRoot = this.attachShadow({ mode: "open" });
        shadowRoot.appendChild(root);

        this.deviceLogoMap = {
            "kindle":"assets/kindle.svg"
        }
    }

    update(){
        this.deviceName.innerHTML = this.getAttribute("name")
        var logo = this.deviceLogoMap[this.getAttribute("type")]
        if (logo){
            this.deviceTypeLogo.setAttribute("src",logo)
        }
        if (this.getAttribute("selected")){
            this.container.style.border = "1px solid #ccc"
            this.container.style.boxShadow = "3px 3px 2px gray;"
        }else{
            this.container.style.removeProperty("boxShadow")
            this.container.style.removeProperty("border")
        }
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
window.customElements.define("device-element", DeviceElement);