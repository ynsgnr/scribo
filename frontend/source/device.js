import {GetDevices} from './api.js'

class ScriboDevice extends HTMLElement {
    constructor() {
        super();
        let template = document.createElement("template")
        template.innerHTML = `
        <div style="min-height: 100%; width:100%;">
            <div style = "margin: 0 auto; text-align: center; overflow:auto; height:100%;" id="loading-display">
                <p>Loading...<p/>
            </div>
            <div style = "margin: 0 auto; text-align: center; overflow:auto; height:100%;" id="error-display"></div>
            <div style = "display:none;" id="content">
                <div style = "display:flex; margin: 0 auto; text-align: center; overflow:auto; height:100%;">
                    <div style = "flex:1;">
                        <p>Devices</p>
                        <div style = "text-align: left; id="device-list">
                        </div>
                        <button class="add-device" type="button">
                            + Add a device
                        </button>
                    </div>
                    <div style = "flex:5;">
                        <p>Previous Syncs</p>
                    </div>
                </div>
            </div>
        </div>
        `
        this.root = template.content
        this.loading = this.root.getElementById("loading-display")
        this.error = this.root.getElementById("error-display")
        this.content = this.root.getElementById("content")
        this.deviceList = this.root.getElementById("device-list")
        let shadowRoot = this.attachShadow({ mode: "open" });
        shadowRoot.appendChild(this.root);
        this.addEventListener("signedin",()=>{this.signOutButton.style.removeProperty("display")})
    }

    loaded(showError){
        if (showError){
            this.loading.style.display="none"
            this.error.style.removeProperty("display")
            this.error.innerHTML="<p>Failed to load data<p/>"
        }
        this.error.style.display="none"
        this.loading.style.display="none"
        this.content.style.removeProperty("display")
    }

    update(data){
        if (!data){
            this.loaded(true)
            return
        }
        console.log(data)
        //this.deviceList.appendChild()
        this.loaded(false)
    }

    connectedCallback(){
        GetDevices().then((result)=>{this.update(result)})
    }
      
    adoptedCallback() {
        GetDevices().then((result)=>{this.update(result)})
    }
}
window.customElements.define("scribo-device", ScriboDevice);