import * as cookie from './cookie.js'
import {Baseurl} from './config.js'

const UserKey = "userKey"
const AccessTokenKey = "accessToken"
const RefreshTokenKey = "refreshToken"
const ExpireIn = 86400
const endpoint = Baseurl+"/auth/user"

class ScriboAuth extends HTMLElement {
    constructor() {
        super();
        let template = document.createElement("template")
        template.innerHTML = `
        <div style = "width:100%; height:100%; flex:1;">
            <span id="ScriboErrorMessage" style="display:none;">Message</span>
            <form id="ScriboLoginForm">
                <label for="email">Email:</label><br>
                <input type="text" id="ScriboAuthEmail" name="email" value="email"><br>
                <label for="pass">Password: </label><br>
                <input type="password" id="ScriboAuthPassword" name="pass" value="pass"><br>
                <input type="button" id="ScriboAuthLoginButton" value="Log In">
            </form>
            <span id="ScriboLoginLoading" style="display:none;">Loading..</span>
        </div>
        `
        this.root = this.attachShadow({ mode: "open" });
        this.root.appendChild(template.content.cloneNode(true));
        this.emailInput = this.root.getElementById("ScriboAuthEmail")
        this.passInput = this.root.getElementById("ScriboAuthPassword")
        this.formElement = this.root.getElementById("ScriboLoginForm")
        this.loadingElement = this.root.getElementById("ScriboLoginLoading")
        this.messageElement = this.root.getElementById("ScriboErrorMessage")
        this.root.getElementById("ScriboAuthLoginButton").addEventListener("click",()=>this.login(this.emailInput.value,this.passInput.value))
    }

    connectedCallback() {
        this.refreshToken()
    }

    login(username, password, token){
        this.loading()
        fetch(endpoint,{method: 'PUT', credentials: 'same-origin',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify({
            "email":username,
            "pass":password,
            "token":token,
        })}).then(response=>{
            if (response.status == 403){
                throw"Username or password is wrong"
            }
            return response.json()
        }).then(data=>this.setToken(username, data))
        .then(()=>this.signedIn())
        .then(()=>{
            event = document.createEvent("HTMLEvents");
            event.initEvent("signedin", true, true);
            this.dispatchEvent(event);
        })
        .catch((err)=>{
            this.displayMessage(err)
            this.done()
        })
    }
    
    displayMessage(msg){
        this.messageElement.style.display = "block"
        this.messageElement.innerHTML = msg
    }

    loading(){
        this.formElement.style.display = "none"
        this.messageElement.style.display = "none"
        this.loadingElement.style.display = "block"
    }

    signedIn(){
        this.formElement.style.display = "none"
        this.messageElement.style.display = "block"
        this.messageElement.innerHTML = "Signed In"
    }

    done(){
        this.formElement.style.display = "block"
        this.loadingElement.style.display = "none"
    }

    setToken(username, data){
        cookie.setCookie(AccessTokenKey,data.AccessToken)
        cookie.setCookie(RefreshTokenKey,data.RefreshToken, ExpireIn)
        cookie.setCookie(UserKey, username, ExpireIn)
    }
    
    refreshToken(){
        let userName = cookie.getCookie(UserKey)
        let refreshToken = cookie.getCookie(RefreshTokenKey)
        if (refreshToken && userName){
            this.login(userName, "", refreshToken)
        }
    }
}
window.customElements.define("scribo-auth", ScriboAuth);