class ScriboHeader extends HTMLElement {
        constructor() {
            super();
            this.template = document.createElement("template")
            this.template.innerHTML = `
            <div style = "width:100%; height:50px; padding:3px; justify-content: center; text-align: left; color:white; background-color:#f28c1f; ">
                <div style= "padding:3px 10px 3px 10px; align-items: center; display: flex;height: 100%; justify-content: space-between;">
                    <span>Scribo</span>
                    <button id="sign-out" type="button" style= "display: none;">
                       Sign Out
                    </button>
                </div>
            </div>
            `
            this.signOutButton = this.template.content.getElementById("sign-out")
            this.signOutButton.onclick = this.signOut
            let shadowRoot = this.attachShadow({ mode: "open" });
            shadowRoot.appendChild(this.template.content);
            this.addEventListener("signedin",()=>{this.signOutButton.style.removeProperty("display")})
        }

        signOut(){
            this.dispatchEvent(new Event("signout",{composed: true}))
        }
    }
window.customElements.define("scribo-header", ScriboHeader);