import {AddDevice} from './api/api.js'

//TODO make this part data oriented (get required data from backend and render form accordingly by device type)
class DeviceAdd extends HTMLElement {
    constructor() {
        super();
        let template = document.createElement("template")
        template.innerHTML = `
        <form id="DeviceInfo">
            <select name = "dropdown">
                <option value = "Kindle" selected>Kindle</option>
            </select><br>
            <label for="name">Device Name:</label><br>
            <input type="text" id="DeviceName" name="deviceName" value="Device Name"><br>
            <label for="email">Kindle Email:</label><br>
            <input type="text" id="KindleEmail" name="email" value="Kindle Email"><br>
            <input type="button" id="DeviceAddButton" value="Add Kindle">
        </form>
        `

        this.nameInput = template.content.getElementById("DeviceName")
        this.kindleEmailInput = template.content.getElementById("KindleEmail")
        template.content.getElementById("DeviceAddButton").addEventListener("click",()=>{
            AddDevice(this.nameInput.value, "kindle",{"kindleEmail":this.kindleEmailInput.value})
            .then(()=>{
                this.dispatchEvent(new Event("deviceadded",{composed: true}))
                this.nameInput.value = "Device Name"
                this.kindleEmailInput.value = "Kindle Email"
            })
        })

        this.root = template.content
        let shadowRoot = this.attachShadow({ mode: "open" });
        shadowRoot.appendChild(this.root);
    }
}
window.customElements.define("device-add", DeviceAdd);