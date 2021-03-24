import {GetDevices,SyncRequest} from './api/api.js'
import {Upload} from './api/file.upload.js'

import {Popup} from './components/basic.popup.js'
import {Toast} from './components/basic.toast.js'
import './components/file.upload.js'
import './components/list.recycle.js'
import './components/loading.placeholder.js'

import './device.add.js'
import './device.element.js'

class ScriboDevice extends HTMLElement {
    constructor() {
        super();
        let template = document.createElement("template")
        template.innerHTML = `
        <div style="min-height: 100%; width:100%;">
            <loading-place-holder id="loading-display">
                <div slot="content" id="content">
                    <div style = "display:flex; margin: 0 auto; text-align: center; overflow:auto; height:100%;">
                        <div style = "flex:1;">
                            <p>Devices</p>
                            <recycle-list style = "text-align: left; padding:5%;" id="device-list"></recycle-list>
                            <button class="add-device" type="button" id="add-device">
                                + Add a device
                            </button>
                        </div>
                        <div style="flex:5;">
                            <file-upload id="device-details" style="display:none;">
                                <p slot="top">Previous Syncs</p>
                                <recycle-list slot="bottom" id="prev-syncs">
                                    <br>
                                </recycle-list>
                            </file-upload>
                        </div>
                    </div>
                </div>
            </loading-place-holder>
        </div>
        `
        let root = template.content
        this.loading = root.getElementById("loading-display")
        this.content = root.getElementById("content")
        this.addDeviceButton = root.getElementById("add-device")
        this.prevSyncs = root.getElementById("prev-syncs")

        this.deviceDetails = root.getElementById("device-details")
        this.deviceDetails.addEventListener("filedrop",(e)=>{this.sendFile(e.detail.files)})

        this.deviceAdd = document.createElement("device-add") //singleton
        this.addDeviceButton.onclick = ()=>{this.popup = Popup(this.shadowRoot,this.deviceAdd)}
        this.deviceAdd.addEventListener("deviceadded",()=>{this.updateWithAPI().then(()=>{this.popup && this.popup.remove()})})

        this.deviceList = root.getElementById("device-list")
        this.deviceList.base = document.createElement("device-element")
        this.deviceList.addEventListener("itemSelect",(e)=>{this.itemSelect(e.detail.selected)})

        let shadowRoot = this.attachShadow({ mode: "open" });
        shadowRoot.appendChild(root);
        this.addEventListener("signedin",()=>{this.signOutButton.style.removeProperty("display")})
    }

    sendFile(file){
        if (this.selectedDevice){
            var id = this.selectedDevice.getAttribute("deviceid")
            if (id != ""){
                Upload(file).then((fileLocation)=>{
                        console.log(fileLocation)
                        SyncRequest(id,fileLocation)
                        this.updateWithAPI()
                    },
                    ()=>{console.log("fail")},
                )
            }
        }
    }

    itemSelect(element){
        this.selectedDevice = element
        this.deviceDetails.style.removeProperty("display")
    }

    update(data){
        if (!data){
            this.loading.setAttribute("loaded","true")
            Toast(this.shadowRoot,"Failed to load data")
            return Promise.resolve()
        }
        data.sort( (a,b)=> a.deviceName<b.deviceName ? -1 : a.deviceName>b.deviceName ? 1 : 0);
        this.deviceList.items = data    
        this.loading.setAttribute("loaded","true")
        return Promise.resolve()
    }

    updateWithAPI(){
        return GetDevices().then((result)=>{this.update(result)}).catch((e)=>{this.loading.setAttribute("loaded","true");console.log(e)})
    }

    connectedCallback(){
        this.updateWithAPI()
    }
      
    adoptedCallback() {
        this.updateWithAPI()
    }

}
window.customElements.define("scribo-device", ScriboDevice);