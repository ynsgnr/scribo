import * as cookie from './cookie.js'
import {Baseurl} from './config.js'
import * as auth from './auth.js'

const commandEndpoint = Baseurl+"/command/v1/user/:userID/command"
const getDevicesEndpoint = Baseurl+"/sync-device/v1/user/:userID/devices"

const userIDKey = ":userID"

export function GetDevices(){
    return fetch(getDevicesEndpoint.replace(userIDKey,cookie.getCookie(auth.UserIDKey)),
    {method: 'GET',
    headers: {
        'Content-Type': 'application/json',
        'Authorization': 'Bearer '+ cookie.getCookie(auth.AccessTokenKey),
    }}).then(response=>{
        if (response.status != 200){
            throw "Can't get resources right now"
        }
        return response.json()
    })
}

export function AddDevice(deviceName, type, data){

}

export function SyncRequest(deviceID, fileLocation){

}