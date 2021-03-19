import {GetDevices,SyncRequest} from './api.js'
import {Upload} from './api.file.upload.js'

import {Popup} from './basic.popup.js'
import './file.upload.js'

import './device.add.js'
import './device.element.js'

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
                        <div style = "text-align: left; padding:5%;" id="device-list">
                        </div>
                        <button class="add-device" type="button" id="add-device">
                            + Add a device
                        </button>
                    </div>
                    <div style="flex:5;">
                        <file-upload id="device-details" style="display:none;">
                            <p slot="top">Previous Syncs</p>
                            <div slot="bottom" id="prev-syncs">
                                <br>
                            </div>
                        </file-upload>
                    </div>
                </div>
            </div>
        </div>
        `
        let root = template.content
        this.loading = root.getElementById("loading-display")
        this.error = root.getElementById("error-display")
        this.content = root.getElementById("content")
        this.deviceList = root.getElementById("device-list")
        this.addDeviceButton = root.getElementById("add-device")
        this.prevSyncs = root.getElementById("prev-syncs")

        this.deviceDetails = root.getElementById("device-details")
        this.deviceDetails.addEventListener("filedrop",(e)=>{this.sendFile(e.detail)})

        this.deviceAdd = document.createElement("device-add") //singleton
        this.addDeviceButton.onclick = ()=>{this.popup = Popup(this.shadowRoot,this.deviceAdd)}
        this.deviceAdd.addEventListener("deviceadded",()=>{this.updateWithAPI().then(()=>{this.popup && this.popup.remove()})})

        this.deviceElements = []

        let shadowRoot = this.attachShadow({ mode: "open" });
        shadowRoot.appendChild(root);
        this.addEventListener("signedin",()=>{this.signOutButton.style.removeProperty("display")})
    }

    sendFile(file){
        var id = this.selectedDevice.getAttribute("id")
        if (id && id != ""){
            Upload(file).then((fileLocation)=>{
                    console.log(fileLocation)
                    SyncRequest(id,fileLocation)
                },
                ()=>{console.log("fail")},
            )
        }
    }

    loadedWithError(error){
        if (error){
            this.loading.style.display="none"
            this.error.style.removeProperty("display")
            this.error.innerHTML="<p>"+error+"<p/>"
        }else{
            this.error.style.display="none"
            this.loading.style.display="none"
            this.content.style.removeProperty("display")
        }
    }

    
    elementOnclick(element){
        if (element && element != this.selectedDevice){
            if(this.selectedDevice){
                this.selectedDevice.removeAttribute("selected")
            }
            element.setAttribute("selected","true")
            this.selectedDevice = element
            this.deviceDetails.style.removeProperty("display")
        }
    }

    update(data){
        console.log(data)
        if (!data){
            this.loadedWithError("Failed to load data")
            return Promise.resolve()
        }
        data.sort( (a,b)=> a.deviceName<b.deviceName ? -1 : a.deviceName>b.deviceName ? 1 : 0);
        for (var i=0;i<data.length;i++){
            if (i<this.deviceElements.length){
                this.deviceElements[i].setAttribute("id",data[i].deviceID)
                this.deviceElements[i].setAttribute("name",data[i].deviceName)
                this.deviceElements[i].setAttribute("type",data[i].deviceType)
                this.deviceElements[i].onclick = (event)=>{this.elementOnclick(event.target)}
            }else{
                let deviceElement = document.createElement("device-element")
                deviceElement.setAttribute("id",data[i].deviceID)
                deviceElement.setAttribute("name",data[i].deviceName)
                deviceElement.setAttribute("type",data[i].deviceType)
                deviceElement.onclick = (event)=>{this.elementOnclick(event.target)}
                this.deviceElements.push(deviceElement)
                this.deviceList.appendChild(deviceElement)
            }
        }
        if (i<this.deviceElements.length){
            for (var k = i; k<this.deviceElements.length; k++){
                this.deviceList.removeChild(this.deviceElements[k])
            }
            this.deviceElements = this.deviceElements.slice(0,i)
        }
        this.loadedWithError(null)
        return Promise.resolve()
    }

    updateWithAPI(){
        return GetDevices().then((result)=>{this.update(result)}).catch(()=>this.loadedWithError("Failed to load data"))
    }

    connectedCallback(){
        this.updateWithAPI()
    }
      
    adoptedCallback() {
        this.updateWithAPI()
    }

}
window.customElements.define("scribo-device", ScriboDevice);