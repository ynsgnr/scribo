import * as cookie from './cookie.js'
import {Baseurl} from './config.js'

export const UserKey = "userKey"
export const UserIDKey = "userIDKey"
export const AccessTokenKey = "accessToken"
export const IDTokenKey = "idToken"
const RefreshTokenKey = "refreshToken"
const ExpireIn = 86400
const endpoint = Baseurl+"/authenticator/v1/user/session"

class ScriboAuth extends HTMLElement {
    constructor() {
        super();
        this.refreshToken()
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
        this.intervalID = setInterval(this.refreshToken.bind(this), ExpireIn-200);

        this.root = this.attachShadow({ mode: "open" });
        this.root.appendChild(template.content.cloneNode(true));
        this.emailInput = this.root.getElementById("ScriboAuthEmail")
        this.passInput = this.root.getElementById("ScriboAuthPassword")
        this.formElement = this.root.getElementById("ScriboLoginForm")
        this.loadingElement = this.root.getElementById("ScriboLoginLoading")
        this.messageElement = this.root.getElementById("ScriboErrorMessage")
        this.root.getElementById("ScriboAuthLoginButton").addEventListener("click",()=>{
            this.loading()
            this.login(this.emailInput.value,this.passInput.value)
            .then(()=>this.signedIn())
            .catch((err)=>{
                this.displayMessage(err)
                this.done()
            })
        })
    }

    disconnectedCallback() {
        clearInterval(this.intervalID)
    }

    login(username, password, token){
        if (username != "" && username!= null) {
            return fetch(endpoint,{method: 'PUT',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({
                "email":username,
                "pass":password,
                "token":token,
            })}).then(response=>{
                if (response.status == 403){
                    this.dispatchEvent(new CustomEvent("authrequired"))
                    throw "Username or password is wrong"
                }
                return response.json()
            }).then(data=>{this.setToken(username, data);return data})
            .then(data=>this.validate(data.token))
        }
    }

    validate(token){
        if (token != "" && token != null){
            return fetch(endpoint,{method: 'GET',
            headers: {
                'Content-Type': 'application/json',
                'Authorization': 'Bearer '+token,
            }
            }).then(response=>{
                if (response.status != 204){
                    this.dispatchEvent(new CustomEvent("authrequired"))
                    throw "Username or password is wrong"
                }
                cookie.setCookie(UserIDKey, response.headers.get("User"))
            })
            .then(()=>{
                this.dispatchEvent(new CustomEvent("signedin"))
            })
        }
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
        cookie.setCookie(AccessTokenKey,data.token)
        cookie.setCookie(RefreshTokenKey,data.refreshToken, ExpireIn)
        cookie.setCookie(UserKey, username, ExpireIn)
        cookie.setCookie(IDTokenKey, data.idToken, ExpireIn)
    }
    
    refreshToken(){
        let userName = cookie.getCookie(UserKey)
        let refreshToken = cookie.getCookie(RefreshTokenKey)
        if (refreshToken && userName && refreshToken!="" && userName!=""){
            this.login(userName, "", refreshToken)
        }else{
            this.dispatchEvent(new CustomEvent("authrequired"))
        }

    }
}
window.customElements.define("scribo-auth", ScriboAuth);